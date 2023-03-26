package main

import (
	"context"
	"encoding/hex"
	"ethereum/constants"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rlp"
)

func main() {
	//1.连接客户端
	client, err := ethclient.Dial(constants.BaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	//2.原始编码
	rawTx := "f86e078504a817c8008252089409e0e9cfce1338c7c0a86f8b125d8d7d7d0010dd880de0b6b3a764000080820a96a009eba198855c9c953af14dff5f57840fa658ec80b6ae4221239e608ea7fef9fda05e827fdae11a567230cacd2f90795ef1fd65ff9ed76390dbe7311091930dafe9"

	//3.解码
	rawTxBytes, err := hex.DecodeString(rawTx)
	tx := new(types.Transaction)
	rlp.DecodeBytes(rawTxBytes, &tx)

	//4.发送交易
	err = client.SendTransaction(context.Background(), tx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", tx.Hash().Hex()) // tx sent: 0xc429e5f128387d224ba8bed6885e86525e14bfdc2eb24b5e9c3351a1176fd81f
}
