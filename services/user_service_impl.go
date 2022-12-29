package services

import (
	"fmt"
	"strings"

	"github.com/alonelegion/api_golang_mongodb/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/net/context"
)

type UserServiceImpl struct {
	collection *mongo.Collection
	ctx        context.Context
}

func NewUserServiceImpl(collection *mongo.Collection, ctx context.Context) UserService {
	return &UserServiceImpl{collection, ctx}
}

func (us UserServiceImpl) FindUserById(id string) (*models.DBResponse, error) {
	oid, _ := primitive.ObjectIDFromHex(id)

	var user *models.DBResponse

	query := bson.M{"_id": oid}
	err := us.collection.FindOne(us.ctx, query).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.DBResponse{}, err
		}
		return nil, err
	}

	return user, nil
}

func (us UserServiceImpl) FindUserByEmail(email string) (*models.DBResponse, error) {
	var user *models.DBResponse

	query := bson.M{"email": strings.ToLower(email)}
	err := us.collection.FindOne(us.ctx, query).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return &models.DBResponse{}, err
		}
		return nil, err
	}

	return user, nil
}

func (us UserServiceImpl) UpdateUserById(id, field, value string) (*models.DBResponse, error) {
	userId, _ := primitive.ObjectIDFromHex(id)
	query := bson.D{{Key: "_id", Value: userId}}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: field, Value: value}}}}
	result, err := us.collection.UpdateOne(us.ctx, query, update)
	fmt.Print(result.ModifiedCount)
	if err != nil {
		fmt.Println(err)
		return &models.DBResponse{}, err
	}

	return &models.DBResponse{}, nil
}

func (us UserServiceImpl) UpdateOne(field string, value interface{}) (*models.DBResponse, error) {
	query := bson.D{{Key: field, Value: value}}
	update := bson.D{{Key: "$key", Value: bson.D{{Key: field, Value: value}}}}
	result, err := us.collection.UpdateOne(us.ctx, query, update)
	fmt.Println(result.ModifiedCount)
	if err != nil {
		fmt.Println(err)
		return &models.DBResponse{}, err
	}

	return &models.DBResponse{}, nil
}
