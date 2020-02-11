package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	_ "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	. "github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client := ConnectClient("https://rinkeby.infura.io/v3/8e2834b158fa48b0a5fb9ca0f72ce6e6")

	// Get Block Number
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	blockNumber := header.Number.String()
	fmt.Println(blockNumber)

	blockTest := big.NewInt(5952590)

	fmt.Println(blockNumber)

	block, err := client.BlockByNumber(context.Background(), blockTest)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(block.Hash().Hex())
}

// Connexion à un noeud geth Rinkeby via Infura
func ConnectClient(url string) *Client {
	client, err := ethclient.Dial(url) //"rinkeby.infura.io/v3/8e2834b158fa48b0a5fb9ca0f72ce6e6"
	if err != nil {
		fmt.Println("Error while connecting to infura")
		log.Fatal(err)
	}
	return client
}
