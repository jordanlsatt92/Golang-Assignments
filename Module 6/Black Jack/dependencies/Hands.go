package dependencies

import "fmt"

type Hand struct {
	Cards      []Card
	PlayerType string // Dealer or Player
	Busted     bool
	HasAce     bool
}

func (h *Hand) IsBust() {
	var total int
	for _, card := range h.Cards {
		if card.Value == 1 {
			h.HasAce = true
			total += 10
		}
		total += card.Value
	}
	if h.HasAce && total > 21 {
		total -= 10
	}
	if total > 21 {
		h.Busted = true
	}
}

func DisplayHands(hands []Hand, currentPlayer int) {
	var totalCards int = 0
	for i:= 0; i < len(hands); i++{
		if i != currentPlayer{
			if hands[i].PlayerType == "dealer"{
				fmt.Print("Dealer's hand: ")
			}else{
				fmt.Print("Player ",i,"'s hand: ")
			}
			for pos, card := range hands[i].Cards {
				if pos == 0 {
					fmt.Print(card.Value," ",card.Suit)
				} else {
					totalCards++
				}
			}
			if (totalCards > 1){
				fmt.Print(" + ", totalCards, " Cards\n")
			}else{
				fmt.Print(" + ", totalCards, " Card\n")
			}
		}
		totalCards = 0
	}
}

func (h *Hand) DisplayCurrentPlayerHand(){
	for _, card := range h.Cards{
		fmt.Print(card.Value," ",card.Suit, "  ")
	}
	fmt.Println()
}