package store

// Store interface
type Store interface {
	User() UserRepository
	Organization() OrganizationRepository
	TezosAccount() TezosAccountRepository
	EthereumAccount() EthereumAccountRepository
	NucypherAccount() NucypherAccountRepository
	Transaction() TransactionRepository
	Genome() GenomeRepository
	Virus() VirusRepository
}
