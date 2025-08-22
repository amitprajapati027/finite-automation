package builder_test

import (
	"testing"

	"github.com/amitprajapati027/finite-automation/builder"
	"github.com/amitprajapati027/finite-automation/internal/automaton"
	"github.com/amitprajapati027/finite-automation/transition"
	"github.com/stretchr/testify/assert"
)

func TestNewAutomationBuilder(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ab := builder.NewAutomatonBuilder()
		assert.NotEmpty(t, ab)
	})
}

func TestAutomationBuilder_States(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ab := builder.NewAutomatonBuilder()
		abWithStates := ab.States("s1", "s2")

		assert.Equal(t, ab, abWithStates)
	})
}

func TestAutomationBuilder_AddState(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ab := builder.NewAutomatonBuilder()
		abWithStates := ab.AddState("s1")

		assert.Equal(t, ab, abWithStates)
	})
}

func TestAutomationBuilder_InitialState(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ab := builder.NewAutomatonBuilder()
		abWithInitialState := ab.InitialState("s1")

		assert.Equal(t, ab, abWithInitialState)
	})
}

func TestAutomationBuilder_FinalStates(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ab := builder.NewAutomatonBuilder()
		abWithFinalStates := ab.FinalStates("s1", "s2")

		assert.Equal(t, ab, abWithFinalStates)
	})
}

func TestAutomationBuilder_AddFinalState(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ab := builder.NewAutomatonBuilder()
		abWithFinalStates := ab.AddFinalState("s1")

		assert.Equal(t, ab, abWithFinalStates)
	})
}

func TestAutomationBuilder_Transitions(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ab := builder.NewAutomatonBuilder()
		abWithTransitions := ab.Transitions(transition.Transition{
			StartState:  "s1",
			Input:       "1",
			ResultState: "s2",
		}, transition.Transition{
			StartState:  "s1",
			Input:       "0",
			ResultState: "s2",
		})

		assert.Equal(t, ab, abWithTransitions)
	})
}

func TestAutomationBuilder_AddTransition(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ab := builder.NewAutomatonBuilder()
		abWithTransitions := ab.AddTransition(transition.Transition{
			StartState:  "s1",
			Input:       "1",
			ResultState: "s2",
		})

		assert.Equal(t, ab, abWithTransitions)
	})
}

func TestAutomationBuilder_ValidateAll(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ab := builder.
			NewAutomatonBuilder().
			States("s1", "s2").
			InitialState("s1").
			FinalStates("s1", "s2").
			AddTransition(transition.Transition{
				StartState:  "s1",
				Input:       "1",
				ResultState: "s2",
			})

		assert.NoError(t, ab.Validate())
	})

	t.Run("validation returns error", func(t *testing.T) {
		ab := builder.
			NewAutomatonBuilder().
			States("s1", "s2").
			InitialState("s1").
			FinalStates("s1", "s2")

		assert.Error(t, ab.Validate())
	})
}

func TestAutomationBuilder_Reset(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ab := builder.
			NewAutomatonBuilder().
			States("s1", "s2")

		assert.NotEqual(t, ab, ab.Reset())
	})
}

func TestAutomationBuilder_Build(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ab := builder.
			NewAutomatonBuilder().
			States("s1", "s2").
			InitialState("s1").
			FinalStates("s1", "s2").
			AddTransition(transition.Transition{
				StartState:  "s1",
				Input:       "1",
				ResultState: "s2",
			})

		a, err := ab.Build()
		assert.NoError(t, err)

		expected, err := automaton.NewFiniteAutomation([]string{"s1", "s2"}, "s1", []string{"s1", "s2"}, transition.Transitions{{
			StartState:  "s1",
			Input:       "1",
			ResultState: "s2",
		}})
		assert.NoError(t, err)

		assert.Equal(t, expected, a)
	})

	t.Run("validation error", func(t *testing.T) {
		ab := builder.
			NewAutomatonBuilder().
			States("s1", "s2").
			InitialState("s1").
			FinalStates("s1", "s2")

		a, err := ab.Build()
		assert.Error(t, err)
		assert.Nil(t, a)
	})
}
