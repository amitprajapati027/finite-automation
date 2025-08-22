package finiteautomation_test

import (
	"testing"

	finiteautomation "github.com/amitprajapati027/finite-automation"
	"github.com/amitprajapati027/finite-automation/transition"
	"github.com/stretchr/testify/assert"
)

func TestFiniteAutomation_GetFiniteAutomation(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		Q := []string{"S0", "S1", "S2"}
		Sigma := []string{"0", "1"}
		q0 := "S0"
		F := []string{"S0", "S1", "S2"}
		Delta := []transition.Transition{
			{StartState: "S0", Input: "0", ResultState: "S0"},
			{StartState: "S0", Input: "1", ResultState: "S1"},
			{StartState: "S1", Input: "0", ResultState: "S2"},
			{StartState: "S1", Input: "1", ResultState: "S0"},
			{StartState: "S2", Input: "0", ResultState: "S1"},
			{StartState: "S2", Input: "1", ResultState: "S2"},
		}

		fa, err := finiteautomation.GetFiniteAutomation(Q, Sigma, q0, F, Delta)
		assert.NoError(t, err)
		assert.NotNil(t, fa)

		result, err := fa.Execute()
		assert.NoError(t, err)
		assert.Equal(t, "S1", result)
	})

	t.Run("states not defined", func(t *testing.T) {
		Q := []string{}
		Sigma := []string{"0", "1"}
		q0 := "S0"
		F := []string{"S0", "S1", "S2"}
		Delta := []transition.Transition{
			{StartState: "S0", Input: "0", ResultState: "S0"},
			{StartState: "S0", Input: "1", ResultState: "S1"},
			{StartState: "S1", Input: "0", ResultState: "S2"},
			{StartState: "S1", Input: "1", ResultState: "S0"},
			{StartState: "S2", Input: "0", ResultState: "S1"},
			{StartState: "S2", Input: "1", ResultState: "S2"},
		}

		fa, err := finiteautomation.GetFiniteAutomation(Q, Sigma, q0, F, Delta)
		assert.Nil(t, fa)
		assert.Error(t, err)
		assert.ErrorIs(t, err, finiteautomation.ErrStatesNotDefined)
		assert.EqualError(t, err, "error Q is empty")
	})

	t.Run("duplicate states", func(t *testing.T) {
		Q := []string{"S1", "S1", "S2"}
		Sigma := []string{"0", "1"}
		q0 := "S0"
		F := []string{"S0", "S1", "S2"}
		Delta := []transition.Transition{
			{StartState: "S0", Input: "0", ResultState: "S0"},
			{StartState: "S0", Input: "1", ResultState: "S1"},
			{StartState: "S1", Input: "0", ResultState: "S2"},
			{StartState: "S1", Input: "1", ResultState: "S0"},
			{StartState: "S2", Input: "0", ResultState: "S1"},
			{StartState: "S2", Input: "1", ResultState: "S2"},
		}

		fa, err := finiteautomation.GetFiniteAutomation(Q, Sigma, q0, F, Delta)
		assert.Nil(t, fa)
		assert.Error(t, err)
		assert.ErrorIs(t, err, finiteautomation.ErrDuplicateState)
		assert.EqualError(t, err, "error Q contains duplicate states - S1")
	})

	t.Run("empty q0", func(t *testing.T) {
		Q := []string{"S0", "S1", "S2"}
		Sigma := []string{"0", "1"}
		q0 := ""
		F := []string{"S0", "S1", "S2"}
		Delta := []transition.Transition{
			{StartState: "S0", Input: "0", ResultState: "S0"},
			{StartState: "S0", Input: "1", ResultState: "S1"},
			{StartState: "S1", Input: "0", ResultState: "S2"},
			{StartState: "S1", Input: "1", ResultState: "S0"},
			{StartState: "S2", Input: "0", ResultState: "S1"},
			{StartState: "S2", Input: "1", ResultState: "S2"},
		}

		fa, err := finiteautomation.GetFiniteAutomation(Q, Sigma, q0, F, Delta)
		assert.Nil(t, fa)
		assert.Error(t, err)
		assert.ErrorIs(t, err, finiteautomation.ErrInitialStateNotDefined)
		assert.EqualError(t, err, "error q0 is empty")
	})

	t.Run("invalid q0", func(t *testing.T) {
		Q := []string{"S0", "S1", "S2"}
		Sigma := []string{"0", "1"}
		q0 := "S3"
		F := []string{"S0", "S1", "S2"}
		Delta := []transition.Transition{
			{StartState: "S0", Input: "0", ResultState: "S0"},
			{StartState: "S0", Input: "1", ResultState: "S1"},
			{StartState: "S1", Input: "0", ResultState: "S2"},
			{StartState: "S1", Input: "1", ResultState: "S0"},
			{StartState: "S2", Input: "0", ResultState: "S1"},
			{StartState: "S2", Input: "1", ResultState: "S2"},
		}

		fa, err := finiteautomation.GetFiniteAutomation(Q, Sigma, q0, F, Delta)
		assert.Nil(t, fa)
		assert.Error(t, err)
		assert.ErrorIs(t, err, finiteautomation.ErrInvalidInitialState)
		assert.EqualError(t, err, "error q0 not present in Q - S3")
	})

	t.Run("final states not defined", func(t *testing.T) {
		Q := []string{"S0", "S1", "S2"}
		Sigma := []string{"0", "1"}
		q0 := "S0"
		F := []string{}
		Delta := []transition.Transition{
			{StartState: "S0", Input: "0", ResultState: "S0"},
			{StartState: "S0", Input: "1", ResultState: "S1"},
			{StartState: "S1", Input: "0", ResultState: "S2"},
			{StartState: "S1", Input: "1", ResultState: "S0"},
			{StartState: "S2", Input: "0", ResultState: "S1"},
			{StartState: "S2", Input: "1", ResultState: "S2"},
		}

		fa, err := finiteautomation.GetFiniteAutomation(Q, Sigma, q0, F, Delta)
		assert.Nil(t, fa)
		assert.Error(t, err)
		assert.ErrorIs(t, err, finiteautomation.ErrFinalStatesNotDefined)
		assert.EqualError(t, err, "error F is empty")
	})

	t.Run("final state not present in Q", func(t *testing.T) {
		Q := []string{"S0", "S1", "S2"}
		Sigma := []string{"0", "1"}
		q0 := "S0"
		F := []string{"S3"}
		Delta := []transition.Transition{
			{StartState: "S0", Input: "0", ResultState: "S0"},
			{StartState: "S0", Input: "1", ResultState: "S1"},
			{StartState: "S1", Input: "0", ResultState: "S2"},
			{StartState: "S1", Input: "1", ResultState: "S0"},
			{StartState: "S2", Input: "0", ResultState: "S1"},
			{StartState: "S2", Input: "1", ResultState: "S2"},
		}

		fa, err := finiteautomation.GetFiniteAutomation(Q, Sigma, q0, F, Delta)
		assert.Nil(t, fa)
		assert.Error(t, err)
		assert.ErrorIs(t, err, finiteautomation.ErrInvalidFinalState)
		assert.EqualError(t, err, "error F contains a state not present in Q - S3")
	})

	t.Run("invalid delta start state", func(t *testing.T) {
		Q := []string{"S0", "S1", "S2"}
		Sigma := []string{"0", "1"}
		q0 := "S0"
		F := []string{"S0", "S1", "S2"}
		Delta := []transition.Transition{
			{StartState: "S3", Input: "0", ResultState: "S0"},
			{StartState: "S0", Input: "1", ResultState: "S1"},
			{StartState: "S1", Input: "0", ResultState: "S2"},
			{StartState: "S1", Input: "1", ResultState: "S0"},
			{StartState: "S2", Input: "0", ResultState: "S1"},
			{StartState: "S2", Input: "1", ResultState: "S2"},
		}

		fa, err := finiteautomation.GetFiniteAutomation(Q, Sigma, q0, F, Delta)
		assert.Nil(t, fa)
		assert.Error(t, err)
		assert.ErrorIs(t, err, finiteautomation.ErrInvalidDeltaState)
		assert.EqualError(t, err, "error Delta contains a state not present in Q - S3")
	})

	t.Run("invalid delta start state", func(t *testing.T) {
		Q := []string{"S0", "S1", "S2"}
		Sigma := []string{"0", "1"}
		q0 := "S0"
		F := []string{"S0", "S1", "S2"}
		Delta := []transition.Transition{
			{StartState: "S0", Input: "0", ResultState: "S3"},
			{StartState: "S0", Input: "1", ResultState: "S1"},
			{StartState: "S1", Input: "0", ResultState: "S2"},
			{StartState: "S1", Input: "1", ResultState: "S0"},
			{StartState: "S2", Input: "0", ResultState: "S1"},
			{StartState: "S2", Input: "1", ResultState: "S2"},
		}

		fa, err := finiteautomation.GetFiniteAutomation(Q, Sigma, q0, F, Delta)
		assert.Nil(t, fa)
		assert.Error(t, err)
		assert.ErrorIs(t, err, finiteautomation.ErrInvalidDeltaState)
		assert.EqualError(t, err, "error Delta contains a state not present in Q - S3")
	})

	t.Run("invalid sigma", func(t *testing.T) {
		Q := []string{"S0", "S1", "S2"}
		Sigma := []string{"0", "2"}
		q0 := "S0"
		F := []string{"S0", "S1", "S2"}
		Delta := []transition.Transition{
			{StartState: "S0", Input: "0", ResultState: "S0"},
			{StartState: "S0", Input: "1", ResultState: "S1"},
			{StartState: "S1", Input: "0", ResultState: "S2"},
			{StartState: "S1", Input: "1", ResultState: "S0"},
			{StartState: "S2", Input: "0", ResultState: "S1"},
			{StartState: "S2", Input: "1", ResultState: "S2"},
		}

		fa, err := finiteautomation.GetFiniteAutomation(Q, Sigma, q0, F, Delta)
		assert.Nil(t, fa)
		assert.Error(t, err)
		assert.ErrorIs(t, err, finiteautomation.ErrInvalidSigma)
		assert.EqualError(t, err, "error Sigma contains an invalid value - 2")
	})
}
