package services_test

import (
	"elevator/domain"
	"elevator/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_move_elevator_to_next_floor(t *testing.T) {
	action := setupAction(domain.Down, 0)

	elevator := setupElevator(domain.Down, 7, domain.Stopped)

	services.AddNewActionToQueue(elevator, action)

	assert.NotEmpty(t, elevator.ActionQueue)

	services.MoveToNextFloor(elevator)

	assert.Empty(t, elevator.ActionQueue)
	assert.Equal(t, elevator.State, domain.Stopped)
	assert.Equal(t, elevator.CurrentFloor, 0)
}

func Test_dont_move_elevator_to_next_floor_without_items_in_queue(t *testing.T) {

	elevator := setupElevator(domain.Down, 7, domain.Stopped)

	assert.Empty(t, elevator.ActionQueue)

	services.MoveToNextFloor(elevator)

	assert.Equal(t, elevator.CurrentFloor, 7)
}
