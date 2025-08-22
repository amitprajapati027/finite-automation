# finite-automation
A go library to construct finite automations

## GetFiniteAutomation

### Parameters

*Q* contains a finite set of states. Example: []string{"S0", "S1", "S2"}
*Sigma* contains the input strings. Example: []string{"1", "0", "1", "0", "1", "0", "1", "0"}
*q0* is the initial state. It must be present in Q. Example: "S0"
*F* is the set of final or accepting states. These must be present in Q. Example []string{"S0", "S1", "S2"}
*Delta* contains the transition functions. The different transition functions are separated by a semi-colon.
        Example: []Transition{{StartState: "state1", Input: "input", ResultState: "state2"}}

### Example

```go
package main

import (
	finiteautomation "github.com/amitprajapati027/finite-automation"
	"github.com/amitprajapati027/finite-automation/transition"
)

func main() {
	Q := []string{"S0", "S1", "S2"}
	Sigma := []string{"0", "1"}
	q0 := "S0"
	F := []string{"S0", "S1", "S2"}
	Delta := []transition.Transition{
		{StartState: "S0", Input: "0", ResultState: "S0"},
		{StartState: "S0", Input: "1", ResultState: "S1"},
		{StartState: "S1", Input: "0", ResultState: "S2"},
		{StartState: "S1", Input: "1", ResultState: "S0"},
		{StartState: "S2", Input: "0", ResultState: "S1"},
		{StartState: "S2", Input: "1", ResultState: "S2"},
	}

	fa, err := finiteautomation.GetFiniteAutomation(Q, Sigma, q0, F, Delta)
	if err != nil {
		println(err.Error())
	}

	result, err := fa.Execute()
	if err != nil {
		println(err.Error())
	}

	println(result)
}
```

## Development

### Running tests

Run the following command to run tests

```bash
make test
```
