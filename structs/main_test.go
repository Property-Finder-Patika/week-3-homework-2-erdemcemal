package main

import "testing"

func TestRectangle_Area(t *testing.T) {
	rect, err := NewRectangle(10, 20)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if rect.Area() != 200 {
		t.Errorf("area is not 200")
	}
}

// TestRectangle_Perimeter tests the perimeter of rectangle.
func TestRectangle_Perimeter(t *testing.T) {
	rect := Rectangle{width: 10, height: 20}
	if rect.Perimeter() != 60 {
		t.Errorf("perimeter is not 60")
	}
}

// TestRectangle_NewRectangle tests the NewRectangle function.
func TestRectangle_NewRectangle(t *testing.T) {
	rect, err := NewRectangle(10, 20)
	if err != nil {
		t.Errorf("error: %v", err)
	}
	if rect.width != 10 || rect.height != 20 {
		t.Errorf("rectangle is not 10x20")
	}
}

// TestRectangle_NewRectangle_Error tests the NewRectangle function with negative values.
func TestRectangle_NewRectangle_Error(t *testing.T) {
	_, err := NewRectangle(-10, 20)
	if err == nil {
		t.Errorf("error is nil")
	}
	_, err = NewRectangle(10, -20)
	if err == nil {
		t.Errorf("error is nil")
	}
}

// TestRectangle_NewRectangle_Error tests the NewRectangle function with zero values.
func TestRectangle_NewRectangle_Error_Zero(t *testing.T) {
	_, err := NewRectangle(0, 20)
	if err == nil {
		t.Errorf("error is nil")
	}
	_, err = NewRectangle(10, 0)
	if err == nil {
		t.Errorf("error is nil")
	}
}
