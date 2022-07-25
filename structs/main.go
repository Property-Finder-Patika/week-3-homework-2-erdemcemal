package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

const LogFile = "log.txt"

func init() {
	logFile, err := os.OpenFile(LogFile, os.O_APPEND|os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Panic(err)
	}
	log.SetOutput(logFile)
	log.SetFlags(log.Lshortfile | log.LstdFlags)
}

// Rectangle is a struct that has width and height.
type Rectangle struct {
	width  float64
	height float64
}

// Area calculate area of rectangle.
func (r *Rectangle) Area() float64 {
	return r.width * r.height
}

// Perimeter calculate the circumference of rectangle.
func (r *Rectangle) Perimeter() float64 {
	return 2 * (r.width + r.height)
}

// NewRectangle create a function that create an instance of rectangle struct and return it.
func NewRectangle(width float64, height float64) (*Rectangle, error) {
	// check if width and height are positive.
	if width < 0 || height < 0 {
		return nil, errors.New("width and height must be positive")
	}
	return &Rectangle{width, height}, nil
}

func main() {
	// create a new rectangle.
	rect, err := NewRectangle(10, 20)
	if err != nil {
		log.Println(err)
	}
	// print the area of rectangle.
	fmt.Println("area:", rect.Area())
	// print the perimeter of rectangle.
	fmt.Println("perimeter:", rect.Perimeter())
}
