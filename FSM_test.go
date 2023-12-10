package fsm

import (
	"bytes"
	"fmt"
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

// ---------------------------------------------------------------------
// Transition functions for testing
// ---------------------------------------------------------------------

func F0(event Event[int]) (State, error) {
	switch event {
	case '1', '4', '7':
		return q1, nil
	case '2', '5', '8':
		return q2, nil
	case '0', '3', '9':
		return q0, nil
	default:
		return ERROR, fmt.Errorf("Unrecognized event %v", event)
	}
}

func F1(event Event[int]) (State, error) {
	switch event {
	case '1', '4', '7':
		return q2, nil
	case '2', '5', '8':
		return q0, nil
	case '0', '3', '9':
		return q1, nil
	default:
		return ERROR, fmt.Errorf("Unrecognized event %v", event)
	}
}

func F2(event Event[int]) (State, error) {
	switch event {
	case '1', '4', '7':
		return q0, nil
	case '2', '5', '8':
		return q1, nil
	case '0', '3', '9':
		return q2, nil
	default:
		return ERROR, fmt.Errorf("Unrecognized event %v", event)

	}
}

// ---------------------------------------------------------------------
// Unit tests
// ---------------------------------------------------------------------

func TestBadInitialState(t *testing.T) {
	fsm := NewFSM[int]()
	inch := make(chan Event[int])
	defer close(inch)
	_, err := fsm.Run(inch)
	assert.NotNil(t, err)
}

func TestBadStateList(t *testing.T) {
	fsm := FSM[int]{
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
	_, err := fsm.Run(inch)
	assert.NotNil(t, err)
}

func TestBadTransitionMap(t *testing.T) {
	fsm := FSM[int]{
		States:       []State{q0, q1, q2},
		InitialState: q0,
		Trace:        OFF,
	}
	inch := make(chan Event[int])
	defer close(inch)
	_, err := fsm.Run(inch)
	assert.NotNil(t, err)
}

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

			ouch, err := fsm.Run(inch)
			assert.Nil(t, err)

			var state State
			for _, r := range input {
				inch <- Event[int](r)
				state = <-ouch
			}

			assert.Equal(t, want, state)
		})
	}
}

func TestNewFSM(t *testing.T) {
	machine := NewFSM[int]()
	assert.NotNil(t, machine)
	assert.Equal(t, 0, len(machine.States))
	assert.Equal(t, UNKNOWN, machine.InitialState)
	assert.Equal(t, State(0), machine.CurrentState)
	assert.Equal(t, 0, len(machine.TransitionMap))
	assert.Equal(t, OFF, machine.Trace)
}

func TestSetTrace(t *testing.T) {
	fsm := FSM[int]{
		States:       []State{q0, q1, q2},
		InitialState: q0,
		TransitionMap: map[State]Transition[int]{
			q0: F0,
			q1: F1,
			q2: F2,
		},
	}
	assert.Equal(t, OFF, fsm.Trace)
	fsm.SetTrace(ON)
	assert.Equal(t, ON, fsm.Trace)
}

func TestTraceLog(t *testing.T) {
	logbuf := bytes.Buffer{}
	log.SetOutput(&logbuf)
	fsm := FSM[int]{
		States:       []State{q0, q1, q2},
		InitialState: q0,
		TransitionMap: map[State]Transition[int]{
			q0: F0,
			q1: F1,
			q2: F2,
		},
	}
	fsm.Trace = ON

	inch := make(chan Event[int])
	defer close(inch)

	ouch, err := fsm.Run(inch)
	assert.Nil(t, err)

	var state State
	input := "241"
	for _, r := range input {
		inch <- Event[int](r)
		state = <-ouch
	}
	want := State(1)
	assert.Equal(t, want, state)

	logString := logbuf.String()
	assert.Contains(t, logString, "event=50")
	assert.Contains(t, logString, "event=52")
	assert.Contains(t, logString, "event=49")
}
