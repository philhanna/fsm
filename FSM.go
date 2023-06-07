package fsm

import "errors"

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// FSM is a finite-state machine
type FSM[T any] struct {

	// The set of states this FSM can have
	States []State

	// The set of events that this FSM can respond to
	Events []Event[T]

	// The function that maps a tuple of state and event to
	// the transition that happens.
	TransitionMap map[State]Transition[T]

	// The initial state
	InitialState State

	// The set of final states
	FinalStates []State
}

// ---------------------------------------------------------------------
// Constants and variables
// ---------------------------------------------------------------------

var (
	ERR_NO_EVENTS        = errors.New("No events defined")
	ERR_NO_INITIAL_STATE = errors.New("No initial state defined")
	ERR_NO_STATES        = errors.New("No states defined")
	ERR_NO_TRANSITIONS   = errors.New("No transitions defined")
)

// ---------------------------------------------------------------------
// Constructors
// ---------------------------------------------------------------------

// NewFSM[T] creates a new finite-state machine of type T that
func NewFSM[T any]() *FSM[T] {
	fsm := new(FSM[T])
	fsm.States = make([]State, 0)
	fsm.Events = make([]Event[T], 0)
	fsm.FinalStates = make([]State, 0)
	fsm.InitialState = UNKNOWN
	return fsm
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

// Run runs the finite state machine using the channel of events sent to
// it
func (fsm *FSM[T]) Run() (chan Event[T], error) {

	// Check for valid structure
	if len(fsm.Events) == 0 {
		return nil, ERR_NO_EVENTS
	}
	if fsm.InitialState == UNKNOWN {
		return nil, ERR_NO_INITIAL_STATE
	}
	if len(fsm.States) == 0 {
		return nil, ERR_NO_STATES
	}
	if len(fsm.TransitionMap) == 0 {
		return nil, ERR_NO_TRANSITIONS
	}

	// Start running
	state := fsm.InitialState
	ch := make(chan Event[T])
	go func() {
		defer close(ch)
		for {
			event := <-ch
			transition := fsm.TransitionMap[state]
			state = transition(state, event)
			if state == DONE {
				break
			}
		}
	}()

	return ch, nil
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

func (fsm *FSM[T]) AddEvent(event Event[T]) {
	fsm.Events = append(fsm.Events, event)
}

func (fsm *FSM[T]) AddState(state State) {
	fsm.States = append(fsm.States, state)
}

func (fsm *FSM[T]) AddFinalState(state State) {
	fsm.FinalStates = append(fsm.FinalStates, state)
}

func (fsm *FSM[T]) SetInitialState(state State) {
	fsm.InitialState = state
}

func (fsm *FSM[T]) SetTransition(state State, transition Transition[T]) {
	fsm.TransitionMap[state] = transition
}
