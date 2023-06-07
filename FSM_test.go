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
	fsm := NewFSM[int]()
	fsm.SetTrace(true)
	fsm.States = []State{
		q0, q1, q2,
	}

	fsm.TransitionMap = map[State]Transition[int]{
		q0: F0,
		q1: F1,
		q2: F2,
	}
	fsm.InitialState = q0

	inch := make(chan Event[int])
	ouch := fsm.Run(inch)
	var state State
	for _, r := range "333" {
		inch <- Event[int](r)
		state = <-ouch
	}
	log.Printf("Final state=%d\n", state)

}

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

func F1(event Event[int]) (State, error) {
	switch event {
	case '1', '4', '7':
		return q2, nil
	case '2', '5', '8':
		return q0, nil
	default:
		return q1, nil
	}
}

func F2(event Event[int]) (State, error) {
	switch event {
	case '1', '4', '7':
		return q0, nil
	case '2', '5', '8':
		return q1, nil
	default:
		return q2, nil
	}
}
