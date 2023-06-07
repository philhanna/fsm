package fsm

import (
	"log"
	"testing"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

const (
	q0 State = iota
	q1
	q2
)

func TestFSM_DivisibleBy3(t *testing.T) {
	fsm := NewFSM()
	fsm.SetTrace(true)
	fsm.States = []State{
		q0, q1, q2,
	}

	F0 := func(event Event) (State, error) {
		switch event {
		case '1', '4', '7':
			return q1, nil
		case '2', '5', '8':
			return q2, nil
		default:
			return q0, nil
		}
	}

	F1 := func(event Event) (State, error) {
		switch event {
		case '1', '4', '7':
			return q2, nil
		case '2', '5', '8':
			return q0, nil
		default:
			return q1, nil
		}
	}

	F2 := func(event Event) (State, error) {
		switch event {
		case '1', '4', '7':
			return q0, nil
		case '2', '5', '8':
			return q1, nil
		default:
			return q2, nil
		}
	}

	fsm.TransitionMap = map[State]Transition{
		q0: F0,
		q1: F1,
		q2: F2,
	}
	fsm.InitialState = q0

	inch := make(chan Event)
	ouch := fsm.Run(inch)
	var state State
	for _, r := range "112" {
		inch <- Event(r)
		state = <- ouch
	}
	log.Printf("Final state=%d\n", state)
	
}
