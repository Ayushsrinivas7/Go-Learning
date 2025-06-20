package main

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/golang-jwt/jwt/v5"
)


var jwtSecret = []byte("mysecret123")

func GenerateToken(userID uint , email string  , role string ) (string , error ){
	// claims ie the data which we want to store  in the token 

	claims := jwt.MapClaims{
          "userID" : userID ,
		  "email" : email ,
		  "role" : role ,
		  "exp" :  time.Now().Add(time.Hour * 24 * 30).Unix(),
	}

	token := jwt.NewWithClaims( jwt.SigningMethodHS256  , claims ) 

	tokenStr , err  := token.SignedString(jwtSecret)

	if err != nil {

		return "" , errors.New("we got an error while jwt generation ")
	}

	return tokenStr , nil 


}

func VerifyToken(tokenStr string ) (jwt.MapClaims, error) {
	// check same hashing algo or not 
       token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// Check the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		// thing is we have to see the secret key andother  things are in the bytes format ie the []bytes(secret string )
		return jwtSecret , nil
	})
     
	if err != nil || !token.Valid {
		return nil, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, errors.New("could not parse claims")
	}
	return claims, nil

}

func main(){
    token , err := GenerateToken( 1 , "ayushsrinivas7@gmail.com"  , "admin" )

	if err !=  nil {
		log.Printf("%v" , err )
		return 
	}

	fmt.Println(" generated token is ", token )

	claims , err := VerifyToken(token)

	if err !=  nil {
		log.Printf("%v" , err )
		return 
	}

    fmt.Println("\n✅ Token Verified! Claims:")

	for k , v := range claims {
    fmt.Printf("\n %v %v " , k , v )
	}

}