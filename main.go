package main

// import "fmt"
func main() {
	cards := buildDeck()
	cards.print()
	hand, cards := deal(cards, 5)
	cards.saveToFile("cards.txt")
	hand.saveToFile("hand.txt")
	cards2 := readDeckFromFile("cards.txt")
	cards2.print()
}
