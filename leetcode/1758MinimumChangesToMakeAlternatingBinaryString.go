package leetcode

func minOperation1758(s string) int {
	pz, po := 0, 0
	if s[0] == '1' {
		pz = 1
	} else {
		po = 1
	}
	for idx := 1; idx < len(s); idx++ {
		if s[idx] == '1' {
			pre := po
			po = pz
			pz = pre + 1
			continue
		}
		pre := pz
		pz = po
		po = pre + 1
	}
	if pz > po {
		return po
	}
	return pz
}
