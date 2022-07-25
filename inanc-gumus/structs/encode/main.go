package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ---------------------------------------------------------
// EXERCISE: Encode
//
//  Add a new command: "save". Encode the games to json, and
//  print it, then terminate the loop.
//
//  1. Create a new struct type with exported fields: ID, Name, Genre and Price.
//
//  2. Create a new slice using the new struct type.
//
//  3. Save the games into the new slice.
//
//  4. Encode the new slice.
//
//
// RESTRICTION
//  Do not export the fields of the game struct.
//
//
// EXPECTED OUTPUT
//  Inanc's game store has 3 games.
//
//    > list   : lists all the games
//    > id N   : queries a game by id
//    > save   : exports the data to json and quits
//    > quit   : quits
//
//  save
//
//  [
//          {
//                  "id": 1,
//                  "name": "god of war",
//                  "genre": "action adventure",
//                  "price": 50
//          },
//          {
//                  "id": 2,
//                  "name": "x-com 2",
//                  "genre": "strategy",
//                  "price": 40
//          },
//          {
//                  "id": 3,
//                  "name": "minecraft",
//                  "genre": "sandbox",
//                  "price": 20
//          }
//  ]
//
// ---------------------------------------------------------

type game struct {
	ID    int
	Name  string
	Genre string
	Price int
}

func main() {
	// use your solution from the previous exercise

	games := []game{
		{ID: 1, Name: "god of war", Genre: "action adventure", Price: 50},
		{ID: 2, Name: "x-com 2", Genre: "strategy", Price: 40},
		{ID: 3, Name: "minecraft", Genre: "sandbox", Price: 20},
	}

	for {
		fmt.Printf(`
  > list   : lists all the games
  > id N   : queries a game by id
  > save   : exports the data to json and quits
  > quit   : quits
`) // index the games by id
		byID := make(map[int]game)
		for _, g := range games {
			byID[g.ID] = g
		}

		fmt.Printf("Inanc's game store has %d games.\n", len(games))

		in := bufio.NewScanner(os.Stdin)
		for {
			fmt.Printf(`
  > list   : lists all the games
  > id N   : queries a game by id
  > save   : exports the data to json and quits
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
					fmt.Printf("#%d: %-15q %-20s $%d\n",
						g.ID, g.Name, "("+g.Genre+")", g.Price)
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

				fmt.Printf("#%d: %-15q %-20s $%d\n",
					g.ID, g.Name, "("+g.Genre+")", g.Price)

			case "save":
				type jsonGame struct {
					ID    int    `json:"id"`
					Name  string `json:"name"`
					Genre string `json:"genre"`
					Price int    `json:"price"`
				}

				// load the data into the encodable game values
				var encodable []jsonGame
				for _, g := range games {
					encodable = append(encodable,
						jsonGame{g.ID, g.Name, g.Genre, g.Price})
				}

				out, err := json.MarshalIndent(encodable, "", "\t")
				if err != nil {
					fmt.Println("Sorry:", err)
					continue
				}

				fmt.Println(string(out))
				return
			}
		}
	}
}
