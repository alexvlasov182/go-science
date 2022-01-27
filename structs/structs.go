package main

import "fmt"

type Person struct {
	Id      int
	Name    string
	Address string
}

type Account struct {
	Id      int
	Name    string
	Cleaner func(string) string
	Owner   Person
	Person
}

func main() {
	// полное обьявление структуры
	var acc Account = Account{
		Id:   1,
		Name: "valex",
		Person: Person{
			Name:    "Александр",
			Address: "Берн",
		},
	}
	fmt.Printf("%#v\n", acc)

	// короткое обьявление структуры
	acc.Owner = Person{2, "Vlasov Alex", "Bern"}

	fmt.Printf("%v#\n", acc)

	fmt.Println(acc.Name)
	fmt.Println(acc.Person.Name)
	fmt.Println(acc.Person.Address)
}
