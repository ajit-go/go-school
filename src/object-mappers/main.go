package main

import (
	"fmt"

	"github.com/devfeel/mapper"
)

type Emp struct {
	Name  string
	Class int
}
type Emp2 struct {
	FirstName string `mapper:"Name"`
	Class     int
}

func main() {
	e1 := Emp{Name: "Ajit", Class: 1}

	var e2 Emp2
	if err := mapper.Mapper(&e1, &e2); err != nil {
		fmt.Println(err)
	}
	fmt.Printf("e1:= %v", e1)
	fmt.Printf("e2:= %v", e2)
	e2 = Emp2{}
	a1 := []Emp{
		Emp{Name: "Ajit", Class: 1},
		Emp{Name: "vinit", Class: 2},
	}
	a2 := make([]Emp, len(a1))
	for i, v := range []int{11, 22, 33, 44} {
		fmt.Println(i, v)
	}
	
	for i, a := range a1 {
		if err := mapper.Mapper(&a, &a2[i]); err != nil {
			fmt.Println(err)
		}
	}
	 fmt.Printf("a1:= %v", a1)
	 fmt.Printf("a2:= %v", a2)
}
