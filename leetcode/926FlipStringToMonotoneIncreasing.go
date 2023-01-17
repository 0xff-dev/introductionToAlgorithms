package leetcode

func minFlipsMonoIncr(s string) int {
	endWithZero, endWithOne := 0, 0
	if s[0] == '0' {
		endWithOne = 1
	}
	if s[0] == '1' {
		endWithZero = 1
	}

	// 0 有x个
	// 1 有y个
	// 如果下一个是1，如果前一是0，
	for idx := 1; idx < len(s); idx++ {
		if s[idx] == '1' {
			// 前一个以0结尾，目前也想以0结尾，那就需要将1修改为0
			wantEndWithZero := endWithZero + 1
			endWithOne = endWithOne
			if endWithOne > endWithZero {
				endWithOne = endWithZero
			}
			endWithZero = wantEndWithZero
			continue
		}

		// 如果当前是0，那么就需要考虑，是否需要调整为1
		// 以0结尾，那么直接从前一个弄过来
		endWithZero = endWithZero
		wantEndWithOne := endWithOne
		if wantEndWithOne > endWithZero {
			wantEndWithOne = endWithZero
		}
		wantEndWithOne++
		endWithOne, endWithZero = wantEndWithOne, endWithZero
	}

	if endWithOne > endWithZero {
		return endWithZero
	}
	return endWithOne
}
