# finite-automation
A go library to construct finite automations

## AutomationBuilder

### Parameters

*Q* is a finite set of states. Example: []string{"S0", "S1", "S2"}
*Sigma* is the input strings. Example: []string{"1", "0", "1", "0", "1", "0", "1", "0"}
*q0* is the initial state. It must be present in Q. Example: "S0"
*F* is the set of final or accepting states. These must be present in Q. Example []string{"S0", "S1", "S2"}
*Delta* are the transition functions. The different transition functions are separated by a semi-colon.
        Example: []Transition{{StartState: "state1", Input: "input", ResultState: "state2"}}

### Example

```go
package main

import (
	"github.com/amitprajapati027/finite-automation/builder"
	"github.com/amitprajapati027/finite-automation/transition"
)

func main() {
	automatonBuilder := builder.
		NewAutomatonBuilder().
		States("S0", "S1", "S2").
		InitialState("S0").
		FinalStates("S0", "S1", "S2").
		Transitions(
			transition.Transition{StartState: "S0", Input: "0", ResultState: "S0"},
			transition.Transition{StartState: "S0", Input: "1", ResultState: "S1"},
			transition.Transition{StartState: "S1", Input: "0", ResultState: "S2"},
			transition.Transition{StartState: "S1", Input: "1", ResultState: "S0"},
			transition.Transition{StartState: "S2", Input: "0", ResultState: "S1"},
			transition.Transition{StartState: "S2", Input: "1", ResultState: "S2"},
		)

	modulo3, err := automatonBuilder.Build()
	if err != nil {
		println(err.Error())
		return
	}

	result, err := modulo3.Execute("0", "1")
	if err != nil {
		println(err.Error())
		return
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
