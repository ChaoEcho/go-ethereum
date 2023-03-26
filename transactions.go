package main

import (
	"context"
	"encoding/hex"
	"ethereum/constants"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	transaction()
}

func transaction() {

	//1.连接客户端
	var (
		ctx         = context.Background()
		url         = constants.BaseUrl
		client, err = ethclient.DialContext(ctx, url)
	)
	if err != nil {
		log.Fatal(err)
	}

	//2.查询指定区块
	blockNumber := big.NewInt(1)
	block, err := client.BlockByNumber(context.Background(), blockNumber)
	if err != nil {
		log.Fatal(err)
	}

	//3.查询区块的交易
	transactions := block.Transactions()
	fmt.Println("total transactions count:", len(transactions))

	//4.遍历交易
	for ind, tx := range transactions {
		fmt.Println("----", ind, "----")
		fmt.Println("transaction hash:", tx.Hash().Hex())
		fmt.Println("transaction value:", tx.Value().String())
		fmt.Println("transaction gas limit:", tx.Gas())
		fmt.Println("transaction fee cap per gas:", tx.GasFeeCap())
		fmt.Println("transaction tip cap per gas:", tx.GasTipCap())
		fmt.Println("transaction gas price:", tx.GasPrice())
		fmt.Println("transaction nonce:", tx.Nonce())                   // 110644
		fmt.Println("transaction data:", hex.EncodeToString(tx.Data())) // []
		fmt.Println("transaction to:", tx.To().Hex())                   // 0x55fE59D8Ad77035154dDd0AD0388D09Dd4047A8e

		//4.1.获取交易的链ID
		chainID, err := client.NetworkID(context.Background())
		fmt.Println("transaction chain id:", chainID)
		if err != nil {
			log.Fatal(err)
		}

		//4.2.获取交易的发送者
		//	这里协议发生改变了，所以使用最新的
		if msg, err := tx.AsMessage(types.LatestSignerForChainID(chainID), tx.GasPrice()); err == nil {
			fmt.Println("transaction from:", msg.From().Hex())
		}

		//4.3.根据事务的哈希查询凭证
		receipt, err := client.TransactionReceipt(context.Background(), tx.Hash())
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("transaction gas used:", receipt.GasUsed)
		fmt.Println("transaction status:", receipt.Status)
	}

	//5.根据块的哈希值查询特定的交易
	blockHash := common.HexToHash("0x850e12bab6eb375f2851bda73fa23e1596bc455e2854c263eedcc068947a3916")
	count, err := client.TransactionCount(context.Background(), blockHash)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("----block 1----")
	for idx := uint(0); idx < count; idx++ {
		tx, err := client.TransactionInBlock(context.Background(), blockHash, idx)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("transaction hash:", tx.Hash().Hex())
	}

	//6.获取最新的区块交易
	fmt.Println("----last ethw transaction----")
	lastBlock, err := client.BlockByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}
	tx, isPending, err := client.TransactionByHash(context.Background(), lastBlock.Transactions()[0].Hash())
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("transaction hash:", tx.Hash().Hex())
	fmt.Println("transaction gas limit:", tx.Gas())
	fmt.Println("isPending:", isPending) // false
}
