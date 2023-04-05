package challengeprogrammingdatastructure

import "fmt"

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
