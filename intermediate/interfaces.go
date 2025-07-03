package intermediate

import (
	"fmt"
	"math"
)

type geomatry interface {
	area() float64
	perimeter() float64
}

type rect struct {
	length float64
	width  float64
}

type circle struct {
	radius float64
}

func (r rect) area() float64 {
	return r.length * r.width
}

func (r rect) perimeter() float64 {
	return 2 * (r.length + r.width)
}

func (c circle) area() float64 {
	return math.Pi * c.radius * c.radius
}

func (c circle) perimeter() float64 {
	return 2 * math.Pi * c.radius
}

func (c circle) diameter() float64 {
	return 2 * c.radius
}

func measure(g geomatry) {
	fmt.Println(g)
	fmt.Println(g.area())
	fmt.Println(g.perimeter())

}

func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("Type of i is int")
	case string:
		fmt.Println("Type of i is string")
	default:
		fmt.Println("Type of i is unknown")

	}
}

func main() {

	r := rect{length: 10, width: 20}
	c := circle{radius: 5}

	measure(r)
	measure(c)
	fmt.Println(c.diameter())

	i := 5
	printType(i)

	s := "String"
	printType(s)

	b := true
	printType(b)

}
