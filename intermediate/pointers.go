package intermediate

import "fmt"

func main() {

	var ptr *int

	var a int = 10

	fmt.Println(a)

	ptr = &a

	fmt.Println(ptr)

	fmt.Println(&a)

	fmt.Println(*ptr)

	modifyvalue(ptr)

	fmt.Println(a)

}

func modifyvalue(ptr *int) {
	*ptr++
}
