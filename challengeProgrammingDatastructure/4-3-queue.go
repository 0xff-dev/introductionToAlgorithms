package challengeprogrammingdatastructure

import (
	"fmt"
	"strings"
)

type taskDesc struct {
	name  string
	times int
}

func CpuTasks(n, p int, tasks []taskDesc) string {
	left := n
	cost := 0
	index := 0
	buf := strings.Builder{}
	for ; left > 0; index = (index + 1) % n {
		if tasks[index].times <= 0 {
			continue
		}
		if tasks[index].times <= p {
			left--
			cost += tasks[index].times
			tasks[index].times = 0
			buf.WriteString(fmt.Sprintf("%s%d", tasks[index].name, cost))
			continue
		}
		tasks[index].times -= p
		cost += p
	}
	return buf.String()
}
