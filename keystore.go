package main

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/ethereum/go-ethereum/accounts/keystore"
)

func main() {
	//createKs()
	importKs()
}

func createKs() {
	//1.指定目录创建keystore
	ks := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)

	//2.设定密码
	password := "secret"

	//3.创建账户
	account, err := ks.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}

	//4.打印地址
	fmt.Println("create Account", account.Address.Hex())
}

func importKs() {
	//1.读取keystore文件
	file := "./wallet/UTC--2023-03-24T14-29-43.690348600Z--c59e2c4de3a11677309bbed45d80ea53c3a3e458"
	ks := keystore.NewKeyStore("./tmp", keystore.StandardScryptN, keystore.StandardScryptP)
	jsonBytes, err := os.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	//2.相同的密码
	password := "secret"

	//3.导入账户（第三个参数是设置新的加密密码）
	account, err := ks.Import(jsonBytes, password, password)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("import Address", account.Address.Hex())

	//4.删除临时目录
	dir, err := os.ReadDir("./tmp")
	for _, d := range dir {
		os.RemoveAll(path.Join([]string{"./tmp", d.Name()}...))
	}
	os.Remove("./tmp")
}
