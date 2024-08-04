package services

import (
	"elevator/models"
)

func AddNewActionToQueue(elevator *models.Elevator, newAction *models.Action) *models.Elevator {
	if elevator.CurrentFloor == newAction.Floor {
		handleSameFloorAction(elevator, newAction)
		return elevator
	}

	setNewActionDirection(elevator, newAction)
	elevator.ActionQueue = append(elevator.ActionQueue, *newAction)

	return elevator
}

func handleSameFloorAction(elevator *models.Elevator, newAction *models.Action) {
	if elevator.State == models.Moving {
		toggleDirection(newAction, elevator.CurrentDirection)
		elevator.ActionQueue = append(elevator.ActionQueue, *newAction)
	}
}

func setNewActionDirection(elevator *models.Elevator, newAction *models.Action) {
	if elevator.CurrentFloor > newAction.Floor {
		newAction.Direction = models.Down
	} else {
		newAction.Direction = models.Up
	}
}

func toggleDirection(action *models.Action, currentDirection models.Direction) {
	if currentDirection == models.Up {
		action.Direction = models.Down
	} else {
		action.Direction = models.Up
	}
}
