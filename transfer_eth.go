package main

import (
	"context"
	"crypto/ecdsa"
	"ethereum/constants"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	//1.连接以太坊客户端
	client, err := ethclient.Dial(constants.BaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	//2.加载私钥
	privateKey, err := crypto.HexToECDSA(constants.MainPrivateKey[2:])
	if err != nil {
		log.Fatal(err)
	}

	//3.获取公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	//4.发送方地址
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)

	//5.获取nonce
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	//6.设置发送数量以及汽油费
	value := big.NewInt(1000000000000000000) // in wei (1 eth)
	gasLimit := uint64(42000)                // in units
	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	//7.接收方地址
	toAddress := common.HexToAddress(constants.TransactionAddress)

	//8.创建交易
	//data := []byte("hello world")
	var data []byte
	tx := types.NewTx(&types.LegacyTx{Nonce: nonce, GasPrice: gasPrice, Gas: gasLimit, To: &toAddress, Value: value, Data: data})

	//9.获取链ID，并签名交易
	chainID, err := client.NetworkID(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	signedTx, err := types.SignTx(tx, types.LatestSignerForChainID(chainID), privateKey)
	if err != nil {
		log.Fatal(err)
	}

	//10.发送交易
	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())
}
