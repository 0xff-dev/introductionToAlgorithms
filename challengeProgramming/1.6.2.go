/*
	n只蚂蚁在L长度的杆子爬行，当蚂蚁到达端点就会掉下去，两只蚂蚁相遇就会各自反方向怕回去
	求所以蚂蚁的掉下杆子的最长时间和最短时间
*/
package challengeProgramming

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func AntFellOutOfPole(length int, dis []int) (int, int) {
	if len(dis) == 0 {
		return 0, 0
	}
	minTime, maxTime := dis[0], 0
	for i := 0; i < len(dis); i++ {
		minTime = max(minTime, min(dis[i], length-dis[i]))
		maxTime = max(maxTime, max(dis[i], length-dis[i]))
	}
	return minTime, maxTime
}
