package store

import (
	"time"

	"github.com/bahadylbekov/vacinex_api/internal/app/model"
)

// UserRepository interface
type UserRepository interface {
	Create(*model.User) error
	Find(int) (*model.User, error)
	FindByEmail(string) (*model.User, error)
}

// OrganizationRepository interface
type OrganizationRepository interface {
	Create(*model.Organization, time.Time) error
	GetMyOrganization(string) (*model.Organization, error)
	Update(*model.Organization, time.Time) error
}
