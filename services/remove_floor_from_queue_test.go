package services_test

import (
	"elevator/domain"
	"elevator/services"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_remove_floor_from_queue(t *testing.T) {
	action := setupAction(domain.Up, 5)

	elevator := setupElevator(domain.Up, 5, domain.Stopped)

	elevator.ActionQueue = []domain.Action{*action}

	services.RemoveFloorFromQueue(elevator, 5)

	assert.Empty(t, elevator.ActionQueue)
}

func Test_cannot_remove_invalid_floor(t *testing.T) {
	action := setupAction(domain.Up, 5)

	elevator := setupElevator(domain.Up, 5, domain.Stopped)

	elevator.ActionQueue = []domain.Action{*action}

	_, error := services.RemoveFloorFromQueue(elevator, 9)

	assert.Error(t, error)
}
