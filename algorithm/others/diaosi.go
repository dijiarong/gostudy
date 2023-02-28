package others

import (
	"strings"
	"time"
)

func ShuangXiu() int {
	res := 0
	startTime, err := time.ParseInLocation("2006-01-02", "2023-01-01", time.Local)
	if err != nil {
		return res
	}
	endTime, err := time.ParseInLocation("2006-01-02", "2024-01-01", time.Local)
	if err != nil {
		return res
	}
	for startTime.Unix() < endTime.Unix() {
		// 周六
		six := startTime.AddDate(0, 0, 6).Format("2006-01-02")
		// 周日
		startTime = startTime.AddDate(0, 0, 7)
		seven := startTime.Format("2006-01-02")
		if strings.HasSuffix(six, "2") || strings.HasSuffix(six, "5") || strings.HasSuffix(six, "8") {
			continue
		}
		if strings.HasSuffix(seven, "2") || strings.HasSuffix(seven, "5") || strings.HasSuffix(seven, "8") {
			continue
		}
		res += 1
	}
	return res
}
