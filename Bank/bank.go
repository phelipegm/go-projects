package main

import "fmt"

func main() {
	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")

	var choice int
	fmt.Print("\nYour choice: ")
	fmt.Scan(&choice)

	if choice == 1 {
		fmt.Println("\nYour balance is:", accountBalance)
	} else if choice == 2 {
		fmt.Print("\nYour deposit: ")
		var depositAmout float64
		fmt.Scan(&depositAmout)

		if depositAmout <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0.")
			return
		}

		accountBalance += depositAmout
		fmt.Println("Balance updated! New amount:", accountBalance)
	} else if choice == 3 {
		fmt.Print("\nWithdrawal amount: ")
		var withdrawalAmount float64
		fmt.Scan(&withdrawalAmount)

		if withdrawalAmount <= 0 {
			fmt.Println("Invalid amount. Must be greater than 0.")
			return
		}

		if withdrawalAmount > accountBalance {
			fmt.Println("Invalid amount. You can't withdraw more than you have.")
			return
		}

		accountBalance -= withdrawalAmount
		fmt.Println("Balance updated! New amount:", accountBalance)
	} else {
		fmt.Println("\nGoodbye!")
	}

}