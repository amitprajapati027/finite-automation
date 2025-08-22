package transition

// Transition hold the properties of a transition
type Transition struct {
	// StartState is the state from which the transition starts.
	// If two transitions have same StartState and Input,
	// the newer transition is used.
	StartState string

	// Input contains the input for transitions.
	// If two transitions have same StartState and Input,
	// the newer transition is used.
	Input string

	// ResultState is the state that the input transtions the FSA into.
	ResultState string
}

// Transitions is a collection of transitions.
type Transitions []Transition

// GetInputs collects and returns all inputs from transitions.
func (ts Transitions) GetInputs() []string {
	inputs := make([]string, 0)
	inputsMap := make(map[string]bool)
	for _, t := range ts {
		if !inputsMap[t.Input] {
			inputs = append(inputs, t.Input)

			// Save inputs in a map, we don't want duplicates.
			inputsMap[t.Input] = true
		}
	}

	return inputs
}
