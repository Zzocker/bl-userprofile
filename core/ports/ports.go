package ports

import (
	"github.com/Zzocker/bl-userprofile/model"
	"github.com/Zzocker/bl-utils/pkg/errors"
)

// UserProfileDatastoreInterface :
type UserProfileDatastoreInterface interface {
	Create(in *model.RegisterUser) *errors.Er
	GetByUsername(username string) (*model.User, *errors.Er)
	Delete(username string) *errors.Er
}
