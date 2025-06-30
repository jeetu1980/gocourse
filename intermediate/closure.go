package intermediate

import "fmt"

func main() {

	// sequence := adder()
	// fmt.Println(sequence())
	// fmt.Println(sequence())
	// fmt.Println(sequence())

	// sequence2 := adder()
	// fmt.Println(sequence2())
	// fmt.Println(sequence2())
	// fmt.Println(sequence())

	subtracter := func() func(int) int {
		countdown := 99
		return func(x int) int {
			countdown -= x
			return countdown
		}
	}()

	fmt.Println(subtracter(1))
	fmt.Println(subtracter(2))
	fmt.Println(subtracter(3))

}

func adder() func() int {
	i := 0
	fmt.Println("Previous value of i: ", i)
	return func() int {
		i++
		fmt.Println("added 1 to i: ", i)
		return i
	}

}
