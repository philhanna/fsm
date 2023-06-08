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
type Transition[T any] func(event Event[T]) State

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
}

// ---------------------------------------------------------------------
// Constants and variables
// ---------------------------------------------------------------------
const (
	UNKNOWN State = -1
	ON            = true
	OFF           = false
)

var (
	ERR_NO_EVENTS        = errors.New("No events defined")
	ERR_NO_INITIAL_STATE = errors.New("No initial state defined")
	ERR_NO_STATES        = errors.New("No states defined")
	ERR_NO_TRANSITIONS   = errors.New("No transitions defined")
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
func (fsm *FSM[T]) Run(inch chan Event[T]) chan State {

	// Check for valid structure
	if fsm.InitialState == UNKNOWN {
		log.Fatal(ERR_NO_INITIAL_STATE)
	}
	if len(fsm.States) == 0 {
		log.Fatal(ERR_NO_STATES)
	}
	if len(fsm.TransitionMap) == 0 {
		log.Fatal(ERR_NO_TRANSITIONS)
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
			fsm.CurrentState = transition(event)
			outState = fsm.CurrentState
			if err != nil {
				log.Fatal(err)
			}
			if fsm.Trace {
				log.Printf("TRACE: input state=%v, event=%v, output state=%v\n", inState, event, outState)
			}
			ouch <- fsm.CurrentState
		}
	}()
	return ouch
}

// SetTrace turns the trace flag on or off.
func (fsm *FSM[T]) SetTrace(value bool) {
	fsm.Trace = value
}
