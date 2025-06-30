package intermediate

import "fmt"

func main() {

	fmt.Println(factorial(5))
	fmt.Println(factorial(10))

	fmt.Println(sumofdigits(9))
	fmt.Println(sumofdigits(123))
	fmt.Println(sumofdigits(12345))

	fmt.Println(fibonacci(3))

}

func factorial(n int) int {

	if n == 0 {
		return 1
	}

	return n * factorial(n-1)

}

func sumofdigits(n int) int {
	if n < 10 {
		return n
	}

	return n%10 + sumofdigits(n/10)
}

func fibonacci(n int) int {
	fmt.Println("value of n: ", n)
	if n <= 1 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)

}
