package repository

import (
	"context"
	app_repository "helley/src/app/repository"
	domain_user "helley/src/domain"
	mongo_model "helley/src/infra/repository/mongo/model"
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
func (u *userRepository) Store(user domain_user.User) (domain_user.User, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	result, err := u.collection.InsertOne(ctx, mongo_model.User{
		Name:      user.Name,
		Email:     user.Email,
		Password:  user.Password,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdateAt,
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
func (u *userRepository) FindById(id string) (domain_user.User, error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return domain_user.User{}, WrapError(err)
	}

	result := u.collection.FindOne(context.Background(), bson.M{"_id": objectId})

	var model = mongo_model.User{}
	err = result.Decode(&model)

	if err != nil {
		return domain_user.User{}, WrapError(err)
	}

	return domain_user.User{
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
func (u *userRepository) Update(user domain_user.User) error {
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
func (u *userRepository) Find() ([]domain_user.User, error) {
	cur, err := u.collection.Find(context.Background(), bson.M{})

	if err != nil {
		return []domain_user.User{}, WrapError(err)
	}

	defer cur.Close(context.TODO())

	var users []domain_user.User
	for cur.Next(context.TODO()) {
		var model mongo_model.User

		if err := cur.Decode(&model); err != nil {
			log.Println(err.Error())
			continue
		}

		users = append(users, domain_user.User{
			ID:       model.ID.Hex(),
			Name:     model.Name,
			Email:    model.Email,
			Password: model.Password,
		})
	}

	err = cur.Err()

	return users, err

}

func (u *userRepository) FindByEmail(email string) (domain_user.User, error) {
	result := u.collection.FindOne(context.Background(), bson.M{"email": email})

	var model = mongo_model.User{}
	err := result.Decode(&model)

	if err != nil {
		return domain_user.User{}, WrapError(err)
	}

	return domain_user.User{
		ID:       model.ID.Hex(),
		Name:     model.Name,
		Email:    model.Email,
		Password: model.Password,
	}, nil
}
