package gapi

import (
	"github.com/alonelegion/api_golang_mongodb/config"
	"github.com/alonelegion/api_golang_mongodb/pb"
	"github.com/alonelegion/api_golang_mongodb/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type Server struct {
	pb.UnimplementedAuthServiceServer
	config         config.Config
	authService    services.AuthService
	userService    services.UserService
	userCollection *mongo.Collection
}

func NewGrpcServer(config config.Config,
	authService services.AuthService,
	userService services.UserService,
	userCollection *mongo.Collection,
) *Server {
	server := &Server{
		config:         config,
		authService:    authService,
		userService:    userService,
		userCollection: userCollection,
	}

	return server
}
