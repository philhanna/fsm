package fsm

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// FSM is a finite-state machine
type FSM struct {

	// The set of events that this FSM can respond to
	events []Event

	// The function that maps a tuple of state and event to
	// the transition that happens.
	deltaFunction func(State, Event) Transition

	// The initial state
	initialState State

	// The set of final states
	finalStates []State
}

// ---------------------------------------------------------------------
// Constructors
// ---------------------------------------------------------------------

func NewFSM() *FSM {
	fsm := new(FSM)
	fsm.events = nil
	fsm.deltaFunction = nil
	fsm.initialState = *new(State)
	fsm.finalStates = make([]State, 0)
	return fsm
}

// ---------------------------------------------------------------------
// Methods
// ---------------------------------------------------------------------

