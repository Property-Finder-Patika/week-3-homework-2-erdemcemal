// Copyright © 2018 Inanc Gumus
// Learn Go Programming Course
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/
//
// For more tutorials  : https://learngoprogramming.com
// In-person training  : https://www.linkedin.com/in/inancgumus/
// Follow me on twitter: https://twitter.com/inancgumus

// ---------------------------------------------------------
// EXERCISE: Refactor the game store to funcs - step #1
//
//  Remember the game store program from the structs exercises?
//  Now, it's time to refactor it to funcs.
//
//  Create games.go file
//
//     1- Move the structs there
//
//     2- Add a func that creates and returns a game.
//
//        Name  : newGame
//        Inputs: id, price, name and genre
//        Output: game
//
//     3- Add a func that adds a `game` to `[]game` and returns `[]game`:
//
//        Name  : addGame
//        Inputs: []game, game
//        Output: []game
//
//     4- Add a func that uses newGame and addGame funcs to load up the game
//        store:
//
//        Name  : load
//        Inputs: none
//        Output: []game
//
//     5- Add a func that indexes games by ID:
//
//        Name  : indexByID
//        Inputs: []game
//        Output: map[int]game
//
//     6- Add a func that prints the given game:
//
//        Name  : printGame
//        Inputs: game
//
//
//  Go back to main.go and change the existing code with
//  the new funcs that you've created. There are hints
//  inside the code.
//
// ---------------------------------------------------------

package main

import (
	"bufio"
	"fmt"
	"github.com/Property-Finder-Patika/week-3-homework-2-erdemcemal/inanc-gumus/functions/refactor-to-funcs-1/models"
	"os"
	"strconv"
	"strings"
)

func main() {

	games := load()
	byID := indexByID(games)

	fmt.Printf("Inanc's game store has %d games.\n", len(games))

	in := bufio.NewScanner(os.Stdin)

	for {
		fmt.Printf(`
  > list   : lists all the games
  > id N   : queries a game by id
  > quit   : quits
`)

		if !in.Scan() {
			break
		}

		fmt.Println()

		cmd := strings.Fields(in.Text())
		if len(cmd) == 0 {
			continue
		}

		switch cmd[0] {
		case "quit":
			fmt.Println("bye!")
			return

		case "list":
			for _, g := range games {
				printGame(g)
			}

		case "id":
			if len(cmd) != 2 {
				fmt.Println("wrong id")
				continue
			}

			id, err := strconv.Atoi(cmd[1])
			if err != nil {
				fmt.Println("wrong id")
				continue
			}

			g, ok := byID[id]
			if !ok {
				fmt.Println("sorry. I don't have the game")
				continue
			}

			printGame(g)
		}
	}
}

func newGame(id, price int, name, genre string) models.Game {
	return models.Game{
		Item:  models.Item{ID: id, Name: name, Price: price},
		Genre: genre,
	}
}

func addGame(games []models.Game, game models.Game) []models.Game {
	return append(games, game)
}

func indexByID(games []models.Game) map[int]models.Game {
	byID := make(map[int]models.Game)
	for _, g := range games {
		byID[g.ID] = g
	}
	return byID
}

func load() []models.Game {
	games := make([]models.Game, 0)
	addGame(games, newGame(1, 50, "god of war", "action adventure"))
	addGame(games, newGame(2, 40, "x-com 2", "strategy"))
	addGame(games, newGame(3, 20, "minecraft", "sandbox"))
	return games
}
func printGame(g models.Game) {
	fmt.Printf("#%d: %-15q %-20s $%d\n",
		g.ID, g.Name, "("+g.Genre+")", g.Price)
}
