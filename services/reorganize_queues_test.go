package services_test

import (
	"testing"

	"elevator/domain"
	"elevator/services"

	"github.com/stretchr/testify/require"
)

func Test_organize_queue(t *testing.T) {
	firstAction := setupAction(domain.Up, 5)
	secondAction := setupAction(domain.Down, 2)
	thirdAction := setupAction(domain.Up, 9)
	fourthAction := setupAction(domain.Up, 7)

	require.NoError(t, firstAction.Validate())
	require.NoError(t, secondAction.Validate())
	require.NoError(t, thirdAction.Validate())
	require.NoError(t, fourthAction.Validate())

	elevator := setupElevator(domain.Up, 3, domain.Moving)
	elevator.ActionQueue = []domain.Action{*firstAction, *secondAction, *thirdAction, *fourthAction}

	organizedActionQueue := services.Reorganize(elevator)

	require.NoError(t, elevator.Validate())

	expectedUpQueue := []int{5, 7, 9}
	expectedDownQueue := []int{2}

	require.Equal(t, expectedUpQueue, organizedActionQueue[0], "Expected actions going up to be organized correctly")
	require.Equal(t, expectedDownQueue, organizedActionQueue[1], "Expected actions going down to be organized correctly")
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
