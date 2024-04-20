package domain

import "fmt"

type RideStatus string

const (
	RideStatusRequested RideStatus = "requested"
	RideStatusCompleted RideStatus = "completed"
)

func BuildRideStatusFromString(status string) (RideStatus, error) {
	switch status {
	case "requested":
		return RideStatusRequested, nil
	case "completed":
		return RideStatusCompleted, nil
	default:
		return "", RaiseDomainError(fmt.Errorf("ride status %s, isn't valid", status))
	}
}
