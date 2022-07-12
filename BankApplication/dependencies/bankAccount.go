package dependencies

import (
	"fmt"
	"strings"
)

type BankAccount struct {
	AccountNum   string
	AccountOwner Owner
	Balance      float64
	InterestRate float64
	AccountType  string
}

// Withdraws the float "value" from the account
func (a *BankAccount) Withdraw(value float64) {
	if value > a.Balance{
		fmt.Println("ERROR\nYou cannot withdraw more than you have in the account\nERROR")
		return
	}else {
		a.Balance -= value
	}
}

// Deposits the float "value" into the account
func (a *BankAccount) Deposit(value float64) {
	if value < 0{
		fmt.Println("ERROR\nyou cannot deposit a negative ammount\nERROR")
		return
	}else{
		a.Balance += value
	}
}

// Applies the interest specified to each account based on the type of account
func (a *BankAccount) ApplyInterest(){
	// Individual Accounts
	if strings.ToLower(a.AccountOwner.EntityType) == "individual"{
		if strings.ToLower(a.AccountType) == "checking"{
			a.Balance = 1.01 * a.Balance
		}
		if strings.ToLower(a.AccountType) == "investment"{
			a.Balance = 1.02 * a.Balance
		}
		if strings.ToLower(a.AccountType) == "saving"{
			a.Balance = 1.05 * a.Balance
		}
	}
	// Business Accounts
	if strings.ToLower(a.AccountOwner.EntityType) == "business"{
		if strings.ToLower(a.AccountType) == "checking"{
			a.Balance = 1.005 * a.Balance
		}
		if strings.ToLower(a.AccountType) == "investment"{
			a.Balance = 1.01 * a.Balance
		}
		if strings.ToLower(a.AccountType) == "saving"{
			a.Balance = 1.02 * a.Balance
		}
	}
}

// Wires money from the source accound to the Target account
func (a BankAccount) Wire(source, target *BankAccount, wireAmount float64) {
	if a.Balance < wireAmount{
		fmt.Println("ERROR\nyou cannot wire more money than you have in the account\nERROR")
		return
	}else{
		target.Balance += wireAmount
		source.Balance -= wireAmount

	}
}

