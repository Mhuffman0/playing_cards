package main

import "testing"

func testBuildDeck(t *testing.T) {
	cards := buildDeck()
	if len(cards) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(cards))
	}
}

// 	suits := map[string]string{
// 		"diamonds": "red",
// 		"hearts":   "red",
// 		"clubs":    "black",
// 		"spades":   "black",
// 	}

// 	faces := map[string]int{
// 		"ace":   1,
// 		"two":   2,
// 		"three": 3,
// 		"four":  4,
// 		"five":  5,
// 		"six":   6,
// 		"seven": 7,
// 		"eight": 8,
// 		"nine":  9,
// 		"ten":   10,
// 		"jack":  10,
// 		"queen": 10,
// 		"king":  10,
// 	}

// 	for suit := range suits {
// 		for face := range faces {
// 			cards = append(cards, fmt.Sprintf("%s of %s", face, suit))
// 		}

// 	}
// 	return cards
// }

// func (d deck) Testprint() {

// 	for i, card := range d {
// 		fmt.Println(i, card)
// 	}
// }

// func Testdeal(d deck, handsize int) (deck, deck) {
// 	return d[:handsize], d[handsize:]
// }

// func (d deck) Testshuffle() {
// 	source := rand.NewSource(time.Now().UnixNano())
// 	r := rand.New(source)
// 	for i := range d {
// 		newPosition := r.Intn(len(d) - 1)
// 		d[i], d[newPosition] = d[newPosition], d[i]
// 	}
// }
