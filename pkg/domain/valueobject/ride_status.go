package valueobject

import (
	"errors"
	"fmt"
)

type RideStatus interface {
	fmt.Stringer
	Requested(prevStatus RideStatus) (RideStatus, error)
	Accepted(prevStatus RideStatus) (RideStatus, error)
	InProgress(prevStatus RideStatus) (RideStatus, error)
	Completed(prevStatus RideStatus) (RideStatus, error)
}

func BuildRideStatus(status string) (RideStatus, error) {
	r, ok := rideStatusStates[status]
	if !ok {
		return nil, errors.New("invalid ride status")
	}
	return r, nil
}

var rideStatusStates = map[string]RideStatus{
	"requested":   RideStatusRequested{},
	"accepted":    RideStatusAccepted{},
	"in_progress": RideStatusInProgress{},
	"completed":   RideStatusCompleted{},
}

type RideStatusRequested struct{}

func (r RideStatusRequested) String() string {
	return "requested"
}

func (r RideStatusRequested) Requested(prevStatus RideStatus) (RideStatus, error) {
	return RideStatusAccepted{}, nil
}

func (r RideStatusRequested) Accepted(prevStatus RideStatus) (RideStatus, error) {
	return RideStatusAccepted{}, nil
}

func (r RideStatusRequested) InProgress(prevStatus RideStatus) (RideStatus, error) {
	return nil, errors.New("ride not accepted yet")
}

func (r RideStatusRequested) Completed(prevStatus RideStatus) (RideStatus, error) {
	return nil, errors.New("ride not accepted yet")
}

type RideStatusAccepted struct{}

func (r RideStatusAccepted) String() string {
	return "accepted"
}

func (r RideStatusAccepted) Requested(prevStatus RideStatus) (RideStatus, error) {
	return nil, errors.New("ride already accepted")
}

func (r RideStatusAccepted) Accepted(prevStatus RideStatus) (RideStatus, error) {
	return r, errors.New("ride already accepted")
}

func (r RideStatusAccepted) InProgress(prevStatus RideStatus) (RideStatus, error) {
	return RideStatusInProgress{}, nil
}

func (r RideStatusAccepted) Completed(prevStatus RideStatus) (RideStatus, error) {
	return nil, nil
}

type RideStatusInProgress struct{}

func (r RideStatusInProgress) String() string {
	return "in_progress"
}

func (r RideStatusInProgress) Requested(prevStatus RideStatus) (RideStatus, error) {
	return nil, errors.New("ride already in progress")
}

func (r RideStatusInProgress) Accepted(prevStatus RideStatus) (RideStatus, error) {
	return nil, errors.New("ride already in progress")
}

func (r RideStatusInProgress) InProgress(prevStatus RideStatus) (RideStatus, error) {
	return r, errors.New("ride already in progress")
}

func (r RideStatusInProgress) Completed(prevStatus RideStatus) (RideStatus, error) {
	return RideStatusCompleted{}, nil
}

type RideStatusCompleted struct{}

func (r RideStatusCompleted) String() string {
	return "completed"
}

func (r RideStatusCompleted) Requested(prevStatus RideStatus) (RideStatus, error) {
	return nil, errors.New("ride already completed")
}

func (r RideStatusCompleted) Accepted(prevStatus RideStatus) (RideStatus, error) {
	return nil, errors.New("ride already completed")
}

func (r RideStatusCompleted) InProgress(prevStatus RideStatus) (RideStatus, error) {
	return nil, errors.New("ride already completed")
}

func (r RideStatusCompleted) Completed(prevStatus RideStatus) (RideStatus, error) {
	return nil, errors.New("ride already completed")
}
