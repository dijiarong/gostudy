package findmissingobservations

/*
现有一份n+m次投掷单个六面骰子的观测数据，骰子的每个面从1到6编号。观测数据中缺失了n份，你手上只拿到剩余m次投掷的数据。幸好你有之前计算过的这n+m次投掷数据的平均值 。
给你一个长度为m的整数数组rolls，其中rolls[i]是第i次观测的值。同时给你两个整数mean和n。
返回一个长度为n的数组，包含所有缺失的观测数据，且满足这n+m次投掷的平均值是mean。如果存在多组符合要求的答案，只需要返回其中任意一组即可。如果不存在答案，返回一个空数组。
k个数字的平均值 为这些数字求和后再除以k。
注意mean是一个整数，所以n+m次投掷的总和需要被n+m整除
*/
func missingRolls(rolls []int, mean int, n int) []int {
	target := (len(rolls) + n) * mean
	for _, v := range rolls {
		target -= v
	}
	if target < n || target > 6*n {
		return nil
	}
	res := make([]int, n)
	avg, mod := target/n, target%n
	for i := range res {
		res[i] = avg
		if i < mod {
			res[i]++
		}
	}
	return res
}
