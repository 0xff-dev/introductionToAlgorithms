package leetcode

func sumGame(num string) bool {
	n := len(num)
	leftSum, rightSum := 0, 0
	leftQ, rightQ := 0, 0

	// 1. 统计左右两边的数字和及问号数量
	for i := 0; i < n/2; i++ {
		if num[i] == '?' {
			leftQ++
		} else {
			leftSum += int(num[i] - '0')
		}
	}

	for i := n / 2; i < n; i++ {
		if num[i] == '?' {
			rightQ++
		} else {
			rightSum += int(num[i] - '0')
		}
	}

	diffS := leftSum - rightSum
	diffQ := rightQ - leftQ

	// 3. 判断获胜条件
	// Bob 获胜的唯一可能：
	// 问号差值必须能正好抵消总和差值，且每两个问号贡献 9 点。
	// 即：diffS == (diffQ / 2) * 9
	// 注意：diffQ 必须为偶数，且 diffS 与 diffQ 符号必须一致（这里假设右边多问号，则左边应多和）
	/*
		假设右半边比左半边多出 $Q_{diff}$ 个问号。游戏结束时，为了让左右相等，右半边多出的这些问号所填数字的总和必须恰好等于左半边多出的原始数字之和（即 $\Delta S = S_L - S_R$）。在博弈中，问号是成对消耗的（Alice 走一步，Bob 走一步）：如果 Alice 在问号多的一侧填入 $x$：Bob 立即在同一侧的另一个问号填入 $9-x$。这两个位置的和固定为 9。如果 Alice 在问号少的一侧填入 $x$：Bob 在问号多的一侧也填入 $x$（如果可能），从而实现对称抵消。因此，每 2 个问号（一轮回合）对 Bob 来说，最稳妥的控制范围就是 9。
	*/

	return diffS*2 != diffQ*9
}
