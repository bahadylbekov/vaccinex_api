package store

// Store interface
type Store interface {
	User() UserRepository
	Organization() OrganizationRepository
	Account() AccountRepository
	Transaction() TransactionRepository
}
