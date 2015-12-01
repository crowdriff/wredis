package wredis

import (
	"errors"
	"time"
)

// SetExDuration is a convenience method to set a key's value with and expiry time.
func (w *Wredis) SetExDuration(key, value string, dur time.Duration) error {
	seconds := int(dur.Seconds())
	if seconds <= 0 {
		return errors.New("duration must be at least 1 second")
	}
	return w.SetEx(key, value, seconds)
}
