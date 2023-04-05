package challengeprogrammingdatastructure

import "math"

func KochCurve(startX, startY, endX, endY float32, n int) []point {
	target := []point{{float32(startX), float32(startY)}}
	if n == 0 {
		return []point{{endX, endY}}
	}
	xSteps := (endX - startX) / 3.0
	ySteps := (endY - startY) / 3.0
	sX, sY := startX+xSteps, startY+ySteps
	uX, uY := startX+xSteps*2.0, startY+ySteps*2.0
	math.Sin(60)
	tX := (uX-sX)*float32(math.Cos(math.Pi/3)) - (uY-sY)*float32(math.Sin(math.Pi/3)) + sX
	tY := (uX-sX)*float32(math.Sin(math.Pi/3)) + (uY-sY)*float32(math.Cos(math.Pi/3)) + sY
	target = append(target, KochCurve(startX, startY, sX, sY, n-1)...)
	target = append(target, KochCurve(sX, sY, tX, tY, n-1)...)
	target = append(target, KochCurve(tX, tY, uX, uY, n-1)...)
	target = append(target, KochCurve(uX, uY, endX, endY, n-1)...)
	return target
}
