package errors

import (
	"errors"
	"net/http"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	ErrAuth       = &sentinelAPIError{status: http.StatusUnauthorized, msg: "Username or Password invalid"}
	ErrNotFound   = &sentinelAPIError{status: http.StatusNotFound, msg: "Not found result"}
	ErrDuplicate  = &sentinelAPIError{status: http.StatusBadRequest, msg: "already registered"}
	ErrBadRequest = &sentinelAPIError{status: http.StatusBadRequest, msg: "Bad request"}
)

type APIError interface {
	APIError() (int, string)
}

type sentinelAPIError struct {
	status int
	msg    string
}

func (e sentinelAPIError) Error() string {
	return e.msg
}

func (e sentinelAPIError) APIError() (int, string) {
	return e.status, e.msg
}

type sentinelWrappedError struct {
	error
	sentinel *sentinelAPIError
}

func (e sentinelWrappedError) Is(err error) bool {
	return e.sentinel == err
}

func (e sentinelWrappedError) APIError() (int, string) {
	return e.sentinel.APIError()
}

func WrapError(err error) error {

	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return sentinelWrappedError{err, ErrNotFound}
	case mongo.IsDuplicateKeyError(err):
		return sentinelWrappedError{err, ErrDuplicate}
	case errors.Is(err, mongo.ErrInvalidIndexValue):
		return sentinelWrappedError{err, ErrBadRequest}
	case errors.Is(err, primitive.ErrInvalidHex):
		return sentinelWrappedError{err, ErrNotFound}
	default:
		return err
	}
}
