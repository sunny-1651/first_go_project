package main

import (
	"os"
	"testing"
)

func TestFileRWops(t *testing.T) {
	os.Remove("_deck_file_rw_testing")
	d := new_deck()
	cards := d.dump_deck_to_file("_deck_file_rw_testing")

	if cards != nil {
		t.Errorf("Failed while writing deck to file with error %s", cards)
	}

	retrieved_deck := load_deck_from_file("_deck_file_rw_testing")

	if retrieved_deck.to_string(",") != d.to_string(",") {
		t.Errorf("the dumped deck \n{%s} \ndoes not match loaded deck \n{%s}", retrieved_deck.to_string(", "), d.to_string(", "))
	}

}
