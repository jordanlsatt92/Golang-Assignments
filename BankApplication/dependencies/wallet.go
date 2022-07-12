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

// Finds an account by the account number and returns a pointer to that account
func (w *Wallet) FindAccountByAccountNum(num string) *BankAccount{
	for _,account := range w.Accounts{
		if account.AccountNum == num{
			return account
		}
	}
	return nil
}

// Adds an account to the wallet
func (w *Wallet) AddAccount(account *BankAccount) {
	w.Accounts = append(w.Accounts, account)
}

func GenerateAccountNumber() func() string {
	number:=1
	return func() string{
		temp:= number
		number++
		return strconv.Itoa(temp)
	}
}

var GenerateAccountNumberClosure func()string = GenerateAccountNumber()

// Creates an account and adds it to the wallet
// accountType is "checking", "saving", or "investment"
// Initial amount is the money that will be in the account
// The entity type is based on the entity type of the wallet owner
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

// Displays the accounts in the order specified
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

// Returns the total balance of all accounts
func (w *Wallet) Balance() float64{
	var totalBal float64
	for _,bal := range w.Accounts{
		totalBal += bal.Balance
	}
	return totalBal
}

// Wires money from the source account to the target account
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

func (w *Wallet) findAccountsWithFunds(amount float64) Wallet{
	var newWallet Wallet = Wallet{ID: "", Accounts:  make([]*BankAccount, 0)}
	for _,account := range w.Accounts{
		if account.Balance >= amount{
			newWallet.Accounts = append(newWallet.Accounts, account)
		}
	}
	return newWallet
}