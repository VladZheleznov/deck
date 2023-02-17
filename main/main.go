package main

import (
	"fmt"

	"github.com/VladZheleznov/deck"
)

func main() {
	res := deck.New()
	fmt.Println("Колода из 52 карт: ")
	fmt.Println(deck.New())
	fmt.Println("=============================")
	fmt.Println("=============================")
	fmt.Println("Перетасованная колода карт: ")
	fmt.Println(deck.ShuffleDeck(res))
	fmt.Println("=============================")
	fmt.Println("=============================")
	fmt.Println("Отсортированная колода: ")
	fmt.Println(deck.SortDeck(res))
	fmt.Println("=============================")
	fmt.Println("=============================")
	fmt.Println("Колода с n джокерами")
	fmt.Println(deck.AddJoker(res, 1))
	fmt.Println("=============================")
	fmt.Println("=============================")
	fmt.Println("N+1 колода")
	fmt.Println(deck.AddDeck(res, 4))

}
