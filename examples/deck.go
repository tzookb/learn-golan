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
	cards := deck{}
	cardSuits := []string{"Club", "Heart", "Diamond", "Spade"}
	cardNumbers := []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "10"}
	for _, suit := range cardSuits {
		for _, cardValue := range cardNumbers {
			cards = append(cards, cardValue+" of "+suit)
		}
	}
	return cards
}

func (d deck) print() {
	for _, card := range d {
		fmt.Println(card)
	}
}

func (d deck) deal(amount int) (deck, deck) {
	return d[:amount], d[amount:]
}

func (d deck) toString() string {
	return strings.Join([]string(d), ",")
}

func (d deck) saveToFile(filename string) error {
	return ioutil.WriteFile(filename, []byte(d.toString()), 0666)
}

func (d deck) shuffle() {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)
	for i := range d {
		switchIdx := r.Intn(len(d) - 1)
		d[i], d[switchIdx] = d[switchIdx], d[i]
	}
}

func getDeckFromFile(filename string) deck {
	bs, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println("Error: ", err)
		os.Exit(1)
	}
	return deck(strings.Split(string(bs), ","))
}
