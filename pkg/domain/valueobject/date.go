package valueobject

import (
	"fmt"
	"time"
)

var (
	ErrInvalidDate = fmt.Errorf("invalid date")
)

type Date struct {
	value int64
}

func DateFromUnix(date int64) (Date, error) {
	var newDate Date
	if date <= 0 {
		return newDate, ErrInvalidDate
	}
	newDate.value = date
	return newDate, nil
}

func DateFromNow() Date {
	return Date{time.Now().Unix()}
}

func (d Date) Int64() int64 {
	return int64(d.value)
}

func (d Date) IsZero() bool {
	return d.value == 0
}

func (d Date) String() string {
	return fmt.Sprintf("%d", d.value)
}
