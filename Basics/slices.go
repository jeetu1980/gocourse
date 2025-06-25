package basics

import "fmt"

func main() {

	arr := []int{1, 2, 3, 4, 5, 7}

	slice1 := arr[0:3]
	slice2 := arr[2:5]

	// slice3 := arr[:4] // from index 0 up to (but not including) index 4 -> {10, 20, 30, 40}
	// slice4 := arr[3:] // from index 3 to the end -> {40, 50, 60}
	// slice5 := arr[:]  // a slice containing all elements of the array -> {10, 20, 30, 40, 50, 60}

	s := make([]int, 5)
	r := make([]string, 0, 10)
	t := make([]bool, 3, 5)

	fmt.Println("slice1:", slice1, "len:", len(slice1), "cap:", cap(slice1))
	fmt.Println("slice2:", slice2, "len:", len(slice2), "cap:", cap(slice2))

	fmt.Println(s, "\n", r, "\n", t)

}
