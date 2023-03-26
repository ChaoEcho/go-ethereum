package main

import (
	"ethereum/constants"
	"log"
	"math/big"

	"github.com/davecgh/go-spew/spew"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/miguelmota/go-ethereum-hdwallet"
)

func main() {
	//1.助记词生成钱包
	mnemonic := "tag volcano eight thank tide danger coast health above argue embrace heavy"
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		log.Fatal(err)
	}

	//2.根据派生路径来生成账户
	path := hdwallet.MustParseDerivationPath("m/44'/60'/0'/0/0")
	account, err := wallet.Derive(path, true)
	if err != nil {
		log.Fatal(err)
	}

	//3.定义一个交易
	nonce := uint64(0)
	value := big.NewInt(10000000000000000)
	toAddress := common.HexToAddress(constants.TransactionAddress)
	gasLimit := uint64(2100000)
	gasPrice := big.NewInt(21000)
	var data []byte

	tx := types.NewTransaction(nonce, toAddress, value, gasLimit, gasPrice, data)
	signedTx, err := wallet.SignTx(account, tx, nil)
	if err != nil {
		log.Fatal(err)
	}

	//4.打印交易
	spew.Dump(signedTx)

	//5.发送交易（这里会失败，因为是新创建的账户，没有金额，所以汽油费不足）
	/*client, err := ethclient.Dial(constants.BaseUrl)
	if err != nil {
		log.Fatal(err)
	}

	err = client.SendTransaction(context.Background(), signedTx)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("tx sent: %s", signedTx.Hash().Hex())*/
}
