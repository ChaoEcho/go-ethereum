package main

import (
	"context"
	"ethereum/constants"
	"fmt"
	"log"
	"math"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {

	//1.连接以太坊节点
	client, err := ethclient.Dial(constants.BaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	//2.获取账户哈希值
	account := common.HexToAddress(constants.MainAccount)

	//3.获取账户余额
	balance, err := client.BalanceAt(context.Background(), account, nil)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balance)

	//4.获取特定区块号时的账户余额
	//  这里查看的是初始化时的余额
	blockNumber := big.NewInt(0)
	balanceAt, err := client.BalanceAt(context.Background(), account, blockNumber)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(balanceAt)

	//5.将数字转换为小数，之前默认的单位是wei
	fbalance := new(big.Float)
	fbalance.SetString(balance.String())
	ethValue := new(big.Float).Quo(fbalance, big.NewFloat(math.Pow10(18)))
	fmt.Println(ethValue) // 25.729324269165216041

	/*
		6.获取账户待处理的余额
		待确认余额（pending balance）则是指该账户已经发送了一些交易
		但这些交易还没有被确认打包进区块链中，因此这部分余额暂时无法被使用。
	*/
	pendingBalance, err := client.PendingBalanceAt(context.Background(), account)
	fmt.Println(pendingBalance)
}
