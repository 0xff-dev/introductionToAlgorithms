package offer

func QuickSort(data []int) {
	if len(data) == 0 || len(data) == 1 {
		return
	}
	walker, pre := 0, 0
	compareData := data[0]
	for walker < len(data) {
		if data[walker] < compareData {
			pre++
			data[pre], data[walker] = data[walker], data[pre]
		}
		walker++
	}
	data[0], data[pre] = data[pre], data[0]
	QuickSort(data[:pre+1])
	QuickSort(data[pre+1:])
}

func Partition(data []int) int {
	if len(data) == 1 {
		return 0
	}
	if len(data) == 0 {
		return -1
	}
	walker, pre := 0, 0
	compareData := data[0]
	for walker < len(data) {
		if data[walker] < compareData {
			pre++
			data[pre], data[walker] = data[walker], data[pre]
		}
		walker++
	}
	data[0], data[pre] = data[pre], data[0]
	return pre
}

func QuickSort1(data []int) {
	pos := Partition(data)
	if pos != -1 && pos != 0 {
		QuickSort1(data[:pos+1])
		QuickSort1(data[pos+1:])
	}
}
