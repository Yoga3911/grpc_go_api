package gapi

import (
	"simple_bank/models"
	"simple_bank/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertUser(u *models.User) *pb.User {
	return &pb.User{
		Username:          u.Username,
		FullName:          u.FullName,
		Email:             u.Email,
		PasswordChangedAt: timestamppb.Now(),
		CreateAt:          timestamppb.Now(),
	}
}

func convertUsers(u []*models.User) []*pb.User {
	var users []*pb.User
	for _, v := range u {
		userPb := pb.User{
			Username: v.Username,
			FullName: v.FullName,
			Email: v.Email,
			PasswordChangedAt: timestamppb.Now(),
			CreateAt: timestamppb.Now(),
		}
		users = append(users, &userPb)
	}

	return users
}
