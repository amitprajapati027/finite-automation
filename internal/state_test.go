package internal_test

import (
	"testing"

	"github.com/amitprajapati027/finite-automation/internal"
	"github.com/stretchr/testify/assert"
)

func TestNewState(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s := internal.NewState("s1")

		assert.NotZero(t, s)
	})
}

func TestState_GetName(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s := internal.NewState("s1")

		assert.NotEmpty(t, s)
	})
}

func TestState_SetAsFinal(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s := internal.NewState("s1")
		s.SetAsFinal()

		assert.Equal(t, true, s.IsFinal())
	})
}

func TestState_IsFinal(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s := internal.NewState("s1")

		assert.Equal(t, false, s.IsFinal())
		s.SetAsFinal()
		assert.Equal(t, true, s.IsFinal())
	})
}

func TestState_Transition(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s1 := internal.NewState("s1")
		s2 := internal.NewState("s2")
		states := internal.States{s1, s2}
		err := states.SetDelta("s1", "1", "s2")
		assert.NoError(t, err)

		result, err := s1.Transition("1")

		assert.NoError(t, err)
		assert.Equal(t, s2, result)
	})

	t.Run("transition not found", func(t *testing.T) {
		s1 := internal.NewState("s1")

		result, err := s1.Transition("1")

		assert.Error(t, err)
		assert.ErrorIs(t, err, internal.ErrStateTransitionNotFound)
		assert.Nil(t, result)
	})
}

func TestStates_SetFinalStates(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s1 := internal.NewState("s1")
		s2 := internal.NewState("s2")
		states := internal.States{s1, s2}

		err := states.SetFinalStates([]string{"s2"})

		assert.NoError(t, err)
	})

	t.Run("state not found", func(t *testing.T) {
		s1 := internal.NewState("s1")
		s2 := internal.NewState("s2")
		states := internal.States{s1, s2}

		err := states.SetFinalStates([]string{"s3"})

		assert.Error(t, err)
		assert.ErrorIs(t, err, internal.ErrStateNotFound)
	})
}

func TestStates_SetDelta(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s1 := internal.NewState("s1")
		s2 := internal.NewState("s2")
		states := internal.States{s1, s2}

		err := states.SetDelta("s1", "1", "s2")

		assert.NoError(t, err)
	})

	t.Run("start state not found", func(t *testing.T) {
		s1 := internal.NewState("s1")
		s2 := internal.NewState("s2")
		states := internal.States{s1, s2}

		err := states.SetDelta("s3", "1", "s2")

		assert.Error(t, err)
		assert.ErrorIs(t, err, internal.ErrStateNotFound)
	})

	t.Run("end state not found", func(t *testing.T) {
		s1 := internal.NewState("s1")
		s2 := internal.NewState("s2")
		states := internal.States{s1, s2}

		err := states.SetDelta("s1", "1", "s3")

		assert.Error(t, err)
		assert.ErrorIs(t, err, internal.ErrStateNotFound)
	})
}

func TestStates_Find(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		s1 := internal.NewState("s1")
		s2 := internal.NewState("s2")
		states := internal.States{s1, s2}

		s, err := states.Find("s1")

		assert.Equal(t, s1, s)
		assert.NoError(t, err)
	})

	t.Run("state not found", func(t *testing.T) {
		s1 := internal.NewState("s1")
		s2 := internal.NewState("s2")
		states := internal.States{s1, s2}

		s, err := states.Find("s3")

		assert.Nil(t, s)
		assert.Error(t, err)
		assert.ErrorIs(t, err, internal.ErrStateNotFound)
	})
}
