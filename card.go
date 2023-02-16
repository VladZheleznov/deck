package deck

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Card struct {
	Name  string
	Suit  string
	Value int
}

func (c Card) String() string {
	return fmt.Sprintf("Card: %s of %s \n", c.Name, c.Suit)
}

var names = []string{"A", "2", "3", "4", "5", "6", "7", "8", "9", "10", "J", "Q", "K"}
var suits = []string{"Spades", "Diamonds", "Clubs", "Hearts"}

func New() []Card {
	var cardDeck []Card
	value := 1
	for _, suit := range suits {
		for _, name := range names {
			cardDeck = append(cardDeck, Card{Name: name, Suit: suit, Value: value})
			value++
		}
	}
	return cardDeck
}

func ShuffleDeck(с []Card) []Card {
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(с), func(i, j int) { с[i], с[j] = с[j], с[i] })
	return с
}

func SortDeck(c []Card) []Card {
	sort.Slice(c, func(i, j int) bool {
		return c[i].Value < c[j].Value
	})
	return c
}

func AddJoker(c []Card, count int) []Card {
	for i := 0; i < count; i++ {
		c = append(c, Card{Name: "Joker", Value: 53 + i})
	}
	return c
}

func AddDeck(c []Card, count int) []Card {
	for i := 0; i < count; i++ {
		c = append(c, SortDeck(c)...)
	}
	return c
}
