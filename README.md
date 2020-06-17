# VaccineX

VaccineX is a distributed marketplace for genome data used in vaccine development. Most viruses mutates over time and in order to be up to date in a vaccine development proccess research teams must have access to recent genome mutations. 

* Collaborative - Marketplace allows research groups share data with their partners worldwide for free by controlling access to that information using proxy re-encryption scheme provided by NuCypher.
* Distributed - The network is censorship-resistant and allows research teams share pre-encrypted genome data using IPFS without centralized cloud services.
* Decentralized - VaccineX is a blockchain agnostic project which gives users an ability pay for digital assets using Ethereum and Tezos crypto currencies.

## Requirements
Require go version >=1.12 , so make sure your go version is okay.
WARNING! Building happen only when this project locates outside of GOPATH environment.

## Getting started

1. Clone this repository
2. Make sure go.mod does not contain unused dependencies by running `go mod tidy`
3. Run `make`
4. Check port :8000