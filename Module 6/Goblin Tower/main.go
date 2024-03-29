package main

import (
	"GoblinModule/dependencies"
	"fmt"
	"math/rand"
	"time"
)


func main(){
	
	hero:=dependencies.GenerateHero()
	fmt.Println("Welcome to Goblin Tower. Try not to die, Hero.")
	keepPlaying:=true
	heroIsDead:=false
	totalSteps:=0
	var input string
	for keepPlaying{
		// Hero takes 10 steps before advancing to new level
		for i:=1; i <= 10; i++{
			fmt.Println("Your taking a step.")
			totalSteps++
			rand.Seed(time.Now().UnixNano())
			chance:= rand.Intn(10+1-1)+1
			// Hero encounters goblin by chance and they fight
			if chance > 5{
				goblin:=dependencies.GenerateGoblin()
				fmt.Println("You've encountered a goblin. Fight!")
				for goblin.Health > 0 && hero.CurrentHealth > 0{
					fmt.Println("\nYour Stats:")
					hero.DisplayStats()
					fmt.Println("Goblin Stats:")
					goblin.DisplayStats()
					fmt.Println("What would you like to do? Type the number of the choice below: ")
					fmt.Println("1. Attack\n2. Take potion")
					fmt.Scan(&input)
					//User chooses to attack
					if input == "1"{
						hero.Attack(&goblin)
						fmt.Println("You attacked!")
					// User chooses to take potion
					}else if input == "2"{
						err:= hero.UsePotion()
						// if the hero does not have any potions, an error is returned and the user must attack
						if err != nil{
							fmt.Println(err)
							hero.Attack(&goblin)
							fmt.Println("You attacked!")
						}
					}else{
						continue
					}
					tookDamage:=goblin.Attack(&hero)
					if tookDamage{
						fmt.Println("Goblin attacked!")
					}else{
						fmt.Println("You blocked the goblin's attack!")
					}
				}
				// Hero defeated goblin
				if hero.CurrentHealth > 0{
					fmt.Println("\nYou defeated the goblin!")
					hero.Gold += 2
					hero.GoblinsSlain ++
				// Hero was killed by goblin; prompt user to see if they want to play again
				}else{
					fmt.Println("You've been defeated.")
					heroIsDead = true
					fmt.Print("Would you like to play again? Enter yes or no:")
					fmt.Scan(&input)
					if input == "yes"{
						currentGold := hero.Gold
						hero = dependencies.GenerateCurrentHero(currentGold)
						break
					}else if input == "no"{
						keepPlaying = false
						fmt.Println("Game over")
						fmt.Printf("Current level: %v\n", hero.Level)
						fmt.Printf("Goblins slain: %v\n", hero.GoblinsSlain)
						break
					}
				}
				
			}
		}
		// Hero reached the end of the round and can now buy potions
		if !heroIsDead{
			fmt.Printf("You made it through level %v. You have slain %v goblins\n", hero.Level, hero.GoblinsSlain)
			for {
				fmt.Print("Would you like to buy a health potion? Enter yes or no:")
				time.Sleep(time.Second*2)
				fmt.Scan(&input)
				if input == "yes"{
					err:=hero.PurchasePotion()
					if err != nil{
						fmt.Println("You have the maximum amount of potions or do not have enough gold. You cannot purchase any potions.")
						break
					}
				}else if input == "no"{
					break
				}
			}
		}
		heroIsDead = false
	}
}