package store

// Store interface
type Store interface {
	User() UserRepository
	Organization() OrganizationRepository
	EthereumAccount() EthereumAccountRepository
	NucypherAccount() NucypherAccountRepository
	Transaction() TransactionRepository
	Genome() GenomeRepository
	Virus() VirusRepository
	Vaccine() VaccineRepository
	NucypherPolicy() NucypherPolicyRepository
	NucypherReceipt() NucypherReceiptRepository
	Grants() RequestedGrantsRepository
}
