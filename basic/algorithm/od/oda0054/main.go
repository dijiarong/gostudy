package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var smap = map[string]int{
	"(": 2,
	")": 2, // 右括号优先级应该和左括号一样，但会特殊处理
	"+": 0,
	"-": 0,
	"*": 1,
	"/": 1,
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	str, _ := reader.ReadString('\n')
	str = strings.TrimSpace(str)

	// 解析表达式为tokens，支持连续表达式和负数
	strs := tokenize(str)

	sstack, numstack := make([]string, len(strs)), make([]string, len(strs))
	sind, nind := 0, 0
	for i := range strs {
		cur := strs[i]
		if tmp, ok := smap[cur]; ok {
			if sind == 0 {
				sstack[sind] = cur
				sind++
			} else {
				// 对于右括号，需要持续计算直到遇到左括号
				if cur == ")" {
					// 继续计算直到遇到左括号
					for sind > 0 && sstack[sind-1] != "(" {
						s1, s2 := numstack[nind-2], numstack[nind-1]
						nind -= 2
						r, err := cal(s1, s2, sstack[sind-1])
						if err != nil {
							fmt.Print("ERROR")
							return
						}
						numstack[nind] = r
						nind++
						sind--
					}
					// 弹出左括号
					if sind > 0 {
						sind--
					}
					continue
				}

				// 对于其他操作符，按优先级处理
				if smap[sstack[sind-1]] >= tmp && cur != "(" {
					// 需要弹出计算
					s1, s2 := numstack[nind-2], numstack[nind-1]
					nind -= 2
					r, err := cal(s1, s2, sstack[sind-1])
					if err != nil {
						fmt.Print("ERROR")
						return
					}
					numstack[nind] = r
					nind++
					sind--
				}

				sstack[sind] = cur
				sind++
			}
		} else {
			numstack[nind] = cur
			nind++
		}
	}
	// 发现符号栈还有元素，需要弹栈计算
	for sind > 0 {
		cur := sstack[sind-1]
		// 需要弹出计算
		s1, s2 := numstack[nind-2], numstack[nind-1]
		nind -= 2
		r, err := cal(s1, s2, cur)
		if err != nil {
			fmt.Print("ERROR")
			return
		}
		numstack[nind] = r
		nind++
		sind--
	}
	fmt.Print(numstack[nind-1])
}

// 解析表达式为tokens，支持连续表达式和负数
func tokenize(expr string) []string {
	var tokens []string
	var current strings.Builder

	for i, ch := range expr {
		if unicode.IsSpace(ch) {
			continue
		}

		if unicode.IsDigit(ch) {
			current.WriteRune(ch)
		} else if ch == '-' {
			// 处理负数：如果'-'出现在开头或者前一个token是操作符或左括号，则认为是负数
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}

			// 判断是否是负数
			if i == 0 || (len(tokens) > 0 && (tokens[len(tokens)-1] == "(" || tokens[len(tokens)-1] == "+" || tokens[len(tokens)-1] == "-" || tokens[len(tokens)-1] == "*" || tokens[len(tokens)-1] == "/")) {
				// 这是一个负数的开始
				current.WriteRune(ch)
			} else {
				// 这是一个减法操作符
				tokens = append(tokens, string(ch))
			}
		} else {
			// 遇到其他运算符或括号，先保存之前的数字
			if current.Len() > 0 {
				tokens = append(tokens, current.String())
				current.Reset()
			}
			tokens = append(tokens, string(ch))
		}
	}

	// 保存最后的数字
	if current.Len() > 0 {
		tokens = append(tokens, current.String())
	}

	return tokens
}

func cal(s1, s2, calm string) (string, error) {
	a1, b1 := convert(s1)
	a2, b2 := convert(s2)

	var res1, res2 int

	if calm == "/" {
		if a2 == 0 {
			return "", errors.New("division by zero")
		}
		res1 = a1 * b2
		res2 = b1 * a2
	} else if calm == "+" {
		res1 = a1*b2 + a2*b1
		res2 = b1 * b2
	} else if calm == "-" {
		res1 = a1*b2 - a2*b1
		res2 = b1 * b2
	} else {
		if a1 == 0 || a2 == 0 {
			return "0", nil
		}
		res1 = a1 * a2
		res2 = b1 * b2
	}

	// 使用GCD算法简化分数
	return simplifyFraction(res1, res2), nil
}

// 求最大公约数
func gcd(a, b int) int {
	if a < 0 {
		a = -a
	}
	if b < 0 {
		b = -b
	}
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

// 化简分数为最简形式
func simplifyFraction(numerator, denominator int) string {
	if numerator == 0 {
		return "0"
	}

	// 处理负数情况，确保分母为正
	if denominator < 0 {
		numerator = -numerator
		denominator = -denominator
	}

	// 求最大公约数并化简
	g := gcd(numerator, denominator)
	numerator /= g
	denominator /= g

	// 如果分母为1，直接返回分子
	if denominator == 1 {
		return fmt.Sprintf("%d", numerator)
	}

	return fmt.Sprintf("%d/%d", numerator, denominator)
}

func convert(s string) (int, int) {
	a, b := 0, 1
	strs := strings.Split(s, "/")
	a, _ = strconv.Atoi(strs[0])
	if len(strs) != 1 {
		b, _ = strconv.Atoi(strs[1])
	}
	return a, b
}
