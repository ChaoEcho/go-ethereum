package main

import (
	"context"
	"ethereum/constants"
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	block()
}

func block() {
	var (
		ctx = context.Background()
		//url         = "https://cloudflare-eth.com"
		url         = constants.BaseUrl
		client, err = ethclient.DialContext(ctx, url)
	)

	//1.nil默认查询最新区块的信息
	header, err := client.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("latest block number =", header.Number.String())

	//2.查询指定区块的信息
	blockNumber := big.NewInt(1)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	//3.输出一些详细信息
	fmt.Println("block number =", block.Number().Uint64())
	fmt.Println("block timestamp =", block.Time())
	fmt.Println("block difficulty =", block.Difficulty().Uint64())
	fmt.Println("block hash =", block.Hash().Hex())
	fmt.Println("block transaction count =", len(block.Transactions()))

	//4.查询指定区块的交易数量
	count, err := client.TransactionCount(context.Background(), block.Hash())
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("block transaction count =", count)
}
