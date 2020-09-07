package challengeProgramming

import "testing"

func TestAntFellOutOfPole(t *testing.T) {
	length := 10
	ants := []int{2, 6, 7}
	minTime, maxTime := AntFellOutOfPole(length, ants)
	t.Log(minTime, " ", maxTime)
}
