package main

import (
	"context"
	"ethereum/constants"
	"fmt"
	"log"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	//1.设置正则表达式的规则
	re := regexp.MustCompile("^0x[0-9a-fA-F]{40}$")

	//2.根据规则检查字符串是否符合规则
	fmt.Printf("is valid: %v\n", re.MatchString("0x323b5d4c32345ced77393b3530b1eed0f346429d")) // is valid: true
	fmt.Printf("is valid: %v\n", re.MatchString("0xZYXb5d4c32345ced77393b3530b1eed0f346429d")) // is valid: false

	//3.连接以太坊节点
	client, err := ethclient.Dial(constants.BaseUrl)
	//默认提供的测试节点，下面的合约地址就可以进行判断
	//client, err := ethclient.Dial("https://cloudflare-eth.com")
	if err != nil {
		log.Fatal(err)
	}

	//3.通过检查该地址是否存储了字节码，来判断是否是智能合约
	//	因为我这里还没有创建只能合约，所以两个都是false，后续可以修改地址来测试
	address := common.HexToAddress("0xe41d2489571d322189246dafa5ebde1f4699f498")
	bytecode, err := client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract := len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract) // is contract: true

	// a random user account address
	address = common.HexToAddress("0x8e215d06ea7ec1fdb4fc5fd21768f4b34ee92ef4")
	bytecode, err = client.CodeAt(context.Background(), address, nil) // nil is latest block
	if err != nil {
		log.Fatal(err)
	}

	isContract = len(bytecode) > 0

	fmt.Printf("is contract: %v\n", isContract) // is contract: false
}
