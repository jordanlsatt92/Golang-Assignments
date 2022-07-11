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

func (a *BankAccount) Withdraw(value float64) {
	if value > a.Balance{
		fmt.Println("ERROR\nYou cannot withdraw more than you have in the account\nERROR")
	}else {
		a.Balance -= value
	}
}

func (a *BankAccount) Deposit(value float64) {
	if value < 0{
		fmt.Println("ERROR\nyou cannot deposit a negative ammount\nERROR")
	}else{
		a.Balance += value
	}
}

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

func (a BankAccount) Wire(source, target *BankAccount, wireAmount float64) {
	if a.Balance < wireAmount{
		fmt.Println("ERROR\nyou cannot wire more money than you have in the account\nERROR")
	}else{
		target.Balance += wireAmount
		source.Balance -= wireAmount

	}
}