package leetcode

type to struct {
	distance, count int
}

type inStation struct {
	station string
	start   int
}
type UndergroundSystem struct {
	in  map[int]inStation
	out map[string]map[string]to
}

func Constructor1396() UndergroundSystem {
	return UndergroundSystem{
		in:  make(map[int]inStation),
		out: make(map[string]map[string]to),
	}
}

func (this *UndergroundSystem) CheckIn(id int, stationName string, t int) {
	this.in[id] = inStation{stationName, t}
}

func (this *UndergroundSystem) CheckOut(id int, stationName string, t int) {
	from := this.in[id]
	cost := t - from.start

	v, ok := this.out[from.station]
	if !ok {
		v = map[string]to{
			stationName: {cost, 1},
		}
	} else {
		if vv, ok := v[stationName]; !ok {
			v[stationName] = to{cost, 1}
		} else {
			v[stationName] = to{cost + vv.distance, vv.count + 1}
		}
	}
	this.out[from.station] = v
}

func (this *UndergroundSystem) GetAverageTime(startStation string, endStation string) float64 {
	v := this.out[startStation][endStation]
	return float64(v.distance) / float64(v.count)
}
