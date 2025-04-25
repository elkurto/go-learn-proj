package main

import (
	"fmt"
)

func printMenuAndGetResponse() int {
	var menuChoice = 0
	var err error = nil

	fmt.Print(`Welcome - To a bank simulator:
    Please choose an option from the menu;
	  1. Check Balance
	  2. Deposit Funds
	  3. Withdraw Funds
	  4. Exit
    Your choice: `)
	_, err = fmt.Scan(&menuChoice)
	if err != nil {
		menuChoice = 999
	}
	return menuChoice
}
func runBankLoopSimulation() {

	var bRunning = true
	var menuChoice = 0
	var balance = 1000

	for bRunning {

		// 1. prompt
		menuChoice = printMenuAndGetResponse()

		// 2. process choice
		switch menuChoice {
		case 1:
			checkBalance(balance)
			continue
		case 2:
			balance = depositMoney(balance)
			continue
		case 3:
			balance = withdrawMoney(balance)
			continue
		case 4:
			bRunning = false
			fmt.Println("Exiting... : Have a nice day.")
			continue
		default:
			complainInvalidChoice()
		}
	}
}

func checkBalance(balance int) int {
	fmt.Printf("current balance: %d  ****\n\n", balance)
	return balance
}
func withdrawMoney(balance int) int {
	var amount int
	var err error

	fmt.Print("\nAmount (Withdraw): ")
	_, err = fmt.Scan(&amount)
	if err != nil {
		fmt.Printf("Invalid amount specified - type an amount in [0, %d]\n\n", balance)
	} else if amount < 0 || balance < amount {
		fmt.Printf("Invalid amount specified - type an amount in [0, %d]\n\n", balance)
	} else {
		balance -= amount
		fmt.Printf("Here's is your money: %d  ****\n", amount)
		fmt.Printf("Current Balance after withdrawl: %d    ****\n\n", balance)
	}
	return balance
}
func depositMoney(balance int) int {
	var amount int
	var err error

	fmt.Print("\nAmount (Deposit): ")
	_, err = fmt.Scan(&amount)
	if err != nil {
		fmt.Printf("Invalid amount specified - type an amount in [0, %d]\n", balance)
	} else {
		balance += amount
		fmt.Printf("Your deposit amount = %d  ****\n", amount)
		fmt.Printf("Current Balance after deposit: %d    ****\n\n", balance)
	}
	return balance
}
func complainInvalidChoice() {
	fmt.Println("Invalid option specified - Try again")
}

func main() {
	runBankLoopSimulation()
}
