package main

import "fmt"

type Person struct {
	Id   int
	Name string
}

// не изменит оригинальной структуры, для которой вызван
func (p Person) UpdateName(name string) {
	p.Name = name
}

// изменяет оригинальную структуру
func (p *Person) SetName(name string) {
	p.Name = name
}

type Account struct {
	Id   int
	Name string
	Person
}

func (p *Account) SetName(name string) {
	p.Name = name
}

type MySlice []int

func (sl *MySlice) Add(val int) {
	*sl = append(*sl, val)
}

func (sl *MySlice) Count() int {
	return len(*sl)
}

func main() {
	pers := Person{1, "Alex"}
	pers.SetName("Alex Vlasov")
	// (&pers).SetName("Alex Vlasov")
	fmt.Printf("update person: %#v\n", pers)

	var acc Account = Account{
		Id:   1,
		Name: "valex",
		Person: Person{
			Id:   2,
			Name: "Alex Vlasov",
		},
	}
	acc.SetName("vlasov.alex")
	acc.Person.SetName("Test")

	// fmt.Printf("%#v \n", acc)
	sl := MySlice([]int{1, 2})
	sl.Add(5)
	fmt.Println(sl.Count(), sl)

}
