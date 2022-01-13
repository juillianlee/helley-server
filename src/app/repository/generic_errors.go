package repository

import "errors"

var (
	ErrDuplicateKey     = errors.New("duplicate key violates unique constraint")
	ErrNotFoundRegister = errors.New("register not found")
	ErrDeleteEntity     = errors.New("fail delete register")
	ErrStoreEntity      = errors.New("fail store entity")
	ErrUpdateEntity     = errors.New("fail update entity")
)

func ErrGenericRepository(msg string) error {
	return errors.New(msg)
}
