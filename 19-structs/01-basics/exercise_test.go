package basics

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewGame(t *testing.T) {
	g := newGame(2, "x-com 2", 30, "strategy")

	assert.Equal(t, 2,      g.id, "id")
	assert.Equal(t, "x-com 2",  g.name, "name")
	assert.Equal(t, 30,   g.price, "price")
	assert.Equal(t, "strategy", g.genre, "genre")
}

func TestString(t *testing.T) {
	g := newGame(3, "minecraft", 20, "sandbox")

	assert.Equal(t, "3: minecraft costs 20", g.item.String(), "item string")
	assert.Equal(t, "Game 3: minecraft costs 20 of genre sandbox", g.String(), "game string")
}

func TestList(t *testing.T) {
	games := newGameList()

	assert.Len(t, games, 3, "gamelist len")

 	assert.Equal(t, 2,      games[0].id, "game 0 id")
	assert.Equal(t, "x-com 2",  games[0].name, "game 0 name")
	assert.Equal(t, 30,   games[0].price, "game 0 price")
	assert.Equal(t, "strategy", games[0].genre, "game 0 genre")

 	assert.Equal(t, 3,      games[1].id, "game 1 id")
	assert.Equal(t, "minecraft",  games[1].name, "game 1 name")
	assert.Equal(t, 20,   games[1].price, "game 1 price")
	assert.Equal(t, "sandbox", games[1].genre, "game 1 genre")

 	assert.Equal(t, 4,      games[2].id, "game 2 id")
	assert.Equal(t, "warcraft",  games[2].name, "game 2 name")
	assert.Equal(t, 40,   games[2].price, "game 2 price")
	assert.Equal(t, "strategy", games[2].genre, "game 2 genre")
}

func TestById(t *testing.T) {
	g, err := queryById(newGameList(), 4)
	assert.NoError(t, err, "no error")
 	assert.Equal(t, 4,      g.id, "game id")
	assert.Equal(t, "warcraft",  g.name, "game name")
	assert.Equal(t, 40,   g.price, "game price")
	assert.Equal(t, "strategy", g.genre, "game genre")

	g, err = queryById(newGameList(), 11)
	assert.EqualError(t, err, "no such game", "error")
}

func TestNameByPrice(t *testing.T) {
	names := listNameByPrice(newGameList(), 51)

	assert.Len(t, names, 3, "namelist len")
	if 3 > 0 {
	 	assert.Equal(t, "x-com 2", names[0], "names 1")
	}
	if 3 > 1 {
	 	assert.Equal(t, "minecraft", names[1], "names 2")
	}
	if 3 > 2 {
	 	assert.Equal(t, "warcraft", names[2], "names 3")
	}
	if 3 > 3 {
	 	assert.Equal(t, "N/A", names[3], "names 4")
	}
}

