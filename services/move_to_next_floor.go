package services

import (
	"elevator/domain"
)

// Tem que ser uma goroutine, pois se um novo andar for adicionado, tem que atualizar
func MoveToNextFloor(elevator *domain.Elevator) {

	if len(elevator.ActionQueue) == 0 {
		return
	}

	queue := Reorganize(elevator)

	nextFloor := queue[0][0]

	elevator.State = domain.Moving

	elevator.CurrentFloor = nextFloor

	elevator.State = domain.Stopped

	RemoveFloorFromQueue(elevator, nextFloor)
}
