package automaton

import (
	"fmt"

	"github.com/amitprajapati027/finite-automation/internal/validation"
	"github.com/amitprajapati027/finite-automation/transition"
)

// FiniteAutomation describes a finite automation.
type FiniteAutomation struct {
	// States contains a set of all states.
	States States

	// TransitionInputs contains all valid inputs to FSA.
	TransitionInputs []string

	// InitialState is the initial state.
	InitialState *State
}

// Execute runs the automation.
func (fa *FiniteAutomation) Execute(Sigma ...string) (string, error) {
	err := validation.ValidateInputs(Sigma, fa.TransitionInputs)
	if err != nil {
		return "", fmt.Errorf("failed to execute finite automation: %w", err)
	}

	state := fa.InitialState
	for _, s := range Sigma {
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
func NewFiniteAutomation(Q []string, q0 string, F []string, Delta transition.Transitions) (*FiniteAutomation, error) {
	// Create states.
	states := make(States, len(Q))
	for i, q := range Q {
		states[i] = NewState(q)
	}

	// Set final states.
	err := states.SetFinalStates(F)
	if err != nil {
		return nil, fmt.Errorf("error setting final states: %w", err)
	}

	// Set deltas.
	for _, delta := range Delta {
		states.SetDelta(delta.StartState, delta.Input, delta.ResultState)
	}

	// Get the initial state.
	initialState, err := states.Find(q0)
	if err != nil {
		return nil, fmt.Errorf("error setting initial state: %w", err)
	}

	return &FiniteAutomation{
		States:           states,
		TransitionInputs: Delta.GetInputs(),
		InitialState:     initialState,
	}, nil
}
