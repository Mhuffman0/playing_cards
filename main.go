package main

// import "fmt"
func main() {
	// cards := buildDeck()
	// cards.print()
	// hand, cards := deal(cards, 5)
	// hand.print()
	// cards.print()
	// cards.to_string()
	// hand.to_string()
	// cards.saveToFile("cards.txt")
	// hand.saveToFile("hand.txt")
	cards := readFromFile("C:/Users/mhuff/OneDrive/Documents/go_workspace/playing_cards/cards.txt")
	cards.print()
}
