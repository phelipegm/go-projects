package main

import (
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "Bank/balance.txt"

func getBalanceFromFile() float64 {
	data, _ := os.ReadFile(accountBalanceFile)
	balanceText := string(data)
	balance, _ := strconv.ParseFloat(balanceText, 64)
	return balance
}

func writeBalanceToFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(balanceText), 0644)
}

func main() {
	var accountBalance = getBalanceFromFile()

	fmt.Println("Welcome to Go Bank!")
	
	for {
			fmt.Println("\nWhat do you want to do?")
			fmt.Println("1. Check balance")
			fmt.Println("2. Deposit money")
			fmt.Println("3. Withdraw money")
			fmt.Println("4. Exit")

			var choice int
			fmt.Print("\nYour choice: ")
			fmt.Scan(&choice)

		switch choice {
			case 1:
				fmt.Println("\nYour balance is:", accountBalance)
			case 2:
				fmt.Print("\nYour deposit: ")
				var depositAmout float64
				fmt.Scan(&depositAmout)

				if depositAmout <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					continue
				}

				accountBalance += depositAmout
				fmt.Println("Balance updated! New amount:", accountBalance)
				writeBalanceToFile(accountBalance)
			case 3:
				fmt.Print("\nWithdrawal amount: ")
				var withdrawalAmount float64
				fmt.Scan(&withdrawalAmount)

				if withdrawalAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					continue
				}

				if withdrawalAmount > accountBalance {
					fmt.Println("Invalid amount. You can't withdraw more than you have.")
					continue
				}

				accountBalance -= withdrawalAmount
				fmt.Println("Balance updated! New amount:", accountBalance)
				writeBalanceToFile(accountBalance)
			default:
				fmt.Println("\nGoodbye!")
				return
		}
	}
}