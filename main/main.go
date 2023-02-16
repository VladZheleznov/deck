package main

import (
	"deck"
	"fmt"
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
	fmt.Println("Колода с джокерами")
	fmt.Println(deck.AddJoker(res, 1))
	fmt.Println("=============================")
	fmt.Println("=============================")
	fmt.Println(deck.AddDeck(res, 3))
	fmt.Println("=============================")
	fmt.Println("=============================")

}
