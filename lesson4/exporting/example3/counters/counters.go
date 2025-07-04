// package counters provides alert counter support.
package counters

// alertCounter is an unexported named type that contains an integer counter for alerts.
type alertCounter int

// New creates and returns values of the unexported type alertCounter.
func New(val int) alertCounter {
	return alertCounter(val)
}
