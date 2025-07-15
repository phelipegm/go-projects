package main

import (
	"fmt"
	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "Bank/balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("-------------")
		//panic(err)
	}

	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Today is", randomdata.Day())
	
	for {
			presentOptions()
			
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
				fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
				fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			default:
				fmt.Println("\nGoodbye!")
				return
		}
	}
}