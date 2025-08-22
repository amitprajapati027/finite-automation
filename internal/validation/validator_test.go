package validation_test

import (
	"testing"

	"github.com/amitprajapati027/finite-automation/internal/validation"
	"github.com/amitprajapati027/finite-automation/transition"
	"github.com/stretchr/testify/assert"
)

func TestValidateAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		states := []string{"S0", "S1", "S2"}
		initialState := "S0"
		finalStates := []string{"S0", "S1", "S2"}
		transitions := transition.Transitions{
			{StartState: "S0", Input: "0", ResultState: "S0"},
			{StartState: "S0", Input: "1", ResultState: "S1"},
			{StartState: "S1", Input: "0", ResultState: "S2"},
			{StartState: "S1", Input: "1", ResultState: "S0"},
			{StartState: "S2", Input: "0", ResultState: "S1"},
			{StartState: "S2", Input: "1", ResultState: "S2"},
		}

		err := validation.ValidateAll(states, initialState, finalStates, transitions)
		assert.NoError(t, err)
	})

	t.Run("states not defined", func(t *testing.T) {
		states := []string{}
		initialState := "S0"
		finalStates := []string{"S0", "S1", "S2"}
		transitions := []transition.Transition{
			{StartState: "S0", Input: "0", ResultState: "S0"},
			{StartState: "S0", Input: "1", ResultState: "S1"},
			{StartState: "S1", Input: "0", ResultState: "S2"},
			{StartState: "S1", Input: "1", ResultState: "S0"},
			{StartState: "S2", Input: "0", ResultState: "S1"},
			{StartState: "S2", Input: "1", ResultState: "S2"},
		}

		err := validation.ValidateAll(states, initialState, finalStates, transitions)
		assert.Error(t, err)
		assert.ErrorIs(t, err, validation.ErrStatesNotDefined)
		assert.EqualError(t, err, "error automaton states is empty")
	})

	t.Run("duplicate states", func(t *testing.T) {
		states := []string{"S1", "S1", "S2"}
		initialState := "S0"
		finalStates := []string{"S0", "S1", "S2"}
		transitions := []transition.Transition{
			{StartState: "S0", Input: "0", ResultState: "S0"},
			{StartState: "S0", Input: "1", ResultState: "S1"},
			{StartState: "S1", Input: "0", ResultState: "S2"},
			{StartState: "S1", Input: "1", ResultState: "S0"},
			{StartState: "S2", Input: "0", ResultState: "S1"},
			{StartState: "S2", Input: "1", ResultState: "S2"},
		}

		err := validation.ValidateAll(states, initialState, finalStates, transitions)
		assert.Error(t, err)
		assert.ErrorIs(t, err, validation.ErrDuplicateState)
		assert.EqualError(t, err, "error automaton states contains duplicate state - S1")
	})

	t.Run("empty initialState", func(t *testing.T) {
		states := []string{"S0", "S1", "S2"}
		initialState := ""
		finalStates := []string{"S0", "S1", "S2"}
		transitions := []transition.Transition{
			{StartState: "S0", Input: "0", ResultState: "S0"},
			{StartState: "S0", Input: "1", ResultState: "S1"},
			{StartState: "S1", Input: "0", ResultState: "S2"},
			{StartState: "S1", Input: "1", ResultState: "S0"},
			{StartState: "S2", Input: "0", ResultState: "S1"},
			{StartState: "S2", Input: "1", ResultState: "S2"},
		}

		err := validation.ValidateAll(states, initialState, finalStates, transitions)
		assert.Error(t, err)
		assert.ErrorIs(t, err, validation.ErrInitialStateNotDefined)
		assert.EqualError(t, err, "error initial state is empty")
	})

	t.Run("invalid initialState", func(t *testing.T) {
		states := []string{"S0", "S1", "S2"}
		initialState := "S3"
		finalStates := []string{"S0", "S1", "S2"}
		transitions := []transition.Transition{
			{StartState: "S0", Input: "0", ResultState: "S0"},
			{StartState: "S0", Input: "1", ResultState: "S1"},
			{StartState: "S1", Input: "0", ResultState: "S2"},
			{StartState: "S1", Input: "1", ResultState: "S0"},
			{StartState: "S2", Input: "0", ResultState: "S1"},
			{StartState: "S2", Input: "1", ResultState: "S2"},
		}

		err := validation.ValidateAll(states, initialState, finalStates, transitions)
		assert.Error(t, err)
		assert.ErrorIs(t, err, validation.ErrInvalidInitialState)
		assert.EqualError(t, err, "error initial state not present in automaton states - S3")
	})

	t.Run("final states not defined", func(t *testing.T) {
		states := []string{"S0", "S1", "S2"}
		initialState := "S0"
		finalStates := []string{}
		transitions := []transition.Transition{
			{StartState: "S0", Input: "0", ResultState: "S0"},
			{StartState: "S0", Input: "1", ResultState: "S1"},
			{StartState: "S1", Input: "0", ResultState: "S2"},
			{StartState: "S1", Input: "1", ResultState: "S0"},
			{StartState: "S2", Input: "0", ResultState: "S1"},
			{StartState: "S2", Input: "1", ResultState: "S2"},
		}

		err := validation.ValidateAll(states, initialState, finalStates, transitions)
		assert.Error(t, err)
		assert.ErrorIs(t, err, validation.ErrFinalStatesNotDefined)
		assert.EqualError(t, err, "error final states is empty")
	})

	t.Run("final state not present in states", func(t *testing.T) {
		states := []string{"S0", "S1", "S2"}
		initialState := "S0"
		finalStates := []string{"S3"}
		transitions := []transition.Transition{
			{StartState: "S0", Input: "0", ResultState: "S0"},
			{StartState: "S0", Input: "1", ResultState: "S1"},
			{StartState: "S1", Input: "0", ResultState: "S2"},
			{StartState: "S1", Input: "1", ResultState: "S0"},
			{StartState: "S2", Input: "0", ResultState: "S1"},
			{StartState: "S2", Input: "1", ResultState: "S2"},
		}

		err := validation.ValidateAll(states, initialState, finalStates, transitions)
		assert.Error(t, err)
		assert.ErrorIs(t, err, validation.ErrInvalidFinalState)
		assert.EqualError(t, err, "error final states contains a state not present in automaton states - S3")
	})

	t.Run("invalid delta start state", func(t *testing.T) {
		states := []string{"S0", "S1", "S2"}
		initialState := "S0"
		finalStates := []string{"S0", "S1", "S2"}
		transitions := []transition.Transition{
			{StartState: "S3", Input: "0", ResultState: "S0"},
			{StartState: "S0", Input: "1", ResultState: "S1"},
			{StartState: "S1", Input: "0", ResultState: "S2"},
			{StartState: "S1", Input: "1", ResultState: "S0"},
			{StartState: "S2", Input: "0", ResultState: "S1"},
			{StartState: "S2", Input: "1", ResultState: "S2"},
		}

		err := validation.ValidateAll(states, initialState, finalStates, transitions)
		assert.Error(t, err)
		assert.ErrorIs(t, err, validation.ErrInvalidTransitionState)
		assert.EqualError(t, err, "error transitions contains a state not present in automaton states - S3")
	})

	t.Run("invalid delta end state", func(t *testing.T) {
		states := []string{"S0", "S1", "S2"}
		initialState := "S0"
		finalStates := []string{"S0", "S1", "S2"}
		transitions := []transition.Transition{
			{StartState: "S0", Input: "0", ResultState: "S3"},
			{StartState: "S0", Input: "1", ResultState: "S1"},
			{StartState: "S1", Input: "0", ResultState: "S2"},
			{StartState: "S1", Input: "1", ResultState: "S0"},
			{StartState: "S2", Input: "0", ResultState: "S1"},
			{StartState: "S2", Input: "1", ResultState: "S2"},
		}

		err := validation.ValidateAll(states, initialState, finalStates, transitions)
		assert.Error(t, err)
		assert.ErrorIs(t, err, validation.ErrInvalidTransitionState)
		assert.EqualError(t, err, "error transitions contains a state not present in automaton states - S3")
	})
}

func TestValidateInputs(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		inputs := []string{"0", "1", "0"}
		transitionInputs := []string{"0", "1"}

		err := validation.ValidateInputs(inputs, transitionInputs)
		assert.NoError(t, err)
	})

	t.Run("invalid inputs", func(t *testing.T) {
		inputs := []string{"0", "2"}
		transitionInputs := []string{"0", "1"}

		err := validation.ValidateInputs(inputs, transitionInputs)
		assert.Error(t, err)
		assert.ErrorIs(t, err, validation.ErrInvalidInput)
		assert.EqualError(t, err, "error input contains an invalid value - 2")
	})
}
