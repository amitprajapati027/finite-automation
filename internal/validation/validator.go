package validation

import (
	"errors"
	"fmt"
	"slices"

	"github.com/amitprajapati027/finite-automation/transition"
)

var (
	ErrStatesNotDefined       = errors.New("error automaton states is empty")
	ErrDuplicateState         = errors.New("error automaton states contains duplicate state")
	ErrInitialStateNotDefined = errors.New("error initial state is empty")
	ErrInvalidInitialState    = errors.New("error initial state not present in automaton states")
	ErrFinalStatesNotDefined  = errors.New("error final states is empty")
	ErrInvalidFinalState      = errors.New("error final states contains a state not present in automaton states")
	ErrInvalidTransitions     = errors.New("error transitions is empty")
	ErrInvalidTransitionState = errors.New("error transitions contains a state not present in automaton states")
	ErrInvalidInput           = errors.New("error input contains an invalid value")
)

// ValidateAll performs comprehensive validation of all automaton components.
func ValidateAll(states []string, initialState string, finalStates []string, transitions transition.Transitions) error {
	// Perform all validations
	err := validateStates(states)
	if err != nil {
		return err
	}

	err = validateInitialState(initialState, states)
	if err != nil {
		return err
	}

	err = validateFinalStates(finalStates, states)
	if err != nil {
		return err
	}

	err = validateTransitions(transitions, states)
	if err != nil {
		return err
	}

	return nil
}

// validateStates validates the states.
func validateStates(states []string) error {
	// Check if Q is empty.
	if len(states) < 1 {
		return ErrStatesNotDefined
	}

	// Check for duplicate states.
	statesMap := make(map[string]bool)
	for _, q := range states {
		if statesMap[q] {
			return fmt.Errorf("%w - %s", ErrDuplicateState, q)
		}

		statesMap[q] = true
	}

	return nil
}

// validateInitialState validates the initial state
func validateInitialState(initialState string, states []string) error {
	// Check if initial state is empty.
	if initialState == "" {
		return ErrInitialStateNotDefined
	}

	// Check if the initial state is contained in all states.
	if !slices.Contains(states, initialState) {
		return fmt.Errorf("%w - %s", ErrInvalidInitialState, initialState)
	}

	return nil
}

// validateFinalStates validates the final states set F
func validateFinalStates(finalStates []string, states []string) error {
	// Check if F is empty.
	if len(finalStates) < 1 {
		return ErrFinalStatesNotDefined
	}

	// Not check for duplicates in final states, as it doesn't affect the logic.
	// Check if all final states are present in Q.
	for _, f := range finalStates {
		if !slices.Contains(states, f) {
			return fmt.Errorf("%w - %s", ErrInvalidFinalState, f)
		}
	}

	return nil
}

// validateTransitions validates the transition function Delta
func validateTransitions(transitions []transition.Transition, states []string) error {
	if len(transitions) < 1 {
		return ErrInvalidTransitions
	}

	for _, d := range transitions {
		// Check if all Delta states are present in Q.
		if !slices.Contains(states, d.StartState) {
			return fmt.Errorf("%w - %s", ErrInvalidTransitionState, d.StartState)
		}

		if !slices.Contains(states, d.ResultState) {
			return fmt.Errorf("%w - %s", ErrInvalidTransitionState, d.ResultState)
		}
	}

	return nil
}

// ValidateInputs validates that all symbols in the alphabet have corresponding transitions
func ValidateInputs(inputs []string, transitionInputs []string) error {
	// Check if all Sigma inputs are present in Delta.
	for _, s := range inputs {
		if !slices.Contains(transitionInputs, s) {
			return fmt.Errorf("%w - %s", ErrInvalidInput, s)
		}
	}

	return nil
}
