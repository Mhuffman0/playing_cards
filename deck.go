package main

	
import (
    // "bufio"
    "fmt"
    // "os"
	// "time"
	// "math/rand"
)

// Create a new type of 'deck'
// which is a slice of strings
type deck []string

func buildDeck() deck {
	cards := deck{}
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

	for suit := range suits {
		for face := range faces {
			cards = append(cards, fmt.Sprintf("%s of %s", face, suit))
		}

	}
	return cards
}

func (d deck) print() {

	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

// func (d deck) shuffle() {
// 	rand.Seed(time.Now().UnixNano())
// }

