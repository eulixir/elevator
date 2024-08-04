package models_test

import (
	"elevator/models"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateIfActionIsEmpty(t *testing.T) {
	action := models.NewAction()
	err := action.Validate()

	require.Error(t, err)
}

func TestValidateAction(t *testing.T) {
	firstAction := models.NewAction()

	firstAction.Direction = models.Down
	firstAction.Floor = 5

	secondAction := models.NewAction()

	secondAction.Direction = models.Down
	secondAction.Floor = 2

	require.NoError(t, firstAction.Validate())
	require.NoError(t, secondAction.Validate())
}
