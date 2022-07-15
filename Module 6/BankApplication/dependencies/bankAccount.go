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

/*
@description: Function that withdraws money from the bank account
@param value (float): The value to be withdrawn
*/
func (a *BankAccount) Withdraw(value float64) {
	if value > a.Balance{
		fmt.Println("ERROR\nYou cannot withdraw more than you have in the account\nERROR")
		return
	}else {
		a.Balance -= value
	}
}

/*
@description: Function that deposits an amount into the account
@param value (float): the amount to be added to the account.
*/
func (a *BankAccount) Deposit(value float64) {
	if value < 0{
		fmt.Println("ERROR\nyou cannot deposit a negative ammount\nERROR")
		return
	}else{
		a.Balance += value
	}
}

/*
@description: Function that applies the interest to the bank account depending on whether
it is owned by a business or individual and the account type (checking,savings,investment)
*/
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

/*
@description: Function that wires the wireAmount param from the source account to the target account. If
the funds are not present in the source account, the function prints an error message.
@param source (*BankAccount): the account from which the funds will come
@param target (*BankAccount): the account to which the funds will go
@param wireamount: the dollar/float amount of money to be wired.
*/
func (a BankAccount) Wire(source, target *BankAccount, wireAmount float64) {
	if a.Balance < wireAmount{
		fmt.Println("ERROR\nyou cannot wire more money than you have in the account\nERROR")
		return
	}else{
		target.Balance += wireAmount
		source.Balance -= wireAmount

	}
}

