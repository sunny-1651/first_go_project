package main

// import (
// 	"fmt"
// )

func main() {
	new_hand := new_deck()
	_, b := new_hand.deal(13)
	b.string_to_file("./dealt_hand.txt")
	// a.show()
	// test := load_deck_from_file("dealt_hand.txt")
	b.show()
	b.shuffle()
	b.show()
}
