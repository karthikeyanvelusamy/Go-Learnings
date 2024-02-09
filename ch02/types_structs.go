package ch02

import "fmt"

type Rectangle struct {
	Length float64
	Width  float64
}

func (r Rectangle) Area() float64 {
	return r.Length * r.Width
}

func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Length + r.Width)
}

func Struct_ex() {
	r := Rectangle{Length: 2, Width: 5}

	fmt.Println("Area : ", r.Area(), "Perimeter: ", r.Perimeter())
}
