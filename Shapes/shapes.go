package shapes

import "fmt"

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Perimeter() float64 {
	return 4 * s.Side
}

type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return 3.14 * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * 3.14 * c.Radius
}

type Triangle struct {
	A float64
	B float64
	C float64
}

func (t Triangle) Area() float64 {
	return (t.A * t.B) / 2
}

func (t Triangle) Perimeter() float64 {
	return t.A + t.B + t.C
}

func Demo() {
	shapes := []Shape{
		Rectangle{Width: 4, Height: 5},
		Square{Side: 3},
		Circle{Radius: 2},
		Triangle{A: 3, B: 4, C: 5},
	}

	for _, shape := range shapes {
		fmt.Println("Area:", shape.Area())
		fmt.Println("Perimeter:", shape.Perimeter())
	}
}
