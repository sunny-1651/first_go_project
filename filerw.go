package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func (d deck) dump_deck_to_file(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.to_string(",")), 0666)
}

func load_deck_from_file(filename string) deck {
	byteSlice, err := ioutil.ReadFile(filename)
	var saved_hand deck

	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(10)
	} else {
		saved_hand = strings.Split(string(byteSlice), ",")
	}
	return saved_hand
}
