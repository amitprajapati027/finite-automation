package automaton

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
