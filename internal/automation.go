package internal

import (
	"fmt"

	"github.com/amitprajapati027/finite-automation/transition"
)

// FiniteAutomation describes a finite automation.
type FiniteAutomation struct {
	// States contains a set of all states.
	States States

	// Sigma is the input to FSA.
	Sigma []string

	// Q0 is the initial state.
	Q0 *State
}

// Execute runs the automation.
func (fa *FiniteAutomation) Execute() (string, error) {
	state := fa.Q0
	var err error
	for _, s := range fa.Sigma {
		state, err = state.Transition(s)
		if err != nil {
			return "", fmt.Errorf("error executing automation: %w", err)
		}
	}

	// Return an error if the state is not a final state.
	if !state.IsFinal() {
		return "", fmt.Errorf("state %s is not a final state", state.GetName())
	}

	return state.GetName(), nil
}

// NewFiniteAutomation creates a new FiniteAutomation object.
func NewFiniteAutomation(Q []string, Sigma []string, q0 string, F []string, Delta []transition.Transition) (*FiniteAutomation, error) {
	// Create states.
	states := make(States, len(Q))
	for i, q := range Q {
		states[i] = NewState(q)
	}

	// Set final states.
	err := states.SetFinalStates(F)
	if err != nil {
		return nil, fmt.Errorf("error setting F: %w", err)
	}

	// Set deltas.
	for _, delta := range Delta {
		states.SetDelta(delta.StartState, delta.Input, delta.ResultState)
	}

	// Get the initial state.
	initialState, err := states.Find(q0)
	if err != nil {
		return nil, fmt.Errorf("error setting q0: %w", err)
	}

	return &FiniteAutomation{
		States: states,
		Sigma:  Sigma,
		Q0:     initialState,
	}, nil
}
