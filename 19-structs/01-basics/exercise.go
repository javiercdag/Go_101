package basics

import (
	"errors"
	"fmt"
)

// DO NOT REMOVE THIS COMMENT
//go:generate go run ../../exercises-cli.go -student-id=$STUDENT_ID generate

type item struct {
	id    int
	name  string
	price int
}

type game struct {
	item
	genre string
}

// newGame returns a new game struct.
func newGame(id int, name string, price int, genre string) game {
	return game{genre: genre, item: item{id: id, name: name, price: price}}
}

// String stringifies an item.
func (i item) String() string {
	return fmt.Sprintf("%d: %s costs %d", i.id, i.name, i.price)
}

// String stringifies a game.
func (g game) String() string {
	return fmt.Sprintf("Game %s of genre %s", g.item.String(), g.genre)
}

// newGameList creates a game store.
func newGameList() []game {
	return []game{
		{
			item:  item{id: 2, name: "x-com 2", price: 30},
			genre: "strategy",
		},
		{
			item:  item{id: 3, name: "minecraft", price: 20},
			genre: "sandbox",
		},
		{
			item:  item{id: 4, name: "warcraft", price: 40},
			genre: "strategy",
		},
	}
}

// queryById returns the game in the specified store with the given id or returns a "no such game" error.
func queryById(games []game, id int) (game, error) {
	for _, g := range games {
		if g.item.id == id {
			return g, nil
		}
	}

	return game{}, errors.New("no such game")
}

// listNameByPrice returns the name of the game(s) with price equal or smaller than a given price.
func listNameByPrice(games []game, price int) []string {
	var names []string

	for _, g := range games {
		if g.item.price <= price {
			names = append(names, g.item.name)
		}
	}

	return names
}
