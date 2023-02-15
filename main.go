package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	ctx := context.Background()

	ethClient, err := ethclient.Dial("https://mainnet.infura.io/v3/")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	for {
		// Poll for new pending transactions
		pendingTxs, err := ethClient.PendingTransactionCount(ctx)
		if err != nil {
			log.Printf("Error while getting pending transactions: %v", err)
			break
		} else {
			fmt.Println("-------------------------------------------------------")
			fmt.Printf("Pending transaction: %v\n", pendingTxs)
		}
		// Sleep for a few seconds before checking for new pending transactions
		time.Sleep(time.Second * 5)
	}
}
