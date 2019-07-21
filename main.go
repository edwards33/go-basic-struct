package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName string
	age int
}

func main() {
	firstPerson := person{"Name01", "LastName01", 20}
	secondPerson := person{"Name02", "LastName02", 30}

	fmt.Println(firstPerson.firstName, firstPerson.lastName, firstPerson.age)
	fmt.Println(secondPerson.firstName, secondPerson.lastName, secondPerson.age)
}