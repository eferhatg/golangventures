package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"strings"
	"time"
)

type deck []string

func newDeck() deck {
	d := deck{}

	suits := []string{"Spades", "Hearts", "Diamonds", "Clubs"}
	values := []string{"Ace", "Two", "Four", "Three", "Five", "Six", "Seven", "Eight", "Nine", "Ten", "Jack", "Queen", "King"}

	for _, suit := range suits {
		for _, val := range values {
			d = append(d, val+" of "+suit)
		}
	}
	return d
}

func (d deck) print() {
	for i, card := range d {
		fmt.Println(i, card)
	}
}

func deal(d deck, handsize int) (deck, deck) {
	return d[:handsize], d[handsize:]
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(strings.Join([]string(d), ",")), 0666)
}

func newDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile("initial_cards")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)

	}
	cardArray := strings.Split(string(bs), ",")
	return deck(cardArray)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		newPos := r.Intn(len(d) - 1)
		d[i], d[newPos] = d[newPos], d[i]
	}

}
