package main

import "fmt"

type Personn struct {
    name string
    age  int
}

func (p Personn) greet() {
    fmt.Println("Hello,", p.name)
}

func (p *Personn) haveBirthday() {
    p.age++
}

func (p *Personn ) details() string {
    // return  fmt.Sprintf("person name is %v and age is %v " , p.name , p.age )
	return "hello"
}

func  main() {
    p := Personn{name: "Mike", age: 20}
    p.greet()
    p.haveBirthday()
    fmt.Println("After birthday:", p.age)
	fmt.Println(p.details()) 
	// p.details()
}
