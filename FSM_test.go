package fsm

import (
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
		case '0', '3', '6', '9':
			return q1, nil
		case '2', '5', '8':
			return q0, nil
		default:
			return q2, nil
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
}
