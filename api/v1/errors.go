package v1

import "errors"

func NotInitializedError() (err error) {
	return errors.New("method not initialized")
}

func NotImplementedError() (err error) {
	return errors.New("method not implemented")
}
