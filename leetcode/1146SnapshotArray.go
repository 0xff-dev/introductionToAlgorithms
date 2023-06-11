package leetcode

type idWithVal struct {
	id, val int
}
type SnapshotArray struct {
	snpaId int
	data   [][]idWithVal
}

func Constructor1146(length int) SnapshotArray {
	data := make([][]idWithVal, length)
	for i := 0; i < length; i++ {
		data[i] = []idWithVal{{0, 0}}
	}
	return SnapshotArray{snpaId: 0, data: data}
}

func (this *SnapshotArray) Set(index int, val int) {
	length := len(this.data[index])
	if this.data[index][length-1].id != this.snpaId {
		this.data[index] = append(this.data[index], idWithVal{id: this.snpaId, val: val})
		return
	}
	this.data[index][length-1].val = val
}

func (this *SnapshotArray) Snap() int {
	id := this.snpaId
	this.snpaId++
	return id
}

func (this *SnapshotArray) Get(index int, snap_id int) int {
	// 该用二分, 很明显的snap_id 是递增的
	length := len(this.data[index])
	for i := length - 1; i >= 0; i-- {
		if this.data[index][i].id <= snap_id {
			return this.data[index][i].val
		}
	}
	return 0
}
