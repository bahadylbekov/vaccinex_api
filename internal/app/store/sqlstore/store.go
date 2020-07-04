package sqlstore

import (
	"github.com/bahadylbekov/vaccinex_api/internal/app/store"
	"github.com/jmoiron/sqlx"
)

// Store ..
type Store struct {
	db                        *sqlx.DB
	userRepository            *UserRepository
	organizationRepository    *OrganizationRepository
	transactionRepository     *TransactionRepository
	ethereumAccountRepository *EthereumAccountRepository
	nucypherAccountRepository *NucypherAccountRepository
	genomeRepository          *GenomeRepository
	virusRepository           *VirusRepository
	vaccineRepository         *VaccineRepository
	policyRepository          *NucypherPolicyRepository
	receiptRepository         *NucypherReceiptRepository
}

// New ...
func New(db *sqlx.DB) *Store {
	return &Store{
		db: db,
	}
}

// User ...
func (s *Store) User() store.UserRepository {
	if s.userRepository != nil {
		return s.userRepository
	}

	s.userRepository = &UserRepository{
		store: s,
	}

	return s.userRepository
}

// Organization ...
func (s *Store) Organization() store.OrganizationRepository {
	if s.organizationRepository != nil {
		return s.organizationRepository
	}

	s.organizationRepository = &OrganizationRepository{
		store: s,
	}

	return s.organizationRepository
}

// Transaction ...
func (s *Store) Transaction() store.TransactionRepository {
	if s.transactionRepository != nil {
		return s.transactionRepository
	}

	s.transactionRepository = &TransactionRepository{
		store: s,
	}

	return s.transactionRepository
}

// EthereumAccount ...
func (s *Store) EthereumAccount() store.EthereumAccountRepository {
	if s.ethereumAccountRepository != nil {
		return s.ethereumAccountRepository
	}

	s.ethereumAccountRepository = &EthereumAccountRepository{
		store: s,
	}

	return s.ethereumAccountRepository
}

// NucypherAccount ...
func (s *Store) NucypherAccount() store.NucypherAccountRepository {
	if s.nucypherAccountRepository != nil {
		return s.nucypherAccountRepository
	}

	s.nucypherAccountRepository = &NucypherAccountRepository{
		store: s,
	}

	return s.nucypherAccountRepository
}

// Genome ...
func (s *Store) Genome() store.GenomeRepository {
	if s.genomeRepository != nil {
		return s.genomeRepository
	}

	s.genomeRepository = &GenomeRepository{
		store: s,
	}

	return s.genomeRepository
}

// Virus ...
func (s *Store) Virus() store.VirusRepository {
	if s.virusRepository != nil {
		return s.virusRepository
	}

	s.virusRepository = &VirusRepository{
		store: s,
	}

	return s.virusRepository
}

// Vaccine ...
func (s *Store) Vaccine() store.VaccineRepository {
	if s.vaccineRepository != nil {
		return s.vaccineRepository
	}

	s.vaccineRepository = &VaccineRepository{
		store: s,
	}

	return s.vaccineRepository
}

func (s *Store) NucypherPolicy() store.NucypherPolicyRepository {
	if s.policyRepository != nil {
		return s.policyRepository
	}

	s.policyRepository = &NucypherPolicyRepository{
		store: s,
	}

	return s.policyRepository

}

func (s *Store) NucypherReceipt() store.NucypherReceiptRepository {
	if s.receiptRepository != nil {
		return s.receiptRepository
	}

	s.receiptRepository = &NucypherReceiptRepository{
		store: s,
	}

	return s.receiptRepository

}
