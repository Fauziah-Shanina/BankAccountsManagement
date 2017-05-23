package main

import (
	"account"
	"time"
	//"fmt"
)

func mobileAccesInstructions(commandChanel chan account.Account) {

	//open account
	openedAccount := account.OpenAccount(20000)
	account.PrintAccount(*openedAccount)
	commandChanel <- *openedAccount
	time.Sleep(time.Millisecond * 10)

	//deposit amount
	openedAccount.Deposit(5000)
	account.PrintAccount(*openedAccount)
	commandChanel <- *openedAccount
	time.Sleep(time.Millisecond * 10)

	//close account
	openedAccount.CloseAccount()
	account.PrintAccount(*openedAccount)
	commandChanel <- *openedAccount
	time.Sleep(time.Millisecond * 10)

}

func webAccesInstructions(commandChanel chan account.Account) {
	openedAccount := <-commandChanel
	//withdraw amount

	openedAccount.Withdraw(60000)
	account.PrintAccount(openedAccount)
	time.Sleep(time.Millisecond * 40)

}

func main() {

	commandResult := make(chan account.Account)
	go mobileAccesInstructions(commandResult)
	webAccesInstructions(commandResult)
}
