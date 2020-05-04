package main

import (
	"fmt"
	"log"

	"github.com/cuckooq/bank/accounts"
)

func main() {
	account := accounts.CreateAccount("cuckoo")
	account.Deposit(10)
	err := account.Withdraw(5)

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println(account.GetBalance())

	account.ChangeOwner("q")
	fmt.Println(account.GetOwner())

	fmt.Println(account)
}
