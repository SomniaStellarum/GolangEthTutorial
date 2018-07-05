package main

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"context"
	"encoding/json"
	"os"
)

func main() {
	// Setup client
	cl, err := ethclient.Dial("https://ropsten.infura.io")
	if err != nil {
		log.Panic("Client Dial Error: ", err)
	}

	// Fetch Block
	ctx := context.Background()
	i := int64(3563639)
	block, err := cl.BlockByNumber(ctx, big.NewInt(i))
	if err != nil {
		log.Panic("Error Fetching Block: ", err)
	}

	// Encode block as json and write to os.Stout
	v, err := json.Marshal(block.Header())
	if err != nil {
		log.Panic("Marshal Error: ", err)
	}
	_, err = os.Stdout.Write(v)
	if err != nil {
		log.Panic("Error Writing Logs")
	}
}