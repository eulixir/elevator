package models_test

import (
	"elevator/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestOrganizeQueue(t *testing.T) {
	firstAction := models.NewAction()
	secondAction := models.NewAction()
	thrirdAction := models.NewAction()
	fourthAction := models.NewAction()

	firstAction.Direction = models.Up
	firstAction.Floor = 5

	secondAction.Direction = models.Down
	secondAction.Floor = 2

	thrirdAction.Direction = models.Up
	thrirdAction.Floor = 9

	fourthAction.Direction = models.Up
	fourthAction.Floor = 7

	require.NoError(t, firstAction.Validate())
	require.NoError(t, secondAction.Validate())
	require.NoError(t, thrirdAction.Validate())
	require.NoError(t, fourthAction.Validate())

	elevator := models.NewElevator()

	elevator.ActionQueue = []models.Action{*firstAction, *secondAction, *thrirdAction, *fourthAction}
	elevator.CurrentDirection = models.Up
	elevator.CurrentFloor = 3
	elevator.State = models.Moving

	organizedActionQueue := models.Reorganize(elevator)

	require.NoError(t, elevator.Validate())

	expectQueue := []int{5, 7, 9}

	require.Equal(t, organizedActionQueue[0], expectQueue)

	expectQueue = []int{2}

	require.Equal(t, organizedActionQueue[1], expectQueue)
}
