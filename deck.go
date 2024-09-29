package main

import (
	"fmt"
	"math/rand"
	"strings"
)

type deck []string

func new_deck() deck {
	var fresh_deck deck

	suits := [4]string{"hearts", "diamonds", "spades", "clubs"}
	car_num := [13]string{"ace", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "jack", "queen", "king"}

	for _, suit := range suits {
		for _, num := range car_num {
			fresh_deck = append(fresh_deck, num+" of "+suit)
		}
	}
	return fresh_deck
}

func (d deck) show() {
	for i, card := range d {
		fmt.Println(i+1, " -> ", card)
	}
}

func (d deck) remove(index int) deck {
	if len(d)-index <= 1 {
		return d[:index]
	} else {
		return append(d[:index], d[index+1:]...)
	}
}

func (d deck) deal(size int) (deck, deck) {
	var dealt_hand deck
	for size > 0 {
		indx := rand.Intn(len(d))
		dealt_hand = append(dealt_hand, d[indx])
		d = d.remove(indx)
		size--
	}
	return d, dealt_hand
}

func deal(d deck, size int) (deck, deck) {
	var dealt_hand deck
	for size > 0 {
		indx := rand.Intn(len(d))
		dealt_hand = append(dealt_hand, d[indx])
		d = d.remove(indx)
		size--
	}
	return d, dealt_hand
}

func (d deck) to_string(seperator string) string {
	return strings.Join([]string(d), seperator)
}

func (d deck) shuffle() {
	for i, _ := range d {
		x := rand.Intn(len(d))
		d[i], d[x] = d[x], d[i]
	}
}
