package main

import (
	"crypto/ecdsa"
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
	"golang.org/x/crypto/sha3"
)

func main() {

	//1.生成私钥
	privateKey, err := crypto.GenerateKey()
	if err != nil {
		log.Fatal(err)
	}

	//2.私钥转换为字节
	privateKeyBytes := crypto.FromECDSA(privateKey)

	//3.这样输出没有0x
	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])

	//4.根据私钥生成公钥
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	/*
		5.公钥转换为字节
		剥离了0x和前2个字符04，它始终是EC前缀，不是必需的
	*/
	publicKeyBytes := crypto.FromECDSAPub(publicKeyECDSA)
	//fmt.Println(publicKeyBytes)
	//fmt.Println(hexutil.Encode(publicKeyBytes))
	fmt.Println(hexutil.Encode(publicKeyBytes)[4:])

	//6.生成地址
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	fmt.Println(address)

	//7.公钥哈希之后取最后的20个字节
	hash := sha3.NewLegacyKeccak256()
	//这里从1开始，是因为是字符数组，就相当于转换为Hex后的04
	hash.Write(publicKeyBytes[1:])
	fmt.Println(hexutil.Encode(hash.Sum(nil)[12:])) // 0x96216849c49358b10257cb55b28ea603c874b05e
}
