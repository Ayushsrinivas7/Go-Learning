package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

type User struct{

	id int 

	username  string 

	email string 
 
	mobile string 

	address Address 

}

type Address struct {

	id int 

	street string 

	city string 

	state string 

	country string 
}

func main() {
	// Creating a struct using a struct literal
	p1 := Person{Name: "Ayush", Age: 22}
	
	// Accessing fields
	fmt.Println(p1.Name)  // Output: Ayush
	fmt.Println(p1.Age)   // Output: 22


	// nested struct 

	u1 := User{
		id : 1 ,
		username : "ayush" ,
		email : "ayushsrinivas7@gmail.com",
		mobile : "9390060699",
		address : Address{
			id : 1,
			street : "a4b",
            city : "a4b",
            state : "a4b",
            country: "a4b",
		},
	}

	fmt.Println(u1.address)


}
