package account

import (
	"fmt"
)

type Account struct {
	Amount    int64
	Activated bool
}

func OpenAccount(initialAmount int64) (openedAcc *Account) {
	if initialAmount > 0 {
		openedAcc := &Account{initialAmount, true}
		return openedAcc
	}

	return
}

func (acc *Account) CloseAccount() (payout int64, ok bool) {

	if acc.Activated {
		payout = acc.Amount
        fmt.Printf("The payout value is %v \n" ,payout)
		acc.Withdraw(acc.Amount)
		acc.Activated = false
		fmt.Println("The account has closed")
		ok = true
		return
	}
	ok = false

	fmt.Println("The account is already closed")
	return

}

func (acc *Account) Withdraw(amountToWithdraw int64) (ok bool) {
	if acc.Activated {
		newAmount := acc.Amount - amountToWithdraw
		if newAmount >= 0 {
			acc.Amount -= amountToWithdraw
			ok = true
			return
		} else {
			fmt.Println("Account has not the entered amount to withdraw")
			ok = false
			return
		}

	}
	fmt.Println("Withdraw on closed account is not available")
	ok = false
	return
}

func (acc *Account) Deposit(amountToDeposit int64) (ok bool) {

	if acc.Activated {
		newAmount := acc.Amount + amountToDeposit
		if newAmount >= 0 {
			acc.Amount += amountToDeposit
			ok = true
			return
		} else {
			fmt.Println("Account has not the entered amount to withdraw")
			ok = false
			return
		}

	}
	fmt.Println("account is closed")
	ok = false
	return
}

func PrintAccount(acc Account) {

	fmt.Printf("Balance is %v \n", acc.Amount)
	if acc.Activated {
		fmt.Println("Account is activated ")
	} else {
		fmt.Println("Account is closed")
	}
	fmt.Println("-----------------------------------------")
}
