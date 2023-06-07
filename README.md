# fsm
[![Go Report Card](https://goreportcard.com/badge/github.com/philhanna/fsm)][idGoReportCard]
[![PkgGoDev](https://pkg.go.dev/badge/github.com/philhanna/fsm)][idPkgGoDev]


A Go package for using finite state machines.  See the Wikipedia article [Finite-state machine]
for an overview.

## Usage

Set up the finite state machine and then call its `Run` method, sending it events
through a channel.  You can read the final state through the FSM's output channel.

The FSM handles generic events; you must specify the type. The example below just
uses integers as events.  The transitions between states given a particular event
you must define as functions taking an event and returning a new state.
```go
func F0(event Event[int]) (State, error) {

	switch event {
	case '1', '4', '7':
		return q1, nil
	case '2', '5', '8':
		return q2, nil
	default:
		return q0, nil
	}
}
```


```go
fsm := FSM[int]{
    States:       []State{q0, q1, q2},
    InitialState: q0,
    TransitionMap: map[State]Transition[int]{
        q0: F0,
        q1: F1,
        q2: F2,
    },
    Trace: OFF,
}

inch := make(chan Event[int])
defer close(inch)

ouch := fsm.Run(inch)

var state State
for _, r := range input {
    inch <- Event[int](r)
    state = <-ouch
}

finalState := state
```

## References
- [Github repository](https://github.com/philhanna/fsm)
- [Wikipedia article on FSM][Finite-state machine]

[Finite-state machine]:https://en.wikipedia.org/wiki/Finite-state_machine


[idGoReportCard]: https://goreportcard.com/report/github.com/philhanna/fsm
[idPkgGoDev]: https://pkg.go.dev/github.com/philhanna/fsm
