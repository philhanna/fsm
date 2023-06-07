package fsm

import "errors"

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

type Event int
type State int
type Transition func(event Event) (State, error)

// FSM is a finite-state machine
type FSM struct {

	// The set of states this FSM can have
	States []State

	// The initial state
	InitialState State

	// The function that maps a tuple of state and event to
	// the transition that happens.
	TransitionMap map[State]Transition
}


// ---------------------------------------------------------------------
// Constants and variables
// ---------------------------------------------------------------------
const (
	UNKNOWN     State = -1
	ERROR_STATE State = -2
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
func NewFSM() *FSM {
	fsm := new(FSM)
	fsm.States = make([]State, 0)
	fsm.InitialState = UNKNOWN
	return fsm
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Run runs the finite state machine using the channel of events sent to
// it. It returns the final state and any error.
func (fsm *FSM) Run(ch <-chan Event) (State, error) {

	// Check for valid structure
	if fsm.InitialState == UNKNOWN {
		return ERROR_STATE, ERR_NO_INITIAL_STATE
	}
	if len(fsm.States) == 0 {
		return ERROR_STATE, ERR_NO_STATES
	}
	if len(fsm.TransitionMap) == 0 {
		return ERROR_STATE, ERR_NO_TRANSITIONS
	}

	// Start running
	state := fsm.InitialState
	go func() (State, error) {
		for {
			event, OK := <-ch
			if !OK {
				return state, nil
			}
			transition := fsm.TransitionMap[state]
			state, err := transition(event)
			if err != nil {
				return state, err
			}
		}
	}()

	// Return the final state
	return state, nil
}
