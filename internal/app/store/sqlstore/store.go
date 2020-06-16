package sqlstore

import (
	"github.com/bahadylbekov/vacinex_api/internal/app/store"
	"github.com/jmoiron/sqlx"
)

// Store ..
type Store struct {
	db                     *sqlx.DB
	userRepository         *UserRepository
	organizationRepository *OrganizationRepository
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
