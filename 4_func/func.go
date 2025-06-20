package main

import "fmt"

func div( a   , b int ) (int , int ){
	q := a / b 
	 r := a%b
	 return q , r   
}

func main(){     
  fmt.Println( div( 4 , 3 ) )
  fmt.Println( sum(1 , 2, 3 ) )

  a := func( nums ...int) int {
	 tot := 0
	 for _ , val := range nums{
		tot += val
	 }
	 return tot 
  }

  fmt.Println(a( 1 ,2 , 3 , 4 ))  

  sqr := func( a int ) int {
	return a * a 
  }

  cub := func( a int ) int {
	return a* a* a 
  }

  fmt.Println( opr(2  , sqr) )
 fmt.Println( opr(2  , cub) )
}



func sum( nums ...int )  int {
	tot := 0 

	for _ , val := range nums {
     	tot += val 
	}
	return tot 
	
}


// function as the arguments 
func opr(  a int , opf func(int ) int ) int {
	return opf(a) 
}