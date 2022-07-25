package dependencies

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

type Hero struct {
	MaxHealth    int
	CurrentHealth int
	AttackPower  int
	DefensePower int
	Potions      [5]int
	Gold         int
	GoblinsSlain int
	Level int
}

/*
Generates a Hero struct
*/
func GenerateHero() Hero{
	rand.Seed(time.Now().UnixNano())
	health:= rand.Intn(30-20+1)+20
	power:= rand.Intn(3-1+1)+1
	defence:= rand.Intn(5-1+1)+1
	potions:= [5]int{2,2,2,2,2}
	newHero:= Hero{health,health,power,defence,potions,0,0,1}
	return newHero
}

/*
Generates a Hero struct with the gold from the previous level
*/
func GenerateCurrentHero(gold, goblinsSlain,level int) Hero{
	rand.Seed(time.Now().UnixNano())
	health:= rand.Intn(30-20+1)+20
	power:= rand.Intn(3-1+1)+1
	defence:= rand.Intn(5-1+1)+1
	potions:= [5]int{2,2,2,2,2}
	newHero:= Hero{health,health,power,defence,potions,gold,goblinsSlain,level}
	return newHero
}

/*
Displays the stats of the hero
*/
func (h *Hero)DisplayStats(){
	fmt.Printf("Health: %d/%d; Power: %d; Defense: 1-%d; Number of potions: %d\n", h.CurrentHealth,h.MaxHealth,h.AttackPower,h.DefensePower,h.NumberOfPotions())
}

// Determines number of potions
func (h *Hero) NumberOfPotions() int{
	count:=0
	for _,val:=range h.Potions{
		if val == 2{
			count++
		}
	}
	return count
}

func (h *Hero) Attack(goblin *Goblin){
	rand.Seed(time.Now().UnixNano())
	goblinDefence:= rand.Intn(goblin.DefensePower+1-1)+1
	goblin.Health -= h.AttackPower - goblinDefence
}

func (h *Hero) UsePotion() error{
	for i:=0;i < len(h.Potions); i++{
		if h.Potions[i] == 2{
			h.CurrentHealth += h.Potions[i]
			h.Potions[i] = 0
			return nil
		}
	}
	return errors.New("You do not have any potions. You must attack.")
}

func (h *Hero) checkPotionCount() error{
	for i:=0; i < len(h.Potions); i++{
		if h.Potions[i] == 0{
			return nil
		}
	}
	return errors.New("max potions")
}

func (h *Hero) PurchasePotion() error{
	err:= h.checkPotionCount()
	if err != nil{
		return errors.New("You have the maximun amount of potions")
	}
	for i:=0; i < len(h.Potions); i++{
		if h.Potions[i] == 0 && h.Gold >= 2{
			h.Potions[i] = 2
			h.Gold -= 2
			return nil
		}
		if h.Gold < 2{
			return errors.New("You do not have enough gold to purchase a health potion.")
		}
	}
	return nil
}