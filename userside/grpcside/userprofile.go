package grpcside

import (
	"context"
	"log"

	"github.com/Zzocker/bl-proto-go/common"
	pb "github.com/Zzocker/bl-proto-go/userprofile"
	"github.com/Zzocker/bl-userprofile/config"
	"github.com/Zzocker/bl-userprofile/core"
	"github.com/Zzocker/bl-userprofile/model"
	"google.golang.org/grpc/metadata"
)

// UserProfileSide :
type UserProfileSide struct {
	Core core.UserProfile
	pb.UnimplementedUserProfileServer
}

// Register :
func (u *UserProfileSide) Register(ctx context.Context, in *pb.RegisterRequest) (*pb.RegisterResponse, error) {
	log.Println("inside Register")
	header := common.Header{
		Status: common.StatusCode_OK,
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		header.Status = common.StatusCode_INVALID_ARGUMENT
		header.Description = "no metadata found"
		return &pb.RegisterResponse{Header: &header}, nil
	}
	raw := md.Get(config.METADATA_SECRET)
	if len(raw) != 1 {
		header.Status = common.StatusCode_INVALID_ARGUMENT
		header.Description = "no SECRET key found in metadata"
		return &pb.RegisterResponse{Header: &header}, nil
	}
	input := model.RegisterUser{
		Username: in.Username,
		Email:    in.Email,
		Phone:    in.Phone,
		Gender:   in.Gender.String(),
		Password: raw[0],
		Name:     in.Name,
	}
	log.Println(input)
	err := u.Core.Register(&input)
	if err != nil {
		header.Status = common.StatusCode(err.Status)
		header.Description = err.Error()
	}
	return &pb.RegisterResponse{
		Header: &header,
	}, nil
}

// CheckCred :
func (u *UserProfileSide) CheckCred(ctx context.Context, in *pb.CheckCredRequest) (*pb.CheckCredResponse, error) {
	log.Println("inside CheckCred")
	header := common.Header{
		Status: common.StatusCode_OK,
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		header.Status = common.StatusCode_INVALID_ARGUMENT
		header.Description = "no metadata found"
		return &pb.CheckCredResponse{Header: &header}, nil
	}
	raw := md.Get(config.METADATA_SECRET)
	if len(raw) != 1 {
		header.Status = common.StatusCode_INVALID_ARGUMENT
		header.Description = "no SECRET key found in metadata"
		return &pb.CheckCredResponse{Header: &header}, nil
	}
	err := u.Core.CheckCred(in.Username, raw[0])
	ok = true
	if err != nil {
		header.Status = common.StatusCode(err.Status)
		header.Description = err.Error()
		ok = false
	}
	return &pb.CheckCredResponse{
		Header: &header,
		Ok:     ok,
	}, nil
}

// GetUser :
func (u *UserProfileSide) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.GetUserResponse, error) {
	log.Println("inside GetUser")
	header := common.Header{
		Status: common.StatusCode_OK,
	}
	usr, err := u.Core.GetUser(in.Username)
	if err != nil {
		header.Status = common.StatusCode(err.Status)
		header.Description = err.Error()
		return &pb.GetUserResponse{Header: &header}, nil
	}
	return &pb.GetUserResponse{
		Header: &header,
		User:   model.UserTogRPC(usr),
	}, nil
}
