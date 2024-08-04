package services

import (
	"elevator/domain"
)

func AddNewActionToQueue(elevator *domain.Elevator, newAction *domain.Action) *domain.Elevator {
	if elevator.CurrentFloor == newAction.Floor {
		handleSameFloorAction(elevator, newAction)
		return elevator
	}

	setNewActionDirection(elevator, newAction)
	elevator.ActionQueue = append(elevator.ActionQueue, *newAction)

	return elevator
}

func handleSameFloorAction(elevator *domain.Elevator, newAction *domain.Action) {
	if elevator.State == domain.Moving {
		toggleDirection(newAction, elevator.CurrentDirection)
		elevator.ActionQueue = append(elevator.ActionQueue, *newAction)
	}
}

func setNewActionDirection(elevator *domain.Elevator, newAction *domain.Action) {
	if elevator.CurrentFloor > newAction.Floor {
		newAction.Direction = domain.Down
	} else {
		newAction.Direction = domain.Up
	}
}

func toggleDirection(action *domain.Action, currentDirection domain.Direction) {
	if currentDirection == domain.Up {
		action.Direction = domain.Down
	} else {
		action.Direction = domain.Up
	}
}
