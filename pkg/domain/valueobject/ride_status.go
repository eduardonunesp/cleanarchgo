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
		RideStatusFinished{},
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
	Finish() (RideStatus, error)
}

func BuildRideStatus(status string) (RideStatus, error) {
	r, ok := rideStatusStates[status]
	if !ok {
		return nil, errors.New("invalid ride status")
	}
	return r, nil
}

func RideStatusAs(status RideStatus, target any) bool {
	if target == nil {
		return false
	}
	_, ok := target.(RideStatus)
	if !ok {
		return false
	}
	target = status
	return true
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

func (r RideStatusRequested) Finish() (RideStatus, error) {
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

func (r RideStatusAccepted) Finish() (RideStatus, error) {
	return nil, errors.New("invalid status")
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

func (r RideStatusInProgress) Finish() (RideStatus, error) {
	return RideStatusFinished{}, nil
}

type RideStatusFinished struct{}

func (r RideStatusFinished) String() string {
	return "completed"
}

func (r RideStatusFinished) Request() (RideStatus, error) {
	return nil, errors.New("invalid status")
}

func (r RideStatusFinished) Accept() (RideStatus, error) {
	return nil, errors.New("invalid status")
}

func (r RideStatusFinished) Start() (RideStatus, error) {
	return nil, errors.New("invalid status")
}

func (r RideStatusFinished) Finish() (RideStatus, error) {
	return nil, errors.New("invalid status")
}
