package automaton_test

import (
	"testing"

	"github.com/amitprajapati027/finite-automation/internal/automaton"
	"github.com/stretchr/testify/assert"
)

func TestStates_SetFinalStates(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s1 := automaton.NewState("s1")
		s2 := automaton.NewState("s2")
		states := automaton.States{s1, s2}

		err := states.SetFinalStates([]string{"s2"})

		assert.NoError(t, err)
	})

	t.Run("state not found", func(t *testing.T) {
		s1 := automaton.NewState("s1")
		s2 := automaton.NewState("s2")
		states := automaton.States{s1, s2}

		err := states.SetFinalStates([]string{"s3"})

		assert.Error(t, err)
		assert.ErrorIs(t, err, automaton.ErrStateNotFound)
	})
}

func TestStates_SetDelta(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s1 := automaton.NewState("s1")
		s2 := automaton.NewState("s2")
		states := automaton.States{s1, s2}

		err := states.SetDelta("s1", "1", "s2")

		assert.NoError(t, err)
	})

	t.Run("start state not found", func(t *testing.T) {
		s1 := automaton.NewState("s1")
		s2 := automaton.NewState("s2")
		states := automaton.States{s1, s2}

		err := states.SetDelta("s3", "1", "s2")

		assert.Error(t, err)
		assert.ErrorIs(t, err, automaton.ErrStateNotFound)
	})

	t.Run("end state not found", func(t *testing.T) {
		s1 := automaton.NewState("s1")
		s2 := automaton.NewState("s2")
		states := automaton.States{s1, s2}

		err := states.SetDelta("s1", "1", "s3")

		assert.Error(t, err)
		assert.ErrorIs(t, err, automaton.ErrStateNotFound)
	})
}

func TestStates_Find(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s1 := automaton.NewState("s1")
		s2 := automaton.NewState("s2")
		states := automaton.States{s1, s2}

		s, err := states.Find("s1")

		assert.Equal(t, s1, s)
		assert.NoError(t, err)
	})

	t.Run("state not found", func(t *testing.T) {
		s1 := automaton.NewState("s1")
		s2 := automaton.NewState("s2")
		states := automaton.States{s1, s2}

		s, err := states.Find("s3")

		assert.Nil(t, s)
		assert.Error(t, err)
		assert.ErrorIs(t, err, automaton.ErrStateNotFound)
	})
}
