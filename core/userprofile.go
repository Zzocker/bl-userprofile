package core

import (
	"github.com/Zzocker/bl-userprofile/core/ports"
	"github.com/Zzocker/bl-userprofile/model"
	"github.com/Zzocker/bl-utils/pkg/errors"
)

// UserProfileBusiness :
type UserProfileBusiness struct {
	DS ports.UserProfileDatastoreInterface
}

// Register :
func (u *UserProfileBusiness) Register(in *model.RegisterUser) *errors.Er {
	return errors.NewMsgln(errors.UNIMPLEMENTED, "core business")
}

// GetUser :
func (u *UserProfileBusiness) GetUser(username string) (*model.User, *errors.Er) {
	return nil, errors.NewMsgln(errors.UNIMPLEMENTED, "core business")

}

// CheckCred :
func (u *UserProfileBusiness) CheckCred(username, password string) *errors.Er {
	return errors.NewMsgln(errors.UNIMPLEMENTED, "core business")
}
