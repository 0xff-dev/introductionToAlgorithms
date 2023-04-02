package challengeprogrammingdatastructure

type MyHeap struct {
	data []int
	less func(i, j int) bool
}

func (mh *MyHeap) Insert(v int) {
	mh.data = append(mh.data, v)
	index := len(mh.data) - 1
	// 5,2,3,4, 1
	// 0 1 2 3 4
	/*
				  0
			1			2
		3		4	5 		6
	*/
	father := index / 2
	if index&1 == 0 {
		father--
	}
	for index > 0 && !mh.less(mh.data[father], mh.data[index]) {
		mh.data[father], mh.data[index] = mh.data[index], mh.data[father]
		index = father
		father = index / 2
		if index&1 == 0 {
			father--
		}
	}
}

func (mh *MyHeap) Pop() int {
	l := len(mh.data)
	mh.data[0], mh.data[l-1] = mh.data[l-1], mh.data[0]
	x := mh.data[l-1]
	mh.data = mh.data[:l-1]
	maintain(0, l-1, mh.data, mh.less)
	return x
}
