package dependencies

import (
	"fmt"
	"math/rand"
	"time"
)

// Deck Struct that holds the 52 cards.
type Deck struct {
	Cards []Card
}

//Creates a deck of cards
func (d *Deck) GenerateDeck() {
	d.Cards = make([]Card, 0)
	suits := []string{"spade", "heart", "diamond", "club"}
	for i := 0; i < 4; i++ {
		for j := 1; j <= 13; j++ {
			d.Cards = append(d.Cards, Card{j, suits[i]})
		}
	}
}

// Shuffles the deck of cards using random numbers
func (d *Deck) ShuffleDeck(){
	deck:= Deck{}
	deck.GenerateDeck()
	for i:=0; i < 52; i++{
		if d.Cards[i] != deck.Cards[i]{
			fmt.Println("Missing cards in the deck; please try again")
			return
		}
	}
	for i:=0; i < 500; i++{
		rand.Seed(time.Now().UnixNano()+7)
		firstRand := rand.Intn(52)
		rand.Seed(time.Now().UnixNano()+19)
		secondRand := rand.Intn(52)
		d.Cards[firstRand], d.Cards[secondRand] = d.Cards[secondRand], d.Cards[firstRand]
	}
}

// Returns the top card of the deck, and removes that card from the deck
func (d *Deck) Hit() Card{
	temp:= d.Cards[0]
	d.Cards = d.Cards[1:]
	return temp
}

func (d *Deck) DealCards(hands []Hand){
	for i:=0; i < 2; i++{
		for j:=0; j < len(hands);j++{
			hands[j].Cards = append(hands[j].Cards, d.Hit())
		}
	}
}