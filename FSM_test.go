package fsm

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"
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

	tests := []struct {
		name  string
		input string
		want  int
	}{
		{"243", "243", 0},
		{"3715", "3715", 1},
		{"34", "34", 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			log.Printf("Starting new test %v\n", tt.name)
			input := tt.input
			want := State(tt.want)

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

			assert.Equal(t, want, state)
		})
	}

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
