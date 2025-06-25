package basics

import "fmt"

func main() {

	//Iterating through arrays

	// myint := [5]int{1, 2, 3, 4, 5}

	// for i := 0; i < len(myint); i++ {

	// 	fmt.Println(myint[i])
	// }

	// myint := [5]int{1, 2, 3, 4, 5}

	// for index, value := range myint {

	// 	fmt.Printf("Index: %d Value: %d \n", index, value)
	// }

	// If you only need the value (and not the index), use an underscore for the index:
	// myint := [5]int{1, 2, 3, 4, 5}

	// for _, value := range myint {

	// 	fmt.Printf("Value: %d \n", value)
	// }

	// If you only need the index (and not the value):
	// myint := [5]int{1, 2, 3, 4, 5}

	// for index := range myint {

	// 	fmt.Printf("Index: %d \n", index)
	// }

	var myarrays *[5]int

	original := [5]int{1, 2, 3, 4, 5}

	myarrays = &original

	myarrays[3] = 100

	fmt.Println("Original Array is ", original)
	fmt.Println("Referrenced Array is ", myarrays)

}
