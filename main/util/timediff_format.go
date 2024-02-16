package util

import (
	"fmt"
	"time"
)

func FormatTimeDiff(diff time.Duration) string {
	var totalMon = int(diff.Hours() / 24 / 30) // 计算总月数

	years := totalMon / 12  // 计算总年数
	months := totalMon % 12 // 计算剩余月数

	var output string
	if years > 0 {
		output = fmt.Sprintf("%d年", years)
		if months > 0 {
			output += fmt.Sprintf("%d个月", months)
		}
	} else {
		output = fmt.Sprintf("%d个月", totalMon)
	}
	return output
}
