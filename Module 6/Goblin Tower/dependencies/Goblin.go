package dependencies

import (
	"fmt"
	"math/rand"
	"time"
)

type Goblin struct {
	Health       int
	AttackPower  int
	DefensePower int
}


// Creates new goblin
func GenerateGoblin() Goblin {
	rand.Seed(time.Now().UnixNano())
	health := rand.Intn(10-5+1) + 5
	power := rand.Intn(3-2+1) + 2
	defence := rand.Intn(2-1+1) + 1
	newGoblin := Goblin{health, power, defence}
	return newGoblin
}

// Displays the goblins stats
func (g *Goblin) DisplayStats(){
	fmt.Printf("Health: %d; Power: %d; Defence: 1-%d\n", g.Health, g.AttackPower, g.DefensePower)
}

// Attack the hero and decrement hero's health
func (g *Goblin) Attack(hero *Hero) bool{
	rand.Seed(time.Now().UnixNano())
	heroDefence:= rand.Intn(hero.DefensePower+1-1)+1
	hitDamage:=g.AttackPower - heroDefence
	if hitDamage > 0{
		hero.CurrentHealth -= hitDamage
		return true
	}
	return false
}
