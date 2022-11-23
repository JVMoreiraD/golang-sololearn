package main

import "fmt"

type BankAccount struct {
	holder  string
	balance int
}

func (ptr *BankAccount) withdraw(amount int) {
	if ptr.balance < amount {
		fmt.Println("Insufficient Funds")
	} else {
		ptr.balance -= amount

	}
}

func main() {
	acc := BankAccount{"James Smith", 100000}

	var amount int
	fmt.Scanln(&amount)

	acc.withdraw(amount)
	fmt.Println(acc)
}
