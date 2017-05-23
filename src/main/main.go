package main

import (
	"account"
	"fmt"
	"time"
)

func mobileAccesInstructions(commandChanel chan account.Account) {

	//open account
	fmt.Println("Open account from mobile")
	openedAccount := account.OpenAccount(20000)
	account.PrintAccount(*openedAccount)
	commandChanel <- *openedAccount
	time.Sleep(time.Millisecond * 10)

	//deposit amount
	fmt.Println("Deposit to account from mobile")
	openedAccount.Deposit(5000)
	account.PrintAccount(*openedAccount)
	commandChanel <- *openedAccount
	time.Sleep(time.Millisecond * 10)

	//close account
	fmt.Println("Close account from mobile")
	openedAccount.CloseAccount()
	account.PrintAccount(*openedAccount)
	commandChanel <- *openedAccount
	time.Sleep(time.Millisecond * 10)

}

func webAccesInstructions(commandChanel chan account.Account) {
	//withdraw amount
	openedAccount := <-commandChanel
	fmt.Println("Withdraw from website ")

	openedAccount.Withdraw(60000)
	account.PrintAccount(openedAccount)
	time.Sleep(time.Millisecond * 10)

}

func main() {

	commandResult := make(chan account.Account,3)
	go mobileAccesInstructions(commandResult)
	go webAccesInstructions(commandResult)
	time.Sleep(time.Millisecond * 5000)
}
