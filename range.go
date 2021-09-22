package main

import "fmt"

type Person struct {
	Name string
	Age int
}

func main() {
	persons := []Person{}
	persons = append(persons, Person{Name: "zhangshan", Age: 30})
	persons = append(persons, Person{Name: "lisi", Age: 10})
	persons = append(persons, Person{Name: "wangwu", Age: 18})
	persons = append(persons, Person{Name: "zhaoliu", Age: 70})

	adults := []*Person{}
	for _, p := range persons {
		if p.Age >= 18 {
			adults = append(adults, &p)
		}
	}
	for _, adult := range adults {
		fmt.Println("name: ", adult.Name, ", age: ", adult.Age)
	}
	fmt.Println()
	adults = []*Person{}
	for i := range persons {
		if persons[i].Age >= 18 {
			adults = append(adults, &persons[i])
		}
	}
	for _, adult := range adults {
		fmt.Println("name: ", adult.Name, ", age: ", adult.Age)
	}
}