package transition_test

import (
	"testing"

	"github.com/amitprajapati027/finite-automation/transition"
	"github.com/stretchr/testify/assert"
)

func TestTransitions_GetInputs(t *testing.T) {
	t.Run("success", func(t *testing.T) {
		ts := transition.Transitions{
			{StartState: "s1", Input: "1", ResultState: "s1"},
			{StartState: "s1", Input: "0", ResultState: "s2"},
		}

		assert.Equal(t, []string{"1", "0"}, ts.GetInputs())
	})

	t.Run("no duplicates", func(t *testing.T) {
		ts := transition.Transitions{
			{StartState: "s1", Input: "1", ResultState: "s1"},
			{StartState: "s1", Input: "0", ResultState: "s2"},
			{StartState: "s2", Input: "1", ResultState: "s2"},
		}

		assert.Equal(t, []string{"1", "0"}, ts.GetInputs())
	})
}
