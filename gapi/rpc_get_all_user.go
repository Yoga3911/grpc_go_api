package gapi

import (
	"context"
	"simple_bank/pb"

	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *Server) GetAllUser(ctx context.Context, _ *emptypb.Empty) (*pb.GetAllUserResponse, error) {	
	users, _ := userR.GetAllUser()
	resp := pb.GetAllUserResponse{
		Users: convertUsers(users),
	}
	return &resp, nil
}