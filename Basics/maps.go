package basics

import "fmt"

func main() {

	mymap := make(map[string]int)

	fmt.Println("Map created using make: ", mymap)

	mymap2 := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	fmt.Println("Map2: ", mymap2)

	var nilmap map[string]int

	fmt.Println("Nil map: ", nilmap)

	_, ok := mymap2["a"]

	if ok {
		fmt.Println("Value exits in map")
	} else {
		fmt.Println("Value does not exits in map")

	}

}
