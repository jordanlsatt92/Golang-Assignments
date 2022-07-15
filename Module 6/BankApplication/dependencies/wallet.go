package dependencies

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Wallet struct {
	ID          string
	Accounts    []*BankAccount
	WalletOwner Owner
}

/*
@description: Function that finds an account in the wallet by the account number
@num (string): The account number of the account to be returned.
@return *BankAccount: A pointer to the bank account containing the bank account number.
*/
func (w *Wallet) FindAccountByAccountNum(num string) *BankAccount{
	for _,account := range w.Accounts{
		if account.AccountNum == num{
			return account
		}
	}
	return nil
}

/*
@description: Adds an account to the wallet
@param account (type BankAccount): The bank account to be added to the wallet.
*/
func (w *Wallet) AddAccount(account *BankAccount) {
	w.Accounts = append(w.Accounts, account)
}

/*
@description: Function that returns a function to generate account numbers for bank accounts
*/
func GenerateAccountNumber() func() string {
	number:=1
	return func() string{
		temp:= number
		number++
		return strconv.Itoa(temp)
	}
}

//Closure function to generate account numbers
var GenerateAccountNumberClosure func()string = GenerateAccountNumber()

/*
@desciption: Creates an account and adds it to the wallet
@param accountType (string): the account type to be created (checking,savings,investment)
@param initialAmount (float): the initial amount that will be deposited when the account is created.
*/
func (w *Wallet) CreateAccount(accountType string, initialAmount float64) {
	var interest float64
	if strings.ToLower(w.WalletOwner.EntityType) == "business"{
		if strings.ToLower(accountType) == "checking"{
			interest = 0.005
		}else if strings.ToLower(accountType) == "investment"{
			interest = 0.01
		}else{
			interest = 0.02
		}
	}else{
		if strings.ToLower(accountType) == "checking"{
			interest = 0.01
		}else if strings.ToLower(accountType) == "investment"{
			interest = 0.02
		}else{
			interest = 0.05
		}
	}
	
	newAccount:=BankAccount{GenerateAccountNumberClosure(),w.WalletOwner,initialAmount,interest, accountType}
	w.Accounts = append(w.Accounts, &newAccount)
}

/*
@description: Function that displays each accounts information that is in the wallet
*/
func (w *Wallet) DisplayAccounts() {
	accounts := make([]*BankAccount, len(w.Accounts))
	accounts = w.Accounts
	sort.Slice(accounts, func(i, j int) bool {
		return strings.ToLower(accounts[i].AccountType) < strings.ToLower(accounts[j].AccountType)
	})

	for _,val := range accounts{
		fmt.Printf("Account Num: %s, Account Owner: %s, Balance: %.2f, Int. Rate: %f, Type: %s\n", val.AccountNum, val.AccountOwner, val.Balance, val.InterestRate, val.AccountType)
	}
	
}

/*
@description: Function that wires the wireAmount param from the source account to the target account. If
the funds are not present in the source account, the function calls the findAccountsWithFunds()method 
and displays the returned value
@param source: the account from which the funds will come
@param target: the account to which the funds will go
@param wireamount: the dollar/float amount of money to be wired.
*/
func (w *Wallet) Balance() float64{
	var totalBal float64
	for _,bal := range w.Accounts{
		totalBal += bal.Balance
	}
	return totalBal
}

/*
@description: Function that wires the wireAmount param from the source account to the target account. If
the funds are not present in the source account, the function calls the findAccountsWithFunds()method 
and displays the returned value
@param source: the account from which the funds will come
@param target: the account to which the funds will go
@param wireamount: the dollar/float amount of money to be wired.
*/
func (w *Wallet) Wire(source, target *BankAccount, wireAmount float64) {
	if source.Balance < wireAmount{
		fmt.Printf("You do not have enough money in the source account to wire %.2f.\n", wireAmount)
		newWallet:= w.findAccountsWithFunds(wireAmount)
		if len(newWallet.Accounts) > 0{
			fmt.Println("Please try using on of the following accounts from which to wire\n")
			newWallet.DisplayAccounts()
		}else{
			fmt.Println("You do not have funds in any accounts to make this wire.\n")
		}

		return
	}else{
		target.Balance += wireAmount
		source.Balance -= wireAmount
	}
}
/*
@description: Function that searches the wallet and returns a wallet containing accounts that 
have at least as much as the amount parameter
@param amount: a dollar/float value
@return Wallet: a wallet that has all of the accounts that have at least as much as the amount parameter
*/
func (w *Wallet) findAccountsWithFunds(amount float64) Wallet{
	var newWallet Wallet = Wallet{ID: "", Accounts:  make([]*BankAccount, 0)}
	for _,account := range w.Accounts{
		if account.Balance >= amount{
			newWallet.Accounts = append(newWallet.Accounts, account)
		}
	}
	return newWallet
}