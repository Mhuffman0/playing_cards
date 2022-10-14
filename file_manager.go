package main

import (
	"io/ioutil"
	"strings"
)

func (d deck) to_string() string {
	return strings.Join([]string(d), ", ")
}

func (d deck) to_bytes() []byte {
	return []byte(d.to_string())
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(
		filename,     // name of file
		d.to_bytes(), // bytes to be written
		0666,         // default file permissions world readable
	)
}
