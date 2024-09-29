package main

import "testing"

func TestNewDeck(t *testing.T) {
	d := new_deck()
	if len(d) != 52 {
		t.Errorf("Expected deck length of 52 but got %v", len(d))
	}
}

func TestDealfromDeck(t *testing.T) {
	d := new_deck()
	testcases := [4]int{5, 7, 2, 9}
	for _, num := range testcases {
		_, specimen := d.deal(num)
		if len(specimen) != num {
			t.Errorf("Deal function does not hand out specified number of cards, expected %v got %v", num, len(specimen))
		}
	}
}
