package main

import (
	"BankApplicationModule/dependencies"
	"fmt"
	"os"
	"strings"
)

func main(){
	aOwner:= dependencies.Owner{"1","21 Liberty Street","Individual"}
	wallet1 := dependencies.Wallet{ID: "W1", WalletOwner: aOwner}
	wallet1.CreateAccount("saving", 1000)
	wallet1.CreateAccount("checking", 500)
	wallet1.CreateAccount("investment", 5000)

	bOwner:= dependencies.Owner{"1","21 Liberty Street","Individual"}
	wallet2 := dependencies.Wallet{ID: "W1", WalletOwner: bOwner}
	wallet2.CreateAccount("saving", 20000)
	wallet2.CreateAccount("checking", 7500)
	wallet2.CreateAccount("investment", 50000)

	// The commented out section below was to test everything to make sure it worked properly

	// businessOwner:= dependencies.Owner{"2", "33 Business St", "Business"}
	// businessWallet:= dependencies.Wallet{ID:"W2", WalletOwner: businessOwner}
	// businessWallet.A
	// fmt.Println(bAccount1)
	// fmt.Println("Withdrawel")
	// bAccount1.Withdraw(20)
	// fmt.Println(bAccount1)
	// fmt.Println("Deposit")
	// bAccount1.Deposit(20)
	// fmt.Println(bAccount1)
	// fmt.Println("Wire")
	// bAccount1.Wire(&bAccount1,&bAccount2, 100)
	// fmt.Println(bAccount1)
	// fmt.Println(bAccount2)
	// fmt.Println("Display accounts in wallet")
	// wallet1.DisplayAccounts()
	// fmt.Println("Apply Interest")
	// bAccount1.ApplyInterest()
	// bAccount2.ApplyInterest()
	// bAccount3.ApplyInterest()
	// wallet1.DisplayAccounts()
	// fmt.Println("Total Balance")
	// fmt.Println("$",wallet1.Balance())

	// Variables to hold inputs
	var input string
	var amount float64

	fmt.Println("\n\nThe following program allows you to interact with your accounts.")
	fmt.Println("Type \"exit\" to quit.")
	// Runs until the user types exit
	for input != "exit"{
		
		fmt.Println("\n\nYour Wallet")
		wallet1.DisplayAccounts()

		//Ask user what they want to do
		fmt.Print("What would you like to do? Type \"new\" to create new account,\"deposit\", \"withdraw\", or \"wire\": ")
		fmt.Scan(&input)
		if strings.ToLower(input) == "exit"{
			os.Exit(0)
		// Withdraw
		}else if strings.ToLower(input) == "withdraw"{
			fmt.Print("Please enter the account number of the account from which you will withdraw: ")
			fmt.Scan(&input)
			account:= wallet1.FindAccountByAccountNum(input)
			fmt.Print("Enter the amount you would like to withdraw: ")
			fmt.Scan(&amount)
			account.Withdraw(amount)
		// Deposit
		}else if strings.ToLower(input) == "deposit"{
			fmt.Print("Please enter the account number of the account to which you will deposit: ")
			fmt.Scan(&input)
			account:= wallet1.FindAccountByAccountNum(input)
			fmt.Print("Enter the amount you would like to deposit: ")
			fmt.Scan(&amount)
			account.Deposit(amount)		
		// Wire money to an account	
		}else if strings.ToLower(input) == "wire"{
			fmt.Print("Would you like to wire to an account outside of your wallet? Type \"yes\" or \"no\": ")
			fmt.Scan(&input)
			// Wire money between accounts in user wallet
			if strings.ToLower(input) == "no"{
				fmt.Print("Please enter the source account number: ")
				fmt.Scan(&input)
				acc1:=wallet1.FindAccountByAccountNum(input)
				fmt.Print("Please enter the target account number: ")
				fmt.Scan(&input)
				acc2:=wallet1.FindAccountByAccountNum(input)
				fmt.Print("Enter the amount you would like to wire: ")
				fmt.Scan(&amount)
				acc1.Wire(acc1,acc2,amount)
			// Wire money between accounts in other wallet
			}else if strings.ToLower(input) == "yes"{
				fmt.Println("Here are the accounts you can wire to:")
				wallet2.DisplayAccounts()
				fmt.Print("Please enter the source account number from your wallet: ")
				fmt.Scan(&input)
				acc1:=wallet1.FindAccountByAccountNum(input)
				fmt.Print("Please enter the target account number: ")
				fmt.Scan(&input)
				acc2:=wallet2.FindAccountByAccountNum(input)
				fmt.Print("Enter the amount you would like to wire: ")
				fmt.Scan(&amount)
				wallet1.Wire(acc1,acc2,amount)
			}
		//Create new account
		}else if strings.ToLower(input) == "new"{
			fmt.Print("What type of account do you want to create? Type \"checking\", \"saving\", or \"investment\": ")
			fmt.Scan(&input)
			fmt.Print("Please enter the amount of the initial deposit (float): ")
			fmt.Scan(&amount)
			wallet1.CreateAccount(input, amount)
		}


	}



}