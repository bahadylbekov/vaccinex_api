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
	GetOrganizations() ([]*model.Organization, error)
	GetMyOrganization(string) ([]*model.Organization, error)
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

// TezosAccountRepository ...
type TezosAccountRepository interface {
	Create(*model.TezosAccount, time.Time) error
	GetAccounts(string) ([]*model.TezosAccount, error)
	UpdateName(string, string, int, time.Time) error
	Deactivate(string, int, time.Time) error
	Reactivate(string, int, time.Time) error
	Private(string, int, time.Time) error
	Unprivate(string, int, time.Time) error
}

// EthereumAccountRepository ...
type EthereumAccountRepository interface {
	Create(*model.EthereumAccount, time.Time) error
	GetAccounts(string) ([]*model.EthereumAccount, error)
	UpdateName(string, string, int, time.Time) error
	Deactivate(string, int, time.Time) error
	Reactivate(string, int, time.Time) error
	Private(string, int, time.Time) error
	Unprivate(string, int, time.Time) error
}

// NucypherAccountRepository ...
type NucypherAccountRepository interface {
	Create(*model.NucypherAccount, time.Time) error
	GetAccounts(string) ([]*model.NucypherAccount, error)
	UpdateName(string, string, int, time.Time) error
	Deactivate(string, int, time.Time) error
	Reactivate(string, int, time.Time) error
	Private(string, int, time.Time) error
	Unprivate(string, int, time.Time) error
}

// GenomeRepository ...
type GenomeRepository interface {
	Create(*model.Genome, time.Time) error
	GetMyGenomes(string) ([]*model.Genome, error)
	GetGenomes() ([]*model.Genome, error)
	GetGenomesByVirus(string) ([]*model.Genome, error)
	GetGenomesByOrganization(string) ([]*model.Genome, error)
}

// VirusRepository ...
type VirusRepository interface {
	Create(*model.Virus, time.Time) error
	GetViruses() ([]*model.Virus, error)
	GetVirusByID(string) (*model.Virus, error)
	Update(*model.Virus, time.Time) error
}
