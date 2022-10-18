package main

import (
	// "log"
	"os"
	"testing"
)

func TestSaveToDeck(t *testing.T) {
	test_file := "_decktesting"

	os.Remove(test_file)
	deck := buildDeck()
	deck.saveToFile(test_file)
	loadedDeck := readDeckFromFile(test_file)
	if len(loadedDeck) != 52 {
		t.Errorf("Expected deck length of 52, but got %d", len(loadedDeck))
	}
	os.Remove(test_file)
}

// func (d deck) to_bytes() []byte {
// 	s := []string(d)
// 	return []byte(
// 		strings.Join(
// 			s,
// 			", ",
// 		))
// }

// func prependDirectory(filename string) string {
// 	current_dir, err := os.Getwd()
// 	if err != nil {
// 		log.Println(err)
// 		os.Exit(1)
// 	}
// 	return current_dir + "\\" + filename
// }

// func (d deck) saveToFile(filename string) error {
// 	return os.WriteFile(
// 		prependDirectory(filename), // name of file
// 		d.to_bytes(),               // bytes to be written
// 		0666,                       // default file permissions world readable
// 	)
// }

// func readDeckFromFile(filename string) deck {
// 	bs, err := os.ReadFile(prependDirectory(filename))
// 	if err != nil {
// 		log.Println(err)
// 		os.Exit(1)
// 	}

// 	s := strings.Split(
// 		string(bs),
// 		", ",
// 	)
// 	return deck(s)
// }
