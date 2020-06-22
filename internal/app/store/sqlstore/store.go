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
	tezosAccountRepository    *TezosAccountRepository
	ethereumAccountRepository *EthereumAccountRepository
	nucypherAccountRepository *NucypherAccountRepository
	genomeRepository          *GenomeRepository
	virusRepository           *VirusRepository
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

// TezosAccount ...
func (s *Store) TezosAccount() store.TezosAccountRepository {
	if s.tezosAccountRepository != nil {
		return s.tezosAccountRepository
	}

	s.tezosAccountRepository = &TezosAccountRepository{
		store: s,
	}

	return s.tezosAccountRepository
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
