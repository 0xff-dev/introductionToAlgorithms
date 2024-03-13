package leetcode

func flipAndInvertImage(image [][]int) [][]int {
	rows, cols := len(image), len(image[0])
	for r := 0; r < rows; r++ {
		for s, e := 0, cols-1; s <= e; s, e = s+1, e-1 {
			image[r][s], image[r][e] = 1-image[r][e], 1-image[r][s]
		}
	}
	return image
}
