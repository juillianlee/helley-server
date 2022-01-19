package repository

import (
	"errors"
	app_repository "helley/src/app/repository"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func WrapError(err error) error {
	switch {
	case errors.Is(err, mongo.ErrNoDocuments):
		return app_repository.ErrNotFoundRegister
	case mongo.IsDuplicateKeyError(err):
		return app_repository.ErrDuplicateKey
	case errors.Is(err, mongo.ErrInvalidIndexValue):
		return app_repository.ErrGenericRepository(err.Error())
	case errors.Is(err, primitive.ErrInvalidHex):
		return app_repository.ErrGenericRepository(err.Error())
	default:
		return app_repository.ErrGenericRepository(err.Error())
	}
}
