package wredis

import (
	"errors"
	"fmt"
)

var (
	EmptyKeyErr     = errors.New("key cannot be an empty string")
	EmptyPatternErr = errors.New("pattern cannot be an empty string")
)

//
// error helper functions
//

func boolErr(err error) (bool, error) {
	return false, err
}

func boolError(msg string) (bool, error) {
	return false, errors.New(msg)
}

func int64Err(err error) (int64, error) {
	return -1, err
}

func int64Error(msg string) (int64, error) {
	return -1, errors.New(msg)
}

func stringErr(err error) (string, error) {
	return "", err
}

func stringError(msg string) (string, error) {
	return "", errors.New(msg)
}

func stringsErr(err error) ([]string, error) {
	return nil, err
}

func stringsError(msg string) ([]string, error) {
	return nil, errors.New(msg)
}

func unsafeError(method string) error {
	return errors.New(unsafeMessage(method))
}

func unsafeMessage(method string) string {
	return fmt.Sprintf("%s requires an Unsafe client. See wredis.NewUnsafe",
		method)
}

func checkSimpleStringResponse(cmd, res string, err error) error {
	if err != nil {
		return err
	} else if res != "OK" {
		return fmt.Errorf("%s expected OK response, got: %s", cmd, res)
	}
	return nil
}
