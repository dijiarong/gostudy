package maximum_manhattan_distance_after_k_changes

var directionMap = map[byte][2]int{
	'N': {0, 1},
	'S': {0, -1},
	'E': {1, 0},
	'W': {-1, 0},
}

var nexts = map[byte][]byte{
	'N': {'S', 'E', 'W'},
	'S': {'N', 'E', 'W'},
	'E': {'S', 'N', 'W'},
	'W': {'S', 'E', 'N'},
}

// 简陋方法,十分丑陋
func maxDistance(s string, k int) int {
	// 回溯
	// 当前坐标
	res := 0
	track := [2]int{0, 0}
	time := k
	var trackback func(index int)
	trackback = func(index int) {
		res = max(res, getDistance(track))
		if index == len(s) {
			return
		}
		track[0], track[1] = track[0]+directionMap[s[index]][0], track[1]+directionMap[s[index]][1]
		trackback(index + 1)
		track[0], track[1] = track[0]-directionMap[s[index]][0], track[1]-directionMap[s[index]][1]
		if time != 0 {
			for _, next := range nexts[s[index]] {
				track[0], track[1] = track[0]+directionMap[next][0], track[1]+directionMap[next][1]
				time--
				trackback(index + 1)
				track[0], track[1] = track[0]-directionMap[next][0], track[1]-directionMap[next][1]
				time++
			}
		}
	}
	trackback(0)
	return res
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

func getDistance(cur [2]int) int {
	x, y := cur[0], cur[1]
	if x < 0 {
		x = -x
	}
	if y < 0 {
		y = -y
	}
	return x + y
}
