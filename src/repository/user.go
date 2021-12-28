package repository

import (
	"app-helley/src/domain"
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type UserRepository interface {
	Store(user domain.User) (domain.User, error)
	DeleteById(id string) error
	FindById(id string) (domain.User, error)
	Update(user domain.User) error
	Find() ([]domain.User, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(db *mongo.Database) UserRepository {
	return &userRepository{
		collection: db.Collection("user"),
	}
}

func (u *userRepository) Store(user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := u.collection.InsertOne(ctx, user)
	if err != nil {
		return domain.User{}, err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	return user, err
}

func (u *userRepository) FindById(id string) (domain.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return domain.User{}, err
	}

	result := u.collection.FindOne(context.Background(), bson.M{"_id": objectId})
	var user = domain.User{}
	err = result.Decode(&user)
	return user, err

}

func (u *userRepository) DeleteById(id string) error {

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = u.collection.DeleteOne(context.Background(), bson.M{"_id": objectId})

	return err
}

func (u *userRepository) Update(user domain.User) error {
	_, err := u.collection.UpdateByID(context.Background(), user.ID, user)
	return err
}

func (u *userRepository) Find() ([]domain.User, error) {
	cur, err := u.collection.Find(context.Background(), bson.M{})

	if err != nil {
		return []domain.User{}, nil
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
