package main

import (
	"ethereum/constants"
	store "ethereum/store"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//1.连接客户端
	client, err := ethclient.Dial(constants.BaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	//2.根据地址加载合约
	address := common.HexToAddress(constants.StoreAddress)
	instance, err := store.NewStore(address, client)
	if err != nil {
		log.Fatal(err)
	}

	//3.读取合约数据
	version, err := instance.Version(nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(version) // "1.0"
}
