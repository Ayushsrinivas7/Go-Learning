package main


import (
    "fmt"
    "sort"
)




func main(){

     nums := []int{5, 2, 6, 3, 1}

     sort.Ints(nums)

     fmt.Println(nums)

     names := []string{"Charlie", "Alice", "Bob"}

     sort.Sort(   sort.Reverse( sort.StringSlice(names) ) )

     fmt.Println(names) 
}