package intermediate

import "fmt"

func main() {

	p := Person{
		firstname: "Amit",
		lastName:  "Kumar",
		age:       45,
		address: Address{
			city:    "Delhi",
			country: "India",
		},
		Phonebook: Phonebook{
			homephone: "1234567890",
			cell:      "0987654321",
		},
	}

	fmt.Println(p.firstname)
	fmt.Println(p.lastName)
	fmt.Println("Before increment", p.age)
	p.increageAgeByOne()
	fmt.Println("After increment", p.age)
	fmt.Println(p.homephone)
	fmt.Println(p.cell)
	fmt.Println(p)

	p.address.city = "Noida"

	fmt.Println(p)

	user := struct {
		name string
		age  int
	}{name: "John", age: 30}

	fmt.Println(user)

	fmt.Println(p.fullname())

}

type Person struct {
	firstname string
	lastName  string
	age       int
	address   Address
	Phonebook
}

type Address struct {
	city    string
	country string
}

type Phonebook struct {
	homephone string
	cell      string
}

func (p Person) fullname() string {
	return p.firstname + " " + p.lastName
}

func (a *Person) increageAgeByOne() {
	a.age++
}
