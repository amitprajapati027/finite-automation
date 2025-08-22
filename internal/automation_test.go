package internal_test

import (
	"testing"

	"github.com/amitprajapati027/finite-automation/internal"
	"github.com/amitprajapati027/finite-automation/transition"
	"github.com/stretchr/testify/assert"
)

func TestFiniteAutomation_Execute(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		sigma := []string{"1", "0", "1"}
		state1 := internal.NewState("s1")
		state2 := internal.NewState("s2")
		state2.SetAsFinal()

		states := internal.States{state1, state2}
		states.SetDelta("s1", "1", "s1")
		states.SetDelta("s1", "0", "s2")
		states.SetDelta("s2", "1", "s2")

		fa := &internal.FiniteAutomation{
			States: states,
			Sigma:  sigma,
			Q0:     state1,
		}

		result, err := fa.Execute()
		assert.NoError(t, err)
		assert.Equal(t, "s2", result)
	})

	t.Run("state transition returns error", func(t *testing.T) {
		sigma := []string{"1", "0", "1"}
		state1 := internal.NewState("s1")
		state2 := internal.NewState("s2")
		state2.SetAsFinal()

		states := internal.States{state1, state2}

		fa := &internal.FiniteAutomation{
			States: states,
			Sigma:  sigma,
			Q0:     state1,
		}

		result, err := fa.Execute()
		assert.Error(t, err)
		assert.EqualError(t, err, "error executing automation: error state transition not found")
		assert.Zero(t, result)
	})

	t.Run("state is not final", func(t *testing.T) {
		sigma := []string{"1", "0", "1"}
		state1 := internal.NewState("s1")
		state2 := internal.NewState("s2")

		states := internal.States{state1, state2}
		states.SetDelta("s1", "1", "s1")
		states.SetDelta("s1", "0", "s2")
		states.SetDelta("s2", "1", "s2")

		fa := &internal.FiniteAutomation{
			States: states,
			Sigma:  sigma,
			Q0:     state1,
		}

		result, err := fa.Execute()
		assert.Error(t, err)
		assert.EqualError(t, err, "state s2 is not a final state")
		assert.Zero(t, result)
	})
}

func TestNewFiniteAutomation(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		Q := []string{"s1", "s2"}
		Sigma := []string{"1", "0", "1"}
		q0 := "s1"
		F := []string{"s2"}
		Delta := []transition.Transition{
			{StartState: "s1", Input: "1", ResultState: "s1"},
			{StartState: "s1", Input: "0", ResultState: "s2"},
			{StartState: "s2", Input: "1", ResultState: "s2"},
		}

		fa, err := internal.NewFiniteAutomation(Q, Sigma, q0, F, Delta)
		assert.NoError(t, err)
		assert.NotZero(t, fa)
	})

	t.Run("error setting F", func(t *testing.T) {
		Q := []string{"s1", "s2"}
		Sigma := []string{"1", "0", "1"}
		q0 := "s1"
		F := []string{"s3"}
		Delta := []transition.Transition{
			{StartState: "s1", Input: "1", ResultState: "s1"},
			{StartState: "s1", Input: "0", ResultState: "s2"},
			{StartState: "s2", Input: "1", ResultState: "s2"},
		}

		fa, err := internal.NewFiniteAutomation(Q, Sigma, q0, F, Delta)
		assert.Error(t, err)
		assert.EqualError(t, err, "error setting F: error state not found")
		assert.Zero(t, fa)
	})

	t.Run("error setting q0", func(t *testing.T) {
		Q := []string{"s1", "s2"}
		Sigma := []string{"1", "0", "1"}
		q0 := "s3"
		F := []string{"s2"}
		Delta := []transition.Transition{
			{StartState: "s1", Input: "1", ResultState: "s1"},
			{StartState: "s1", Input: "0", ResultState: "s2"},
			{StartState: "s2", Input: "1", ResultState: "s2"},
		}

		fa, err := internal.NewFiniteAutomation(Q, Sigma, q0, F, Delta)
		assert.Error(t, err)
		assert.EqualError(t, err, "error setting q0: error state not found")
		assert.Zero(t, fa)
	})
}
