package main

import (
	"context"
	"ethereum/constants"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//1.连接WebSocket
	client, err := ethclient.Dial(constants.WSUrl)
	if err != nil {
		log.Fatal(err)
	}

	//2.订阅新区块
	headers := make(chan *types.Header)
	sub, err := client.SubscribeNewHead(context.Background(), headers)
	if err != nil {
		log.Fatal(err)
	}

	//3.监听新区块
	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case header := <-headers:
			fmt.Println("监听到新区块:", header.Hash().Hex())

			block, err := client.BlockByHash(context.Background(), header.Hash())
			if err != nil {
				log.Fatal(err)
			}

			fmt.Println("Hash:", block.Hash().Hex())
			fmt.Println("Block Number:", block.Number().Uint64())
			fmt.Println("Block Time:", block.Time())
			fmt.Println("Block Nonce:", block.Nonce())
			fmt.Println("Block Transactions Len:", len(block.Transactions()))
		}
	}
}
