package valueobject

import (
	"fmt"
	"time"
)

type Date int64

func DateFromInt64(date int64) Date {
	return Date(date)
}

func DateFromNow() Date {
	return Date(time.Now().Unix())
}

func (d Date) Int64() int64 {
	return int64(d)
}

func (d Date) IsZero() bool {
	return d == 0
}

func (d Date) String() string {
	return fmt.Sprintf("%d", d)
}
