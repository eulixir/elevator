package domain_test

import (
	"elevator/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrganizeQueue(t *testing.T) {
	firstAction := domain.NewAction()
	secondAction := domain.NewAction()
	thrirdAction := domain.NewAction()
	fourthAction := domain.NewAction()

	firstAction.Direction = domain.Up
	firstAction.Floor = 5

	secondAction.Direction = domain.Down
	secondAction.Floor = 2

	thrirdAction.Direction = domain.Up
	thrirdAction.Floor = 9

	fourthAction.Direction = domain.Up
	fourthAction.Floor = 7

	require.NoError(t, firstAction.Validate())
	require.NoError(t, secondAction.Validate())
	require.NoError(t, thrirdAction.Validate())
	require.NoError(t, fourthAction.Validate())

	elevator := domain.NewElevator()

	elevator.ActionQueue = []domain.Action{*firstAction, *secondAction, *thrirdAction, *fourthAction}
	elevator.CurrentDirection = domain.Up
	elevator.CurrentFloor = 3
	elevator.State = domain.Moving

	organizedActionQueue := domain.Reorganize(elevator)

	require.NoError(t, elevator.Validate())

	expectQueue := []int{5, 7, 9}

	require.Equal(t, organizedActionQueue[0], expectQueue)

	expectQueue = []int{2}

	require.Equal(t, organizedActionQueue[1], expectQueue)
}
