package main

import (
	"BankApplicationModule/dependencies"
	"fmt"
	"os"
	"strings"
)

func main(){
	aOwner:= dependencies.Owner{"1","21 Liberty Street","Individual"}
	bAccount1 := dependencies.BankAccount{"1", aOwner, 300.00, 0.0125,"Checking"}
	bAccount2 := dependencies.BankAccount{"2", aOwner, 1000.00, 0.0125,"Saving"}
	bAccount3 := dependencies.BankAccount{"3", aOwner, 5000.00, 0.0125,"Investment"}
	wallet1 := dependencies.Wallet{ID: "W1", WalletOwner: aOwner}
	wallet1.AddAccount(&bAccount1)
	wallet1.AddAccount(&bAccount2)
	wallet1.AddAccount(&bAccount3)
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

	var input string
	var amount float64

	fmt.Println("The following program allows you to interact with your accounts.")
	fmt.Println("Type \"exit\" to quit.")
	for input != "exit"{
		
		fmt.Println("Your Wallet")
		wallet1.DisplayAccounts()

		fmt.Print("What would you like to do? Type \"deposit\", \"withdraw\", or \"wire\": ")
		fmt.Scan(&input)
		if strings.ToLower(input) == "exit"{
			os.Exit(0)
		}else if strings.ToLower(input) == "withdraw"{
			fmt.Print("Please enter the account number of the account from which you will withdraw: ")
			fmt.Scan(&input)
			account:= wallet1.FindAccountByAccountNum(input)
			fmt.Print("Enter the amount you would like to withdraw: ")
			fmt.Scan(&amount)
			account.Withdraw(amount)
		}else if strings.ToLower(input) == "deposit"{
			fmt.Print("Please enter the account number of the account to which you will deposit: ")
			fmt.Scan(&input)
			account:= wallet1.FindAccountByAccountNum(input)
			fmt.Print("Enter the amount you would like to deposit: ")
			fmt.Scan(&amount)
			account.Deposit(amount)			
		}else if strings.ToLower(input) == "wire"{
			fmt.Print("Please enter the source account number: ")
			fmt.Scan(&input)
			acc1:=wallet1.FindAccountByAccountNum(input)
			fmt.Print("Please enter the target account number: ")
			fmt.Scan(&input)
			acc2:=wallet1.FindAccountByAccountNum(input)
			fmt.Print("Enter the amount you would like to wire: ")
			fmt.Scan(&amount)
			acc1.Wire(acc1,acc2,amount)
		}


	}



}