package challengeprogrammingdatastructure

import "testing"

func TestKochCurve(t *testing.T) {
	startX, startY := float32(0), float32(0)
	endX, endY := float32(100), float32(0)
	ans := KochCurve(startX, startY, endX, endY, 1)
	for _, p := range ans {
		t.Logf("%.4f, %.4f\n", p.x, p.y)
	}
}
