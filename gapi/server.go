package gapi

import (
	"simple_bank/pb"
)

type Server struct {
	pb.UnimplementedSimpleBankServer
}

func NewServer() (*Server, error) {
	return &Server{}, nil
}
