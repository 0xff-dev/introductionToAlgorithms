package leetcode

// 00000010100101000001111010011100
// 00111001011110000010100101000000
func reverseBits(num uint32) uint32 {
	// 01 10 ,期待的结果是 10 01
	// 10 10 -> 0010 >>1 -> 0001
	// 01 01 -> 0100 <<1 -> 1000
	// -> 1001
	//
	// 我的想法是：先两位两位的反转然后四位一组反转，然后八位一组最后，16为一组结束
	// 1010 1010 1010 1010 1010 1010
	var (
		a, b, aa, bb uint32
	)
	// 1
	a = 0xAAAAAAAA
	b = 0x55555555
	aa = (num & a) >> 1
	bb = (num & b) << 1
	num = aa | bb

	// 2位反转
	// 1100
	// 0011
	a = 0xCCCCCCCC
	b = 0x33333333
	aa = (num & a) >> 2
	bb = (num & b) << 2
	num = aa | bb

	// 111100001111000011110000
	a = 0xF0F0F0F0
	b = 0x0F0F0F0F
	aa = (num & a) >> 4
	bb = (num & b) << 4
	num = aa | bb

	// 8
	a = 0xFF00FF00
	b = 0x00FF00FF
	aa = (num & a) >> 8
	bb = (num & b) << 8
	num = aa | bb

	a = 0xFFFF0000
	b = 0x0000FFFF
	aa = (num & a) >> 16
	bb = (num & b) << 16
	num = aa | bb

	// 16
	return num
}
