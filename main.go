package main

import (
	"fmt"

	"github.com/corpetty/go-alethio-api/alethio"
)

func main() {
	var address = "0x2942577508e060ea092c0CD7802ae42c1CEA2BAe"
	client := alethio.NewClient(nil)

	account, err := client.GetAccount(address)
	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(account)
}
