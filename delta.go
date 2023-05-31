package fsm

// ---------------------------------------------------------------------
// Type definitions
// ---------------------------------------------------------------------

// Delta is a tuple of a state and an event, which operate as a map
// to the transition and next state
type Delta struct {
	state State
	event Event
}

// ---------------------------------------------------------------------
// Constructor
// ---------------------------------------------------------------------

func NewDelta(state State, event Event) *Delta {
	d := new(Delta)
	d.state = state
	d.event = event
	return d
}
