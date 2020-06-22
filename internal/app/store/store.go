package store

// Store interface
type Store interface {
	User() UserRepository
	Organization() OrganizationRepository
	TezosAccount() TezosAccountRepository
	EthereumAccount() EthereumAccountRepository
	Transaction() TransactionRepository
	Genome() GenomeRepository
	Virus() VirusRepository
}
