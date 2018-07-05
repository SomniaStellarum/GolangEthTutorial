package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
	"context"
	"encoding/json"
	"os"
)

var fltr ethereum.FilterQuery

func main() {
	// Setup client
	cl, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Panic("Client Dial Error: ", err)
	}

	// Setup filter and query node
	address := "0x61B59a28c621783caD9CB07b8c8b946f93E9B18E"
	fltr.Addresses = []common.Address{common.HexToAddress(address)}
	fltr.FromBlock = big.NewInt(int64(3570300))
	fltr.ToBlock = big.NewInt(int64(3570800))
	ctx := context.Background()
	lgs, err := cl.FilterLogs(ctx, fltr)
	if err != nil {
		log.Panic("Filter Error: ", err)
	}

	// Encode as json and output to Stdout
	b, err := json.Marshal(lgs)
	if err != nil {
		log.Panic("Marshal Error: ", err)
	}
	_, err = os.Stdout.Write(b)
	if err != nil {
		log.Panic("File Write Error: ", err)
	}
}