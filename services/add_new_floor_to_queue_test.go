package services_test

import (
	"testing"

	"elevator/domain"
	"elevator/services"

	"github.com/stretchr/testify/assert"
)

func Test_add_new_floor_when_elevator_is_moving_up(t *testing.T) {
	elevator := setupElevator(domain.Up, 0, domain.Moving)
	action := setupAction(domain.Up, 4)

	assert.Empty(t, elevator.ActionQueue, "Expected action queue to be empty before adding action")

	services.AddNewActionToQueue(elevator, action)

	assert.Equal(t, elevator.ActionQueue, []domain.Action{*action}, "Expected action queue to contain the new action")
}

func Test_add_new_up_floor_when_elevator_is_going_down(t *testing.T) {
	elevator := setupElevator(domain.Up, 5, domain.Moving)
	action := setupAction(domain.Down, 4)

	assert.Empty(t, elevator.ActionQueue, "Expected action queue to be empty before adding action")

	services.AddNewActionToQueue(elevator, action)

	assert.Equal(t, elevator.ActionQueue, []domain.Action{*action}, "Expected action queue to contain the new action")
}

func Test_add_same_floor_action_when_elevator_is_stopped(t *testing.T) {
	elevator := setupElevator(domain.Up, 5, domain.Stopped)
	action := setupAction(domain.Down, 5)

	assert.Empty(t, elevator.ActionQueue, "Expected action queue to be empty before adding action")

	services.AddNewActionToQueue(elevator, action)

	assert.Empty(t, elevator.ActionQueue, "Expected action queue to remain empty after adding action on the same floor when elevator is stopped")
}

func Test_add_same_floor_action_when_elevator_is_moving(t *testing.T) {
	elevator := setupElevator(domain.Up, 5, domain.Moving)
	existingAction := setupAction(domain.Up, 7)
	elevator.ActionQueue = append(elevator.ActionQueue, *existingAction)

	sameFloorAction := setupAction(domain.Up, 5)

	assert.Equal(t, elevator.ActionQueue, []domain.Action{*existingAction}, "Expected action queue to contain the existing action before adding new action")

	services.AddNewActionToQueue(elevator, sameFloorAction)

	expectedAction := domain.Action{
		Direction: domain.Down,
		Floor:     5,
	}

	assert.Equal(t, elevator.ActionQueue, []domain.Action{*existingAction, expectedAction}, "Expected action queue to contain the existing action and the new action with toggled direction")
}

func Test_set_new_action_direction_to_down_when_elevator_is_moving_up_and_new_floor_is_lower(t *testing.T) {
	elevator := setupElevator(domain.Up, 5, domain.Moving)
	existingAction := setupAction(domain.Up, 7)
	elevator.ActionQueue = append(elevator.ActionQueue, *existingAction)

	newAction := setupAction(domain.Up, 5)

	assert.Equal(t, elevator.ActionQueue, []domain.Action{*existingAction}, "Expected action queue to contain the existing action before adding new action")

	services.AddNewActionToQueue(elevator, newAction)

	expectedAction := domain.Action{
		Direction: domain.Down,
		Floor:     5,
	}

	assert.Equal(t, elevator.ActionQueue, []domain.Action{*existingAction, expectedAction}, "Expected action queue to contain the existing action and the new action with direction set to down")
}

func Test_set_new_action_direction_to_up_when_elevator_is_moving_down_and_new_floor_is_higher(t *testing.T) {
	elevator := setupElevator(domain.Down, 9, domain.Moving)
	existingAction := setupAction(domain.Down, 7)
	elevator.ActionQueue = append(elevator.ActionQueue, *existingAction)

	newAction := setupAction(domain.Up, 5)

	assert.Equal(t, elevator.ActionQueue, []domain.Action{*existingAction}, "Expected action queue to contain the existing action before adding new action")

	services.AddNewActionToQueue(elevator, newAction)

	expectedAction := domain.Action{
		Direction: domain.Down,
		Floor:     5,
	}

	assert.Equal(t, elevator.ActionQueue, []domain.Action{*existingAction, expectedAction}, "Expected action queue to contain the existing action and the new action with direction set to down")
}

func setupElevator(direction domain.Direction, floor int, state domain.State) *domain.Elevator {
	elevator := domain.NewElevator()
	elevator.CurrentDirection = direction
	elevator.CurrentFloor = floor
	elevator.State = state
	return elevator
}

func setupAction(direction domain.Direction, floor int) *domain.Action {
	action := domain.NewAction()
	action.Direction = direction
	action.Floor = floor
	return action
}
