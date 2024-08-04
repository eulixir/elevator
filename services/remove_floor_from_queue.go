package services

import (
	"elevator/domain"
	"fmt"
)

func RemoveFloorFromQueue(elevator *domain.Elevator, floor int) (domain.Elevator, error) {

	index := findIndex(elevator, floor)

	if index < 0 {
		return *elevator, fmt.Errorf("floor not found")
	}

	elevator.ActionQueue = append(elevator.ActionQueue[:index], elevator.ActionQueue[index+1:]...)

	return *elevator, nil
}

func findIndex(elevator *domain.Elevator, targetFloor int) int {
	for i, action := range elevator.ActionQueue {
		if action.Floor == targetFloor {
			return i
		}
	}
	return -1
}
