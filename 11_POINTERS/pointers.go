package main 
import (
	"fmt"
)
func  main(){

	jay := " hello wolrd "

    fmt.Println(jay)
	fmt.Println( &jay )
    
	poi  := &jay 
    fmt.Println( poi )
	fmt.Println( *poi )

	var poi2 *string 
	poi2 = &jay 
	fmt.Println( poi2 )
	fmt.Println( *poi2 )


}