// 实现基础math功能：（加、减、乘、除）
package mymath

type myMath struct{}

var Math myMath

// 返回较大值
func (my myMath) GetMaxInt(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

// 加法
func (my myMath) Add(a, b int) int {
	if b == 0 {
		return a
	}
	tmp := a ^ b
	b = a & b << 1
	return my.Add(tmp, b)
}

// 减法
func (my myMath) Sub(a, b int) int {
	return my.Add(a, my.Add(^b, 1))
}

// 判断负数
func (myMath) IsNeg(a int) bool {
	return a < 0
}

// 取相反数
func (my myMath) Neg(a int) int {
	return my.Add(^a, 1)
}

// 除法
func (my myMath) Div(a, b int) int {
	res := 0
	c := a
	if a < 0 {
		c = my.Add(^a, 1)
	}
	d := b
	if b < 0 {
		d = my.Add(^b, 1)
	}
	for i := 30; i >= 0; i-- {
		if (c >> i) >= d {
			res = my.Add(res, 1<<i)
			c = my.Sub(c, my.Multi(d, 1<<i))
		}
	}
	return res
}

// 乘法
func (my myMath) Multi(a, b int) int {
	res := 0
	if a == 0 || b == 0 {
		return res
	}
	c := a
	if a < 0 {
		c = my.Add(^a, 1)
	}
	d := b
	if b < 0 {
		d = my.Add(^b, 1)
	}
	for d != 0 {
		if (d & 1) != 0 {
			res = my.Add(res, c)
		}
		c <<= 1
		d >>= 1
	}
	if (my.IsNeg(a) && my.IsNeg(b)) || (!my.IsNeg(a) && !my.IsNeg(b)) {
		return res
	}
	return my.Neg(res)
}
