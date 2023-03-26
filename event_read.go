package main

import (
	"context"
	"ethereum/constants"
	store "ethereum/store"
	"fmt"
	"log"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//1.连接Websocket客户端
	client, err := ethclient.Dial(constants.WSUrl)
	if err != nil {
		log.Fatal(err)
	}

	//2.获得合约，创建筛选查询
	contractAddress := common.HexToAddress(constants.StoreAddress)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(0),
		ToBlock:   big.NewInt(1000),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	//3.获得所有匹配的日志
	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	//4.abi解析日志
	contractAbi, err := abi.JSON(strings.NewReader(string(store.StoreABI)))
	if err != nil {
		log.Fatal(err)
	}

	//5.遍历日志
	for _, vLog := range logs {
		fmt.Printf("BlockHash: %s\n", vLog.BlockHash.Hex())
		fmt.Printf("BlockNumber: %d\n", vLog.BlockNumber)
		fmt.Printf("TxHash: %s\n", vLog.TxHash.Hex())

		event := struct {
			Key   [32]byte
			Value [32]byte
		}{}
		err := contractAbi.UnpackIntoInterface(&event, "ItemSet", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Event: %s\n", event)

		//5.1.遍历日志的主题
		var topics [4]string
		for i := range vLog.Topics {
			topics[i] = vLog.Topics[i].Hex()
		}
		//5.2.第一个主题总是事件的签名，首个主题只是被哈希过的事件签名
		fmt.Printf("Topic[0]: %s\n", topics[0])
	}

	//6.验证首个主题是否是被哈希过的事件签名
	eventSignature := []byte("ItemSet(bytes32,bytes32)")
	hash := crypto.Keccak256Hash(eventSignature)
	fmt.Println(hash.Hex())
}
