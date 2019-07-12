package main

import (
	"fmt"

	"github.com/corpetty/go-alethio-api/alethio"
	"github.com/davecgh/go-spew/spew"
)

func main() {
	var address = "0x2942577508e060ea092c0CD7802ae42c1CEA2BAe"
	client := alethio.NewClient(nil)

	accountDetails, err := client.GetAccountDetails(address)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(accountDetails)

	// var contract = "0xB8c77482e45F1F44dE1745F52C74426C631bDD52"
	// client2 := alethio.NewClient(nil)
	// contractDetails, err := client2.GetAccountContract(contract)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// fmt.Println(contractDetails)

	// client2 := alethio.NewClient(nil)
	// accountTransactions, err := client2.GetAccountTransactions(address)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// fmt.Println(accountTransactions)

	// client2 := alethio.NewClient(nil)
	// blocks, err := client2.GetAllBlocks()
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// fmt.Println(blocks)

	// client2 := alethio.NewClient(nil)
	// var blockhash = "0x05245795a1b3ed9486387102539b55f71527f7cbe9afd3918642ed46735c30f0"
	// block, err := client2.GetBlockByHash(blockhash)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// fmt.Println(block)

	// client2 := alethio.NewClient(nil)
	// var blockNumber = 43243
	// block, err := client2.GetBlockByNumber(blockNumber)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// fmt.Println(block)

	// client2 := alethio.NewClient(nil)
	// var label = "latest"
	// block, err := client2.GetBlockByLabel(label)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// fmt.Println(block)

	// client2 := alethio.NewClient(nil)
	// var blockhash = "0x05245795a1b3ed9486387102539b55f71527f7cbe9afd3918642ed46735c30f0"
	// block, err := client2.GetBlockParent(blockhash)
	// if err != nil {
	// 	fmt.Print(err)
	// }
	// spew.Dump(block)

	client2 := alethio.NewClient(nil)
	var blockhash = "0x05245795a1b3ed9486387102539b55f71527f7cbe9afd3918642ed46735c30f0"
	account, err := client2.GetBlockBeneficiary(blockhash)
	if err != nil {
		fmt.Print(err)
	}
	spew.Dump(account)
}
