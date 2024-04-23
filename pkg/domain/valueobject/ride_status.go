package valueobject

import (
	"errors"
	"fmt"
)

var rideStatusStates = map[string]RideStatus{}

func init() {
	rideStatuses := []RideStatus{
		RideStatusRequested{},
		RideStatusAccepted{},
		RideStatusInProgress{},
	}
	for _, status := range rideStatuses {
		rideStatusStates[status.String()] = status
	}
}

type RideStatus interface {
	fmt.Stringer
	Request() (RideStatus, error)
	Accept() (RideStatus, error)
	Start() (RideStatus, error)
}

func BuildRideStatus(status string) (RideStatus, error) {
	r, ok := rideStatusStates[status]
	if !ok {
		return nil, errors.New("invalid ride status")
	}
	return r, nil
}

type RideStatusRequested struct{}

func (r RideStatusRequested) String() string {
	return "requested"
}

func (r RideStatusRequested) Request() (RideStatus, error) {
	return nil, errors.New("invalid status")
}

func (r RideStatusRequested) Accept() (RideStatus, error) {
	return RideStatusAccepted{}, nil
}

func (r RideStatusRequested) Start() (RideStatus, error) {
	return nil, errors.New("invalid status")
}

type RideStatusAccepted struct{}

func (r RideStatusAccepted) String() string {
	return "accepted"
}

func (r RideStatusAccepted) Request() (RideStatus, error) {
	return nil, errors.New("invalid status")
}

func (r RideStatusAccepted) Accept() (RideStatus, error) {
	return nil, errors.New("invalid status")
}

func (r RideStatusAccepted) Start() (RideStatus, error) {
	return RideStatusInProgress{}, nil
}

type RideStatusInProgress struct{}

func (r RideStatusInProgress) String() string {
	return "in_progress"
}

func (r RideStatusInProgress) Request() (RideStatus, error) {
	return nil, errors.New("invalid status")
}

func (r RideStatusInProgress) Accept() (RideStatus, error) {
	return nil, errors.New("invalid status")
}

func (r RideStatusInProgress) Start() (RideStatus, error) {
	return nil, errors.New("invalid status")
}
