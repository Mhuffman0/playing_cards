package main

import (
	"log"
	"os"
	"strings"
)

func (d deck) to_bytes() []byte {
	s := []string(d)
	return []byte(
		strings.Join(
			s,
			", ",
		))
}

func prependDirectory(filename string) string {
	current_dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	return current_dir + "\\" + filename
}

func (d deck) saveToFile(filename string) error {
	return os.WriteFile(
		prependDirectory(filename), // name of file
		d.to_bytes(),               // bytes to be written
		0666,                       // default file permissions world readable
	)
}

func readDeckFromFile(filename string) deck {
	bs, err := os.ReadFile(prependDirectory(filename))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	s := strings.Split(
		string(bs),
		", ",
	)
	return deck(s)
}
