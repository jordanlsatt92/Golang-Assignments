package dependencies

import (
	"errors"
	"fmt"
	"sort"
	"strings"
)

type Wallet struct {
	ID          string
	Accounts    []*BankAccount
	WalletOwner Owner
}

func (w *Wallet) FindAccountByAccountNum(num string) *BankAccount{
	for _,account := range w.Accounts{
		if account.AccountNum == num{
			return account
		}
	}
	return nil
}

func (w *Wallet) AddAccount(account *BankAccount) {
	w.Accounts = append(w.Accounts, account)
}

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

func (w *Wallet) Balance() float64{
	var totalBal float64
	for _,bal := range w.Accounts{
		totalBal += bal.Balance
	}
	return totalBal
}

func (w *Wallet) Wire(source, target *BankAccount, wireAmount float64) error{
	if source.Balance < wireAmount{
		return errors.New("you cannot wire more money than you have in the account")
	}else{
		target.Balance += wireAmount
		source.Balance -= wireAmount
		return nil
	}
}