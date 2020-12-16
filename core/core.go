package core

import (
	"github.com/Zzocker/bl-userprofile/model"
	"github.com/Zzocker/bl-utils/pkg/errors"
)

// UserProfile :
type UserProfile interface {
	Register(in *model.RegisterUser) *errors.Er
	GetUser(username string) (*model.User, *errors.Er)
	CheckCred(username, password string) *errors.Er
}
