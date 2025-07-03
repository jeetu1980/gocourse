package intermediate

import "fmt"

type rectangle struct {
	length float64
	width  float64
}

func (r rectangle) Area() float64 {
	return r.length * r.width

}
func (r *rectangle) Scale(factor float64) {
	r.length *= factor
	r.width *= factor
}

type MyInt int

func (i MyInt) IsPositive() bool {
	return i > 0
}

func (MyInt) WelcomeMessage() string {
	return "Welcome to the Jungle"

}

type Shape struct {
	rectangle
}

func main() {

	rect := rectangle{
		length: 10,
		width:  5,
	}

	area := rect.Area()
	fmt.Println("Area of the rectangle before factoring:", area)

	rect.Scale(2)
	area = rect.Area()
	fmt.Println("Area of the rectangle after factoring:", area)

	int1 := MyInt(5)
	int2 := MyInt(-3)

	fmt.Println(int1.IsPositive())
	fmt.Println(int2.IsPositive())

	s := Shape{rectangle: rectangle{length: 20, width: 20}}

	fmt.Println("Area of the Shap is", s.Area())

}
