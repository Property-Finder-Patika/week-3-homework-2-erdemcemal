package models

type Item struct {
	ID    int
	Name  string
	Price int
}

type Game struct {
	Item
	Genre string
}
