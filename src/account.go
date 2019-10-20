package main

import "fmt"

type Account struct {
	currentBalance int
}

type Transaction struct {
	transType   string
	transAmount int
}

func main() {
	var mainAccount Account
	mainAccount.currentBalance = 0

	var transType = [4]string{"D", "D", "W", "D"}
	var transAmount = [4]int{300, 300, 200, 100}

	for i, t := range transType {
		if t == "W" {
			mainAccount.currentBalance += -1 * transAmount[i]
		} else if t == "D" {
			mainAccount.currentBalance += transAmount[i]
		}
	}

	fmt.Println(mainAccount.currentBalance)
}
