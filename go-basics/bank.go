package main

import (
	"fmt"

	"example.com/bank/fileops"
	"github.com/Pallinder/go-randomdata"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)

	if err != nil {
		fmt.Println("| ----- Failed to get balance from file ----- |")
		fmt.Println(err)
		// panic("Cannot continue without balance") // Stop the program
		return
	}

	fmt.Println("Welcome to the Bank!")
	fmt.Println("Reach us on: ", randomdata.PhoneNumber())
	for {
		presentOptions() // Call the function to display options from communication.go file

		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		// Boolean variables to store user's choice
		// wantsCheckBalance := choice == 1
		// wantsDeposit := choice == 2
		// wantsWithdraw := choice == 3
		// wantsExit := choice == 4

		// if wantsCheckBalance {sw

		switch choice {
		case 1:
			fmt.Println("Your balance is:", accountBalance)
		case 2:
			fmt.Print("Your deposit amount:")
			var depositAmount float64
			fmt.Scan(&depositAmount)

			// Check if deposit amount is negative
			if depositAmount < 0 {
				fmt.Println("You cannot deposit a negative amount")
				// return
				continue // Skip the rest of the loop
			}

			accountBalance += depositAmount
			fmt.Println("Your new balance is:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		case 3:
			fmt.Print("Your withdraw amount:")
			var withdrawAmount float64
			fmt.Scan(&withdrawAmount)

			// Check if withdraw amount is negative
			if withdrawAmount < 0 {
				fmt.Println("You cannot withdraw a negative amount")
				continue
			}
			// Check if withdraw amount is more than account balance
			if withdrawAmount > accountBalance {
				fmt.Println("You cannot withdraw more than your balance")
				continue
			}

			accountBalance -= withdrawAmount
			fmt.Println("Your new balance is:", accountBalance)
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thank you for using the Bank!")
			return
		}

		// ======= IF ELSE VERSION =======
		/*
			if choice == 1 {
				fmt.Println("Your balance is:", accountBalance)
			} else if choice == 2 {
				fmt.Print("Your deposit amount:")
				var depositAmount float64
				fmt.Scan(&depositAmount)

				// Check if deposit amount is negative
				if depositAmount < 0 {
					fmt.Println("You cannot deposit a negative amount")
					// return
					continue // Skip the rest of the loop
				}

				accountBalance += depositAmount
				fmt.Println("Your new balance is:", accountBalance)
			} else if choice == 3 {
				fmt.Print("Your withdraw amount:")
				var withdrawAmount float64
				fmt.Scan(&withdrawAmount)

				// Check if withdraw amount is negative
				if withdrawAmount < 0 {
					fmt.Println("You cannot withdraw a negative amount")
					return
				}
				// Check if withdraw amount is more than account balance
				if withdrawAmount > accountBalance {
					fmt.Println("You cannot withdraw more than your balance")
					return
				}

				accountBalance -= withdrawAmount
				fmt.Println("Your new balance is:", accountBalance)
				} else {
					fmt.Println("Goodbye!")
					// return
					break // Exit the loop
					}

		*/
		// ======= END IF ELSE VERSION =======
	}
}
