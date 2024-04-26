package domain

import "github.com/eduardonunesp/cleanarchgo/pkg/domain/valueobject"

func BuildRide(opts ...rideOption) (*Ride, error) {
	r := &Ride{}
	for _, opt := range opts {
		err := opt(r)
		if err != nil {
			return nil, err
		}
	}
	return r, nil
}

func RideWithNewID() rideOption {
	return func(r *Ride) error {
		r.id = valueobject.MustUUID()
		return nil
	}
}

func RideWithID(id string) rideOption {
	return func(r *Ride) error {
		var err error
		r.id, err = valueobject.UUIDFromString(id)
		return err
	}
}

func RideWithPassengerID(passengerID string) rideOption {
	return func(r *Ride) error {
		var err error
		r.passengerID, err = valueobject.UUIDFromString(passengerID)
		return err
	}
}

func RideWithDriverID(driverID string) rideOption {
	return func(r *Ride) error {
		var err error
		r.driverID, err = valueobject.UUIDFromString(driverID)
		return err
	}
}

func RideWithStatus(status string) rideOption {
	return func(r *Ride) error {
		var err error
		r.status, err = valueobject.BuildRideStatus(status)
		return err
	}
}

func RideWithFare(fare string) rideOption {
	return func(r *Ride) error {
		r.fare = fare
		return nil
	}
}

func RideWithSegment(fromLat, fromLong, toLat, ToLong string) rideOption {
	return func(r *Ride) error {
		var err error
		r.segment, err = valueobject.BuildSegmentFromCoords(fromLat, fromLong, toLat, ToLong)
		return err
	}
}

func RideWithDate(date int64) rideOption {
	return func(r *Ride) error {
		var err error
		r.date, err = valueobject.DateFromUnix(date)
		return err
	}
}

func RideWithDateNow() rideOption {
	return func(r *Ride) error {
		r.date = valueobject.DateFromNow()
		return nil
	}
}
