package model

import (
	pb "github.com/Zzocker/bl-proto-go/userprofile"
)

// UserTogRPC :
func UserTogRPC(in *User) *pb.User {
	return &pb.User{}
}

// GRPCToUser :
func GRPCToUser(in *pb.User) *User {
	return &User{}
}
