package main

import "fmt"

type deck []string

func newDeck() deck {
	cards := deck{}

	cardSuits := []string{"\u2660", "\u2666", "\u2665", "\u2663"} // Spades, Diamonds, Hearts, Clubs
	cardValues := []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}

	for _, suite := range cardSuits {
		for _, value := range cardValues {
			// fmt.Println(i, j)
			cards = append(cards, value+suite)
		}
	}

	return cards
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, ")", card)
	}
}

func deal(d deck, handSize int) (deck, deck) {
	return d[:handSize], d[handSize:]
}
