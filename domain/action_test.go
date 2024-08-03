package domain_test

import (
	"elevator/domain"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestValidateIfActionIsEmpty(t *testing.T) {
	action := domain.NewAction()
	err := action.Validate()

	require.Error(t, err)
}

func TestValidateAction(t *testing.T) {
	firstAction := domain.NewAction()

	firstAction.Direction = domain.Down
	firstAction.Floor = 5

	secondAction := domain.NewAction()

	secondAction.Direction = domain.Down
	secondAction.Floor = 2

	require.NoError(t, firstAction.Validate())
	require.NoError(t, secondAction.Validate())
}
