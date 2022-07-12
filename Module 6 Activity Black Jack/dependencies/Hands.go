package dependencies

type Hand struct {
	Cards      []Card
	PlayerType string // Dealer or Player
	Busted     bool
}

func (h *Hand) IsBust() {
	var total int
	for _, val := range h.Cards {
		total += val.Value
	}
	if total > 21 {
		h.Busted = true
	}
}