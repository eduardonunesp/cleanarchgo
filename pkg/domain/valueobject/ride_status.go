package valueobject

import "fmt"

type RideStatus string

const (
	RideStatusRequested RideStatus = "requested"
	RideStatusAccepted  RideStatus = "accepted"
	RideStatusInProgres RideStatus = "in_progress"
	RideStatusCompleted RideStatus = "completed"
)

func RideStatusFromString(status string) (RideStatus, error) {
	switch status {
	case "requested":
		return RideStatusRequested, nil
	case "accepted":
		return RideStatusAccepted, nil
	case "in_progress":
		return RideStatusInProgres, nil
	case "completed":
		return RideStatusCompleted, nil
	default:
		return "", fmt.Errorf("ride status %s, isn't valid", status)
	}
}
