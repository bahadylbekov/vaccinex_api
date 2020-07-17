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

// EthereumAccountRepository ...
type EthereumAccountRepository interface {
	Create(*model.EthereumAccount, time.Time) error
	GetAccounts(string) ([]*model.EthereumAccount, error)
	GetAccountByOrganization(string) ([]*model.EthereumAccount, error)
	UpdateName(string, string, int, time.Time) error
	UpdateAddress(string, string, int, time.Time) error
	Deactivate(string, int, time.Time) error
	Reactivate(string, int, time.Time) error
	Private(string, int, time.Time) error
	Unprivate(string, int, time.Time) error
}

// NucypherAccountRepository ...
type NucypherAccountRepository interface {
	Create(*model.NucypherAccount, time.Time) error
	GetAccounts(string) ([]*model.NucypherAccount, error)
	GetAccountByOrganization(string) ([]*model.NucypherAccount, error)
	UpdateName(string, string, int, time.Time) error
	UpdateAddress(string, string, int, time.Time) error
	UpdateVerifyingKey(string, string, int, time.Time) error
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
	GetGenomesByVaccine(string) ([]*model.Genome, error)
}

// VirusRepository ...
type VirusRepository interface {
	Create(*model.Virus, time.Time) error
	GetViruses() ([]*model.Virus, error)
	GetVirusByID(string) (*model.Virus, error)
	Update(*model.Virus, time.Time) error
}

// VaccineRepository ...
type VaccineRepository interface {
	Create(*model.Vaccine, time.Time) error
	GetVaccines() ([]*model.Vaccine, error)
	GetVaccineByID(string) ([]*model.Vaccine, error)
	UpdateAmount(string, string, int, time.Time) error
	UpdateName(string, string, int, time.Time) error
	UpdateDescription(string, string, int, time.Time) error
}

type NucypherPolicyRepository interface {
	Create(*model.NucypherPolicy, time.Time) error
	GetByID(string) (*model.NucypherPolicy, error)
	GetByLabel(string) (*model.NucypherPolicy, error)
}

type NucypherReceiptRepository interface {
	Create(*model.NucypherReceipt, time.Time) error
	GetByID(string) (*model.NucypherReceipt, error)
	GetReceiptByHash(string) (*model.NucypherReceipt, error)
}

type RequestedGrantsRepository interface {
	Create(*model.RequestedGrant, time.Time) error
	Submit(bool, string, string, string, time.Time) error
	GetGrantsForMe(string) ([]*model.RequestedGrant, error)
	GetCompletedGrantsForMe(string) ([]*model.RequestedGrant, error)
}
