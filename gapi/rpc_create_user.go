package gapi

import (
	"context"
	"simple_bank/models"
	"simple_bank/pb"
	"simple_bank/repository"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

var (
	userR = repository.NewUserRepository()
)

func (server *Server) CreateUser(ctx context.Context, req *pb.CreateUserRequest) (*pb.CreateUserResponse, error) {
	user, err := userR.InsertUser(&models.User{
		Username: req.GetUsername(),
		FullName: req.GetFullName(),
		Email: req.GetEmail(),
	})
	if err != nil {
		return nil, status.Errorf(codes.Internal, err.Error())
	}

	resp := &pb.CreateUserResponse{
		User: convertUser(user),
	}

	return resp, nil
}
