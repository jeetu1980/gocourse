package intermediate

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	message := "Hello \nWorld!"
	message1 := "Hello \tWorld!"
	message2 := "Hello \rWorld!"
	message4 := `Hello \nWorld!`

	fmt.Println(message)
	fmt.Println(message1)
	fmt.Println(message2)
	fmt.Println(message4)

	fmt.Println("Rune count in a greetings:", utf8.RuneCountInString("greetings"))
}
