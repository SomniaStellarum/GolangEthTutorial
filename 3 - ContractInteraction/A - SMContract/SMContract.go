package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"fmt"
)

func main() {
	SMAddress := "0x6d540C9f4357FE128c6B3300a12A16B38a5bCB3b"

	// Setup client
	cl, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Panic("Client Dial Error: ", err)
	}

	// Create SM Factory Contract Instance
	SMFac, err := NewSubscriptionManagerFactory(common.HexToAddress(SMAddress), cl)
	if err != nil {
		log.Panic("SM Instantiation Error: ", err)
	}

	// Fetch SM Addresses
	i := int64(0)
	var addrs []common.Address
	for {
		addr, err := SMFac.SubscriptionManagerFactoryCaller.SubscriptionManagerAddresses(nil, big.NewInt(i))
		if err != nil {
			fmt.Print("End of Array: ", err, "\n")
			break
		}
		addrs = append(addrs, addr)
		i++
	}
	for _, addr := range addrs {
		// Create SM Contract Instance
		SM, err := NewSubscriptionManager(addr, cl)
		if err != nil {
			log.Panic("Error instantiating SM: ", err)
		}
		SMSession := new(SubscriptionManagerSession)
		SMSession.Contract = SM

		// Fetch Billers
		b, err := SMSession.Billers(common.HexToAddress("0xC928027EbeEcF95A695c183a200447649EB87a1F"))
		if err != nil {
			log.Panic("Error fetching biller: ", err)
		}
		fmt.Print("Biller: ", b, "\n")
	}

}