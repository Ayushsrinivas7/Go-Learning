package main 

import (
	"fmt"
)

func main(){

	myMap:=  make( map[string]int )
	myMap["apple"] = 3
    fmt.Println(myMap["apple"])

	value, ok := myMap["orange"]
if ok {
    fmt.Println("Found:", value)
} else {
    fmt.Println("Not found")
}

// curd in map 

mp := make( map[string]int )

mp["1"] = 1
mp["2"] = 2
val , ok := mp["1"]

if ok {
	fmt.Println(val)
}else{
	fmt.Println("not found ")
}


delete(mp , "2" )








 
}