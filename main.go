package LandTradingContract

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
)

func main() {
	client, err := ethclient.Dial("http://127.0.0.1:8545")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	ctx := context.Background()
	blockNumber, err := client.BlockNumber(ctx)
	if err != nil {
		log.Fatalf("Failed to get block number: %v", err)
	}
	fmt.Println("Connected to Ethereum client, latest block number:", blockNumber)
}
