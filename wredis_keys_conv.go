package wredis

import "fmt"

// DelWithPattern is a convenience method that Deletes all keys matching the pattern
func (w *Wredis) DelWithPattern(pattern string) (int64, error) {
	if w.safe {
		return int64Error(unsafeMessage("DelWithPattern"))
	}
	keys, err := w.Keys(pattern)
	if err != nil {
		return int64Error(err.Error())
	}
	if len(keys) == 0 {
		return int64Error(fmt.Sprintf("no keys found with pattern: %s",
			pattern))
	}
	return w.Del(keys...)
}
