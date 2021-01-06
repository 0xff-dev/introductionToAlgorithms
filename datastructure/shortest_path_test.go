package datastructure

import "testing"

func initGraph() []edge {
	r := make([]edge, 0)
	r = append(r, edge{'s', 't', 6})
	r = append(r, edge{'s', 'y', 7})
	r = append(r, edge{'t', 'x', 5})
	r = append(r, edge{'t', 'z', -4})
	r = append(r, edge{'t', 'y', 8})
	r = append(r, edge{'y', 'x', -3})
	r = append(r, edge{'y', 'z', 9})
	r = append(r, edge{'x', 't', -2})
	r = append(r, edge{'z', 'x', 7})
	r = append(r, edge{'z', 's', 2})
	return r
}

func TestBellmanFord(t *testing.T) {
	r := initGraph()
	for _, char := range []byte{'s', 't', 'z', 'y', 'x'} {
		found, res := BellmanFord(r, 's', char)
		if !found {
			t.Fatalf("%c in graph, wrong answer", char)
		}
		t.Log("min distance is: ", res)
	}
}

func initDijkstraGraph() []edge {
	r := make([]edge, 0)
	r = append(r, edge{'s', 't', 10})
	r = append(r, edge{'s', 'y', 5})
	r = append(r, edge{'t', 'x', 1})
	r = append(r, edge{'t', 'y', 2})
	r = append(r, edge{'y', 't', 3})
	r = append(r, edge{'y', 'x', 9})
	r = append(r, edge{'y', 'z', 2})
	r = append(r, edge{'x', 'z', 4})
	r = append(r, edge{'z', 'x', 6})
	r = append(r, edge{'z', 's', 7})
	return r
}

func TestDijkstra(t *testing.T) {
	graph := initDijkstraGraph()
	r := Dijkstra(graph, 's')
	for aim, dist := range r {
		t.Logf("start: s to: %s -- %d", string(aim), dist)
	}
}
