package automaton_test

import (
	"testing"

	"github.com/amitprajapati027/finite-automation/internal/automaton"
	"github.com/stretchr/testify/assert"
)

func TestNewState(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s := automaton.NewState("s1")

		assert.NotZero(t, s)
	})
}

func TestState_GetName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s := automaton.NewState("s1")

		assert.NotEmpty(t, s)
	})
}

func TestState_SetAsFinal(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s := automaton.NewState("s1")
		s.SetAsFinal()

		assert.Equal(t, true, s.IsFinal())
	})
}

func TestState_IsFinal(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s := automaton.NewState("s1")

		assert.Equal(t, false, s.IsFinal())
		s.SetAsFinal()
		assert.Equal(t, true, s.IsFinal())
	})
}

func TestState_Transition(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s1 := automaton.NewState("s1")
		s2 := automaton.NewState("s2")
		states := automaton.States{s1, s2}
		err := states.SetDelta("s1", "1", "s2")
		assert.NoError(t, err)

		result, err := s1.Transition("1")

		assert.NoError(t, err)
		assert.Equal(t, s2, result)
	})

	t.Run("transition not found", func(t *testing.T) {
		s1 := automaton.NewState("s1")

		result, err := s1.Transition("1")

		assert.Error(t, err)
		assert.ErrorIs(t, err, automaton.ErrStateTransitionNotFound)
		assert.Nil(t, result)
	})
}
