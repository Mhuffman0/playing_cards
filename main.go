package main

func main() {
	cards := buildDeck()
	// cards.print()
	hand, cards := deal(cards, 5)
	// hand.print()
	// cards.print()
	// cards.to_string()
	// hand.to_string()
	hand.saveToFile("C:/Users/Mhuff/OneDrive/Documents/go_workspace/playing_cards/hand.txt")
}
