package challengeprogrammingdatastructure

import "fmt"

// 实现一个方法，支持插入字符串，查询字符串
// 字符串仅包含 ACGT4种字符。
// 这可以考虑散列，前缀树
type op struct {
	opName string
	target string
}

func searchWordByHash(ops []op) {
	charMap := make(map[string]struct{})
	for _, p := range ops {
		if p.opName == "insert" {
			charMap[p.target] = struct{}{}
			continue
		}
		output := "YES"
		if _, ok := charMap[p.target]; !ok {
			fmt.Println("NO")
		}
		fmt.Println(output)
	}
}
