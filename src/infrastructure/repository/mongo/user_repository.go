package repository_mongo

import (
	app_repository "app-helley/src/application/repository"
	"app-helley/src/domain"
	mongo_model "app-helley/src/infrastructure/repository/mongo/model"
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

// Create a repository using mongodb as a data storage
func NewUserRepository(db *mongo.Database) app_repository.UserRepository {
	return &userRepository{
		collection: db.Collection("user"),
	}
}

/**
Store user on database mongodb
*/
func (u *userRepository) Store(user domain.User) (domain.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := u.collection.InsertOne(ctx, mongo_model.User{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	})

	if mongo.IsDuplicateKeyError(err) {
		return user, WrapError(err)
	}

	if err != nil {
		return user, WrapError(err)
	}

	user.ID = result.InsertedID.(primitive.ObjectID).Hex()

	return user, err
}

// Find user by ID on database mongodb
func (u *userRepository) FindById(id string) (domain.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return domain.User{}, WrapError(err)
	}

	result := u.collection.FindOne(context.Background(), bson.M{"_id": objectId})

	var model = mongo_model.User{}
	err = result.Decode(&model)

	if err != nil {
		return domain.User{}, WrapError(err)
	}

	return domain.User{
		ID:       model.ID.Hex(),
		Name:     model.Name,
		Email:    model.Email,
		Password: model.Password,
	}, nil

}

// Delete user by ID on database mongodb
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

// Update user on database mongodb
func (u *userRepository) Update(user domain.User) error {
	ID, err := primitive.ObjectIDFromHex(user.ID)
	if err != nil {
		return WrapError(err)
	}

	model := mongo_model.User{
		ID:       ID,
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
	}

	_, err = u.collection.UpdateByID(context.Background(), model.ID, model)
	if err != nil {
		return WrapError(err)
	}
	return nil
}

// Find all users on database mongodb
func (u *userRepository) Find() ([]domain.User, error) {
	cur, err := u.collection.Find(context.Background(), bson.M{})

	if err != nil {
		return []domain.User{}, WrapError(err)
	}

	defer cur.Close(context.TODO())

	var users []domain.User
	for cur.Next(context.TODO()) {
		var model mongo_model.User

		if err := cur.Decode(&model); err != nil {
			log.Println(err.Error())
			continue
		}

		users = append(users, domain.User{
			ID:       model.ID.Hex(),
			Name:     model.Name,
			Email:    model.Email,
			Password: model.Password,
		})
	}

	err = cur.Err()

	return users, err

}
