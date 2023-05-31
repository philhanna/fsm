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
