package main

func main() {
	// card := newCard()
	// fmt.Println("Card: ", card)
	cards := deck{newCard(), "King of Jacks"}
	cards = append(cards, "Seven of Hearts", "Queen of Spades")

	// loop
	cards.print()
}

func newCard() string {
	return "Five of Diamonds"
}
