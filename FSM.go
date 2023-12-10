package fsm

import (
	"errors"
	"log"
)

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

type Event[T any] any
type State int
type Transition[T any] func(event Event[T]) (State, error)

// FSM is a finite-state machine
type FSM[T any] struct {

	// The set of states this FSM can have
	States []State

	// The initial state
	InitialState State

	// CurrentState state
	CurrentState State

	// The function that maps a tuple of state and event to
	// the transition that happens.
	TransitionMap map[State]Transition[T]

	// Set on the Trace flag to trace each transition
	Trace bool

	// Current error
	Error error
}

// ---------------------------------------------------------------------
// Constants and variables
// ---------------------------------------------------------------------
const (
	UNKNOWN State = -1
	ERROR   State = -86
	ON            = true
	OFF           = false
)

var (
	ErrNoEvents       = errors.New("no events defined")
	ErrNoInitialState = errors.New("no initial state defined")
	ErrNoStates       = errors.New("no states defined")
	ErrNoTransitions  = errors.New("no transitions defined")
)

// ---------------------------------------------------------------------
// Constructors
// ---------------------------------------------------------------------

// NewFSM creates a new finite-state machine
func NewFSM[T any]() *FSM[T] {
	fsm := new(FSM[T])
	fsm.States = make([]State, 0)
	fsm.InitialState = UNKNOWN
	return fsm
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Run runs the finite state machine, sending the state back by a channel
func (fsm *FSM[T]) Run(inch chan Event[T]) (chan State, error) {

	// Check for valid structure
	if fsm.InitialState == UNKNOWN {
		return nil, ErrNoInitialState
	}
	if len(fsm.States) == 0 {
		return nil, ErrNoStates
	}
	if len(fsm.TransitionMap) == 0 {
		return nil, ErrNoTransitions
	}

	// Start running
	fsm.CurrentState = fsm.InitialState
	ouch := make(chan State)
	go func() {
		var err error
		var inState, outState State
		defer close(ouch)
		for {
			inState = fsm.CurrentState
			event := <-inch
			transition := fsm.TransitionMap[fsm.CurrentState]
			fsm.CurrentState, err = transition(event)
			outState = fsm.CurrentState
			if err != nil {
				fsm.Error = err
			}
			if fsm.Trace {
				log.Printf("TRACE: input state=%v, event=%v, output state=%v\n", inState, event, outState)
			}
			ouch <- fsm.CurrentState
		}
	}()
	return ouch, nil
}

// SetTrace turns the trace flag on or off.
func (fsm *FSM[T]) SetTrace(value bool) {
	fsm.Trace = value
}
