package automaton

import (
	"errors"
)

var (
	ErrStateNotFound           = errors.New("error state not found")
	ErrStateTransitionNotFound = errors.New("error state transition not found")
)

// State is a node in the FSA.
type State struct {
	// name of the state.
	name string

	// final describes if the state is a final state.
	final bool

	// delta contains transition information.
	delta map[string]*State
}

// NewState constructs and returns a new state
func NewState(name string) *State {
	return &State{
		name:  name,
		delta: make(map[string]*State),
	}
}

// GetName returns the state name.
func (s *State) GetName() string {
	return s.name
}

// SetAsFinal sets the state as a final state.
func (s *State) SetAsFinal() {
	s.final = true
}

// IsFinal returns true if the state is a final state.
func (s *State) IsFinal() bool {
	return s.final
}

// Transition uses the delta field and returns the next state.
func (s *State) Transition(sigma string) (*State, error) {
	newState := s.delta[sigma]
	if newState == nil {
		return nil, ErrStateTransitionNotFound
	}
	return newState, nil
}
