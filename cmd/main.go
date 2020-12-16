package main

import (
	"log"
	"net"

	pb "github.com/Zzocker/bl-proto-go/userprofile"
	"github.com/Zzocker/bl-userprofile/core"
	"github.com/Zzocker/bl-userprofile/userside/grpcside"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", ":8081")
	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()
	side, err := creategRPCSide()
	if err != nil {
		panic(err)
	}
	reflection.Register(srv)
	pb.RegisterUserProfileServer(srv, side)
	log.Println("server started at 8081")
	if err := srv.Serve(lis); err != nil {
		panic(err)
	}
}

func creategRPCSide() (pb.UserProfileServer, error) {
	cre, err := createCore()
	if err != nil {
		return nil, err
	}
	return &grpcside.UserProfileSide{
		Core: cre,
	}, nil
}

func createCore() (core.UserProfile, error) {
	return &core.UserProfileBusiness{}, nil
}

// func createDS()(ports.UserProfileDatastoreInterface,error){
// 	return
// }
