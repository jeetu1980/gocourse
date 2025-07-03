package intermediate

import "fmt"

func swap[T any](a, b T) (T, T) {
	return b, a
}

type Elements[T any] struct {
	elements []T
}

func (e *Elements[T]) push(element T) {
	e.elements = append(e.elements, element)
}

func (e *Elements[T]) pop() (T, bool) {
	if len(e.elements) == 0 {
		var zero T
		return zero, false
	} else {
		element := e.elements[len(e.elements)-1]
		e.elements = e.elements[:len(e.elements)-1]
		return element, true
	}

}

func (e Elements[T]) isempty() bool {
	return len(e.elements) == 0
}

func (e Elements[T]) printAll() {
	if len(e.elements) == 0 {
		fmt.Print("There are elements")
	} else {
		for _, elements := range e.elements {

			fmt.Print(elements)
		}
	}
	fmt.Print("\n")
}

func main() {

	// x, y := 1, 2
	// fmt.Println("x: ", x, "y: ", y)

	// x, y = swap(x, y)
	// fmt.Println("x: ", x, "y: ", y)

	// x1, y1 := "Amit", "kumar"
	// fmt.Println("x1: ", x1, "y1: ", y1)
	// x1, y1 = swap(x1, y1)
	// fmt.Println("x1: ", x1, "y1: ", y1)

	stack := Elements[int]{}
	stack.push(1)
	stack.push(2)
	stack.push(3)
	stack.push(4)
	stack.printAll()
	fmt.Println(stack.pop())
	fmt.Println("is stack empty", stack.isempty())
	stack.printAll()
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println(stack.pop())
	fmt.Println("is stack empty", stack.isempty())
}
