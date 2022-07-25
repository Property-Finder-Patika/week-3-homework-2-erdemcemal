package main

import (
	"encoding/json"
	"fmt"
	models "github.com/Property-Finder-Patika/week-3-homework-2-erdemcemal/inanc-gumus/structs"
)

// ---------------------------------------------------------
// EXERCISE: Decode
//
//  At the beginning of the file:
//
//  1. Load the initial data to the game store from json.
//     (see the data constant below)
//
//  2. Load the decoded values into the usual `game` values (to the games slice as well).
//
//     So the rest of the program can work intact.
//
//
// HINT
//
//  Move the jsonGame type to the top and reuse it both when
//  loading the initial data, and in the "save" command.
//
//
// EXPECTED OUTPUT
//  Please run the solution to see the output.
// ---------------------------------------------------------

const data = `
[
        {
                "id": 1,
                "name": "god of war",
                "genre": "action adventure",
                "price": 50
        },
        {
                "id": 2,
                "name": "x-com 2",
                "genre": "strategy",
                "price": 40
        },
        {
                "id": 3,
                "name": "minecraft",
                "genre": "sandbox",
                "price": 20
        }
]`

func main() {
	// use your solution from the previous exercise
	var games []models.Game
	err := json.Unmarshal([]byte(data), &games)
	if err != nil {
		fmt.Println(err)
	}
	for _, game := range games {
		println(game.ID, game.Name, game.Genre, game.Price)
	}
}
