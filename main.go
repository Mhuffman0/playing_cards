package main

import (
	"fmt"
)

var cards [52]string

func main() {
	BuildDeck()
	fmt.Printf("%+q\n", cards)
}

func BuildDeck() {
	suits := map[string]string{
		"diamonds": "red",
		"hearts":   "red",
		"clubs":    "black",
		"spades":   "black",
	}

	faces := map[string]int{
		"ace":   1,
		"two":   2,
		"three": 3,
		"four":  4,
		"five":  5,
		"six":   6,
		"seven": 7,
		"eight": 8,
		"nine":  9,
		"ten":   10,
		"jack":  10,
		"queen": 10,
		"king":  10,
	}
	i := 0
	for suit := range suits {
		for face := range faces {
			cards[i] = fmt.Sprintf("%s of %s", face, suit)
			i++
		}

	}
}
