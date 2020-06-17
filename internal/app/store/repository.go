package store

import (
	"time"

	"github.com/bahadylbekov/vaccinex_api/internal/app/model"
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
	FindOrganizationsByEmail(string) ([]*model.Organization, error)
	AddOrganizationToMyList(string, string, string, time.Time) error
	GetConnectedOrganizations(string) ([]*model.Organization, error)
	GetOrganization(string) (*model.Organization, error)
	Delete(string, string) error
	ID(string) (*string, error)
}

// TransactionRepository interface
type TransactionRepository interface {
	Create(*model.Transaction, time.Time) error
	GetTransactions(string) ([]*model.Transaction, error)
	GetSendTransactions(string) ([]*model.Transaction, error)
	GetRecievedTransactions(string) ([]*model.Transaction, error)
}

// AccountRepository ...
type AccountRepository interface {
	Create(*model.Account, time.Time) error
	GetAccounts(string) ([]*model.Account, error)
	UpdateName(string, string, int, time.Time) error
	Deactivate(string, int, time.Time) error
	Reactivate(string, int, time.Time) error
	Private(string, int, time.Time) error
	Unprivate(string, int, time.Time) error
}
