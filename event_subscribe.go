package main

import (
	"context"
	"ethereum/constants"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//1.连接Websocket客户端
	client, err := ethclient.Dial(constants.WSUrl)
	if err != nil {
		log.Fatal(err)
	}

	//2.获得ERC20合约，创建筛选查询
	contractAddress := common.HexToAddress(constants.ERC20Address)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{contractAddress},
	}

	//3.订阅事件
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
		}
	}
}
