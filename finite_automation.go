package finiteautomation

import (
	"errors"
	"fmt"
	"slices"

	"github.com/amitprajapati027/finite-automation/internal"
	"github.com/amitprajapati027/finite-automation/transition"
)

var (
	ErrStatesNotDefined       = errors.New("error Q is empty")
	ErrDuplicateState         = errors.New("error Q contains duplicate states")
	ErrInitialStateNotDefined = errors.New("error q0 is empty")
	ErrInvalidInitialState    = errors.New("error q0 not present in Q")
	ErrFinalStatesNotDefined  = errors.New("error F is empty")
	ErrInvalidFinalState      = errors.New("error F contains a state not present in Q")
	ErrInvalidDeltaState      = errors.New("error Delta contains a state not present in Q")
	ErrInvalidSigma           = errors.New("error Sigma contains an invalid value")
)

// GetFiniteAutomation returns a new finite internal.
//
// parameters:
// Q contains a finite set of states.
// example: []string{"S0", "S1", "S2"}
//
// Sigma contains the input strings.
// example: []string{"1", "0", "1", "0", "1", "0", "1", "0"}
//
// q0 is the initial state.
// It must be present in Q.
// example: "S0"
//
// F is the set of final or accepting states.
// These must be present in Q.
// example []string{"S0", "S1", "S2"}
//
// Delta contains the transition functions.
// example: []Transition{{StartState: "state1", Input: "input", ResultState: "state2"}}
func GetFiniteAutomation(Q []string, Sigma []string, q0 string, F []string, Delta []transition.Transition) (*internal.FiniteAutomation, error) {
	// Validations
	// Check if Q is empty.
	if len(Q) < 1 {
		return nil, ErrStatesNotDefined
	}

	// Check for duplicate states.
	statesMap := make(map[string]bool)
	for _, q := range Q {
		if statesMap[q] {
			return nil, fmt.Errorf("%w - %s", ErrDuplicateState, q)
		}

		statesMap[q] = true
	}

	// Check if q0 is empty.
	if q0 == "" {
		return nil, ErrInitialStateNotDefined
	}

	// Check if the initial state is contained in Q.
	if !slices.Contains(Q, q0) {
		return nil, fmt.Errorf("%w - %s", ErrInvalidInitialState, q0)
	}

	// Check if F is empty.
	if len(F) < 1 {
		return nil, ErrFinalStatesNotDefined
	}

	// Not check for duplicates in final states, as it doesn't affect the logic.
	// Check if all final states are present in Q.
	for _, f := range F {
		if !slices.Contains(Q, f) {
			return nil, fmt.Errorf("%w - %s", ErrInvalidFinalState, f)
		}
	}

	var deltaInputs []string
	for _, d := range Delta {
		// Check if all Delta states are present in Q.
		if !slices.Contains(Q, d.StartState) {
			return nil, fmt.Errorf("%w - %s", ErrInvalidDeltaState, d.StartState)
		}

		if !slices.Contains(Q, d.ResultState) {
			return nil, fmt.Errorf("%w - %s", ErrInvalidDeltaState, d.ResultState)
		}

		// Collect delta inputs for validating Sigma.
		deltaInputs = append(deltaInputs, d.Input)
	}

	// Check if all Sigma inputs are present in Delta.
	for _, s := range Sigma {
		if !slices.Contains(deltaInputs, s) {
			return nil, fmt.Errorf("%w - %s", ErrInvalidSigma, s)
		}
	}

	return internal.NewFiniteAutomation(Q, Sigma, q0, F, Delta)
}
