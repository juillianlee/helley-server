package repository_mongo

import (
	app_repository "app-helley/src/application/repository"
	"app-helley/src/domain"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) app_repository.UserRepository {
	return &userRepository{
		collection: db.Collection("user"),
	}
}

func (u *userRepository) Store(user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := u.collection.InsertOne(ctx, user)

	if mongo.IsDuplicateKeyError(err) {
		return user, WrapError(err)
	}

	if err != nil {
		return user, WrapError(err)
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	return user, err
}

func (u *userRepository) FindById(id string) (domain.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return domain.User{}, WrapError(err)
	}

	result := u.collection.FindOne(context.Background(), bson.M{"_id": objectId})

	var user = domain.User{}
	err = result.Decode(&user)

	if err != nil {
		return user, WrapError(err)
	}

	return user, nil

}

func (u *userRepository) DeleteById(id string) error {

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return WrapError(err)
	}

	_, err = u.collection.DeleteOne(context.Background(), bson.M{"_id": objectId})

	if err != nil {
		return WrapError(err)
	}

	return nil
}

func (u *userRepository) Update(user domain.User) error {
	_, err := u.collection.UpdateByID(context.Background(), user.ID, user)
	if err != nil {
		return WrapError(err)
	}
	return nil
}

func (u *userRepository) Find() ([]domain.User, error) {
	cur, err := u.collection.Find(context.Background(), bson.M{})

	if err != nil {
		return []domain.User{}, WrapError(err)
	}

	defer cur.Close(context.TODO())

	var users []domain.User
	for cur.Next(context.TODO()) {
		var user domain.User

		if err := cur.Decode(&user); err != nil {
			log.Println(err.Error())
			continue
		}

		users = append(users, user)
	}

	err = cur.Err()

	return users, err

}
