package basics

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)

	target := random.Intn(100) + 1

	fmt.Println("Welcome to the Guessing game")
	fmt.Println("I have choosen a number between 1 and 100")
	fmt.Println("Can you guess what it is")

	var guess int

	for {
		fmt.Println("Please input your guess:")
		fmt.Scanln(&guess)

		if guess == target {
			fmt.Println("Congratulation, you guessed right number")
			break
		} else if guess > target {
			fmt.Println("Your number is too high, please guess again")
		} else {
			fmt.Println("Your number is too low, please guess again")
		}
	}
}
