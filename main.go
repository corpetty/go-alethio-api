package main

import (
	"fmt"

	"github.com/corpetty/go-alethio-api/alethio"
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

	client2 := alethio.NewClient(nil)
	accountTransactions, err := client2.GetAccountTransactions(address)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(accountTransactions)
}
