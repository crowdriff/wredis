package wredis

import (
	"errors"
	"fmt"
)

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

func checkSimpleStringResponse(cmd, res string, err error) error {
	if err != nil {
		return err
	} else if res != "OK" {
		return fmt.Errorf("%s did not get OK response: %s", cmd, res)
	}
	return nil
}
