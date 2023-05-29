package leetcode

type ParkingSystem struct {
	solt [3]int
}

func Constructor1604(big int, medium int, small int) ParkingSystem {
	return ParkingSystem{solt: [3]int{big, medium, small}}
}

func (this *ParkingSystem) AddCar(carType int) bool {
	if this.solt[carType-1] > 0 {
		this.solt[carType-1]--
		return true
	}
	return false
}
