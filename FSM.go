package fsm

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
// Constructors
// ---------------------------------------------------------------------

// NewFSM[T] creates a new finite-state machine of type T that
func NewFSM[T any]() *FSM[T] {
	fsm := new(FSM[T])
	fsm.States = make([]State, 0)
	fsm.Events = make([]Event[T], 0)
	fsm.InitialState = INIT
	
	fsm.States = append(fsm.States, fsm.InitialState)
	fsm.FinalStates = make([]State, 0)
	return fsm
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

