package builder

import (
	"github.com/amitprajapati027/finite-automation/internal/automaton"
	"github.com/amitprajapati027/finite-automation/internal/validation"
	"github.com/amitprajapati027/finite-automation/transition"
)

// AutomatonBuilder provides an interface for constructing finite automata.
type AutomatonBuilder struct {
	states       []string
	initialState string
	finalStates  []string
	transitions  transition.Transitions
}

// NewAutomatonBuilder creates a new AutomatonBuilder.
func NewAutomatonBuilder() *AutomatonBuilder {
	return &AutomatonBuilder{
		states:      make([]string, 0),
		finalStates: make([]string, 0),
		transitions: make(transition.Transitions, 0),
	}
}

// States sets the states of the automaton.
func (b *AutomatonBuilder) States(states ...string) *AutomatonBuilder {
	b.states = states
	return b
}

// AddState adds a single state to the automaton.
func (b *AutomatonBuilder) AddState(state string) *AutomatonBuilder {
	b.states = append(b.states, state)
	return b
}

// InitialState sets the initial state of the automaton.
func (b *AutomatonBuilder) InitialState(state string) *AutomatonBuilder {
	b.initialState = state
	return b
}

// FinalStates sets the final states of the automaton.
func (b *AutomatonBuilder) FinalStates(states ...string) *AutomatonBuilder {
	b.finalStates = states
	return b
}

// AddFinalState adds a single final state.
func (b *AutomatonBuilder) AddFinalState(state string) *AutomatonBuilder {
	b.finalStates = append(b.finalStates, state)
	return b
}

// Transitions sets the transitions of the automaton.
func (b *AutomatonBuilder) Transitions(transitions ...transition.Transition) *AutomatonBuilder {
	b.transitions = transitions
	return b
}

// AddTransition adds a single transition.
func (b *AutomatonBuilder) AddTransition(transition transition.Transition) *AutomatonBuilder {
	b.transitions = append(b.transitions, transition)
	return b
}

// Validate validates the current configuration.
func (b *AutomatonBuilder) Validate() error {
	return validation.ValidateAll(b.states, b.initialState, b.finalStates, b.transitions)
}

// Reset clears all configuration and returns a fresh builder.
func (b *AutomatonBuilder) Reset() *AutomatonBuilder {
	return NewAutomatonBuilder()
}

// Build constructs and returns the finite automaton.
func (b *AutomatonBuilder) Build() (*automaton.FiniteAutomation, error) {
	// Validate before building
	err := b.Validate()
	if err != nil {
		return nil, err
	}

	return automaton.NewFiniteAutomation(b.states, b.initialState, b.finalStates, b.transitions)
}
