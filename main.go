package main

func main() {

	//go run main.go deck.go
	// var card string = "Ace of spades"
	// card := setCardName()
	// cards := []string{"ACE OF JOKER", setCardName()}
	// cards = append(cards, "ANOTHER JOKER")
	// for i, card := range cards {
	// 	fmt.Println(i+1, card)
	// }

	//connect to another package
	//go run main.go deck.go
	// cards := deck{"ACE OF JOKER", setCardName()}
	// cards = append(cards, "ANOTHER JOKER")
	// cards.printAllCards()

	//loop for in loop for
	cards := newDeck()
	cards.printAllCards()
}

// func setCardName() string {
// 	return "ACE OF DIAMOND"
// }
