package main 

import (
 "sort"
 "fmt"
)


type Person struct{
	name string  
	age int
}

type Per2 struct{
	name string  
	age int
	score int 
}

func main(){
  
	ppl := []Person{
		{"a" , 2 } ,
		{"b" , 3} ,
		{"c" , 1},
	}
    sort.Slice( ppl , func(i   , j int) bool{
		return ppl[i].age < ppl[j].age 
	})

	fmt.Println(ppl)

	sort.Slice(ppl , func(i , j int) bool {
		return ppl[i].name < ppl[j].name
	})

	fmt.Println(ppl)

	nums := []int{5, 3, 1, 4, 2}
	sort.Slice( nums , func( i , j  int ) bool{
      return nums[i] > nums[j]
	})

	fmt.Println(nums)

	ppl2 := []Per2 {
		{"Alice", 25, 85},
		{"Bob", 25, 90},
		{"Alice", 25, 92},
		{"Bob", 20, 70},
		{"Charlie", 25, 80},
	}
	sort.Slice(ppl2 , func(i , j int) bool{
		if ppl2[i].age != ppl2[j].age{
			return ppl2[i].age < ppl2[j].age
		}
		
		if ppl2[i].name != ppl2[j].name {
			return ppl2[i].name < ppl2[j].name
		}
       
		return ppl2[i].score > ppl2[j].score
	})

	fmt.Println(ppl2)

}