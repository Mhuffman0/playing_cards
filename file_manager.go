package main

import (
	"log"
	"os"
	"strings"
)

func (d deck) to_string() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) to_bytes() []byte {
	return []byte(d.to_string())
}



func (d deck) saveToFile(filename string) error {
	current_dir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}


	return os.WriteFile(
		current_dir+"\\"+filename, // name of file
		d.to_bytes(),              // bytes to be written
		0666,                      // default file permissions world readable
	)
}

func readFromFile(filename string) (deck) {
	bs, err := os.ReadFile(filename)
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}

	deck:= strings.Split(string(bs), ", ")
	return deck
}