package unit_test

import (
	"testing"
)

// func TestCheckDivisibility(t *testing.T) {
// 	input := 5
// 	want := "THREE"
// 	got := CheckDivisibilty(input)
// 	if want != got {
// 		t.Error("Incorrect Response")
// 	}
// }

func TestCheckDivisibilityTableDriven( t *testing.T ){

	testCases := []struct{
		name string 
		input int 
		want string 
	}{
		{
			name : "for input 5 " ,
			input : 5 ,
			want : "FIVE" ,
		} ,
		{
			name : "for input 6 " ,
			input : 6 ,
			want : "THREE" ,
		} ,
	}

	for _ , test := range testCases {
		  
		t.Run( test.name  , func(t *testing.T ){
             got := CheckDivisibilty(test.input)
			 
			 if got != test.want{
				t.Errorf(" incorrect response [input :%v ] [want : %v ] , [ got : %v ] " , test.input , test.want , test.name )
			 }
		})
	}
}
