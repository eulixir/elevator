package services

import (
	"elevator/domain"
	"sort"
)

func Reorganize(elevator *domain.Elevator) (result [2][]int) {
	result[0] = []int{}
	result[1] = []int{}

	for _, action := range elevator.ActionQueue {
		if action.Direction == elevator.CurrentDirection {
			result[0] = append(result[0], action.Floor)
			sort.Ints(result[0])
		} else {
			result[1] = append(result[1], action.Floor)
			sortDescending(result[0])
		}
	}

	return
}

func sortDescending(arr []int) {
	sort.Slice(arr, func(i, j int) bool {
		return arr[i] > arr[j]
	})
}
