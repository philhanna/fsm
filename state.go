package fsm

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// State is one of the possible states in the FSM.
type State int
const (
	UNKNOWN State = iota
	INIT
)