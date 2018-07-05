package main

import (
	"crypto/ecdsa"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/ethereum/go-ethereum/common"
)

func main() {
	SMAddress := "0x76C841534b303e2ed91B3Cd60baD09B8270DB755"

	// Load private key from file
	privateKey, err := crypto.LoadECDSA("privateKey.txt")
	if err != nil {
		log.Panic("Loading Private Key Error: ", err)
	}

	pubKey := privateKey.PublicKey
	addr := crypto.PubkeyToAddress(ecdsa.PublicKey(pubKey))

	fmt.Print("Using Account: ", addr.String(), "\n")

	// Create transactor
	transactOpts := bind.NewKeyedTransactor(privateKey)

	// Setup client
	cl, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Panic("Client Dial Error: ", err)
	}

	// Create SM Instance to transact
	SM, err := NewSubscriptionManager(common.HexToAddress(SMAddress), cl)
	if err != nil {
		log.Panic("Error instantiating SM: ", err)
	}
	SMSession := new(SubscriptionManagerSession)
	SMSession.Contract = SM
	SMSession.TransactOpts = *transactOpts

	// Call Bill
	transaction, err := SMSession.CallBill()
	if err != nil {
		log.Panic("Call Bill Error: ", err, "\nTransaction: ", transaction,"\n")
	}
	fmt.Print("Transaction: ", transaction, "\n")
}
