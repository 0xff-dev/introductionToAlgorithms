package leetcode

func angleClock(hour int, minutes int) float64 {
	// 分钟走的度数
	mi := (float64(minutes) / 60.0) * 360.0

	// 时针整点走的度数
	h := (hour % 12) * 30
	// 时针多走的部分
	ho := float64(h) + (float64(minutes)/60.0)*30.0
	diff := mi - ho
	if diff < 0 {
		diff = -diff
	}
	return min(diff, 360.0-diff)
}
