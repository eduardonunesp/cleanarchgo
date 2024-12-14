package valueobject

import (
	"fmt"
)

// Status type to ensure type safety
type Status string

// Action type for transitions
type Action string

// Typed constants for ride statuses
const (
	StatusRequested  Status = "requested"
	StatusAccepted   Status = "accepted"
	StatusInProgress Status = "in_progress"
	StatusFinished   Status = "completed"
)

// Typed constants for actions
const (
	ActionAccept Action = "accept"
	ActionStart  Action = "start"
	ActionFinish Action = "finish"
)

// RideStatus interface
type RideStatus interface {
	fmt.Stringer
	Transition(action Action) (RideStatus, error)
	Is(status Status) bool
}

// Transition map defining valid transitions
var rideStatusTransitions = map[Status]map[Action]Status{
	StatusRequested: {
		ActionAccept: StatusAccepted,
	},
	StatusAccepted: {
		ActionStart: StatusInProgress,
	},
	StatusInProgress: {
		ActionFinish: StatusFinished,
	},
	StatusFinished: {},
}

// RideStatus implementation
type rideStatus struct {
	name Status
}

func (r rideStatus) Is(status Status) bool {
	return r.name == status
}

func (r rideStatus) String() string {
	return string(r.name)
}

func (r rideStatus) Transition(action Action) (RideStatus, error) {
	nextStatus, ok := rideStatusTransitions[r.name][action]
	if !ok {
		return nil, fmt.Errorf("invalid action '%s' for status '%s'", action, r.name)
	}
	return rideStatus{name: nextStatus}, nil
}

// Global map for status lookup
var rideStatusStates = map[Status]RideStatus{
	StatusRequested:  rideStatus{name: StatusRequested},
	StatusAccepted:   rideStatus{name: StatusAccepted},
	StatusInProgress: rideStatus{name: StatusInProgress},
	StatusFinished:   rideStatus{name: StatusFinished},
}

// BuildRideStatus creates a RideStatus from a string
func BuildRideStatus(status Status) (RideStatus, error) {
	r, ok := rideStatusStates[status]
	if !ok {
		return nil, fmt.Errorf("invalid ride status: '%s'", status)
	}
	return r, nil
}

// RideStatusAs checks if two RideStatuses are equivalent
func RideStatusAs(status, target RideStatus) bool {
	if status == nil || target == nil {
		return false
	}
	return status.String() == target.String()
}
