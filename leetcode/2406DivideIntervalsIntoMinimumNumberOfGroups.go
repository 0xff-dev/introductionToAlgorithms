package leetcode

import "sort"

func minGroups(intervals [][]int) int {
	events := []struct {
		time      int
		typeEvent int
	}{}

	// Convert intervals to events
	for _, interval := range intervals {
		events = append(events, struct{ time, typeEvent int }{interval[0], 1})
		events = append(events, struct{ time, typeEvent int }{interval[1] + 1, -1})
	}

	// Sort the events
	sort.Slice(events, func(i, j int) bool {
		if events[i].time == events[j].time {
			return events[i].typeEvent < events[j].typeEvent
		}
		return events[i].time < events[j].time
	})

	concurrentIntervals := 0
	maxConcurrentIntervals := 0

	// Calculate maximum number of concurrent intervals
	for _, event := range events {
		concurrentIntervals += event.typeEvent
		if concurrentIntervals > maxConcurrentIntervals {
			maxConcurrentIntervals = concurrentIntervals
		}
	}

	return maxConcurrentIntervals
}
