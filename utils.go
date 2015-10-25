package wredis

import "errors"

//
// error helper functions
//

func boolError(msg string) (bool, error) {
	return false, errors.New(msg)
}

func int64Error(msg string) (int64, error) {
	return int64(0), errors.New(msg)
}

func stringError(msg string) (string, error) {
	return "", errors.New(msg)
}

func stringsError(msg string) ([]string, error) {
	return nil, errors.New(msg)
}
