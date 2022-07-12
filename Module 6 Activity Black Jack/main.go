package main

import (
	"BlackJackModule/dependencies"
	"fmt"
)


func main(){
	
	//Holds the number of players
	var numPlayers int
	// Holds the number of rounds the user wants to play
	var roundsToPlay int
	fmt.Print("Enter the number of players (limit: 4): ")
	fmt.Scan(&numPlayers)
	if numPlayers > 4 {numPlayers = 4}
	fmt.Println("Enter the number of rounds you want to play (limit:5): ")
	fmt.Scan(&roundsToPlay)
	if roundsToPlay > 5 {roundsToPlay = 5}
	
	// Array to hold the hands of the players and dealer
	hands := make([]dependencies.Hand,1)
	hands[0] = dependencies.Hand{make([]dependencies.Card,0),"dealer", false}
	for i:=0; i < numPlayers; i++{
		hands = append(hands, dependencies.Hand{make([]dependencies.Card,0),"player", false})
	}



	for i:= 0; i < roundsToPlay; i++{
		deck := dependencies.Deck{}
		deck.GenerateDeck()
		deck.ShuffleDeck()

		// Deal Cards
		deck.DealCards(hands)
		
	}
}