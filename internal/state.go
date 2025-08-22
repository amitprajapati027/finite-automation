package internal

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

// States is a collection of state.
type States []*State

// SetFinalStates sets the final field for states defined in F.
func (s States) SetFinalStates(F []string) error {
	for _, f := range F {
		stateFound := false
		for _, state := range s {
			if state.name == f {
				stateFound = true
				state.SetAsFinal()
			}
		}

		if !stateFound {
			return ErrStateNotFound
		}
	}

	return nil
}

// SetDelta sets the delta field in all states in the collection.
func (s States) SetDelta(start, sigma, end string) error {
	startState, err := s.Find(start)
	if err != nil {
		return err
	}

	endState, err := s.Find(end)
	if err != nil {
		return err
	}

	startState.delta[sigma] = endState

	return nil
}

// Find finds a state by it's name and returns it.
func (s States) Find(name string) (*State, error) {
	for _, state := range s {
		if state.name == name {
			return state, nil
		}
	}

	return nil, ErrStateNotFound
}
