package structs_example

import (
	"fmt"
	"testing"
)

func TestPerimeter(t *testing.T) {
	r := Rectangle{10.0, 10.0}
	got := r.Perimeter()
	want := 40.0

	if got != want {
		t.Errorf("got %.2f want perim %.2f", got, want)
	}
}

type Shape interface {
	Area() float64
}

func TestArea(t *testing.T) {

	checkArea := func(t testing.TB, shape Shape, hasArea float64) {
		t.Helper()
		got := shape.Area()
		if got != hasArea {
			t.Errorf("%#v got %.2f hasArea %.2f", shape, got, hasArea)
		}
	}

	tests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "rectangle", shape: Rectangle{Width: 12.0, Height: 6.0}, hasArea: 72.0},
		{name: "circle", shape: Circle{Radius: 10.0}, hasArea: 314.1592653589793},
		{name: "triangle", shape: Triangle{sideA: 12.0, sideB: 6.0}, hasArea: 36.0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			checkArea(t, test.shape, test.hasArea)
		})
	}

}

func TestMutation(t *testing.T) {
	r := Rectangle{1.0, 3.0}
	r.MutA()
	fmt.Println("width: ", r.Width, "height: ", r.Height)
	r.MutB()
	fmt.Println("width: ", r.Width, "height: ", r.Height)
}
