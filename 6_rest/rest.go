package main 

import (
	"io"
	"fmt" 
	"net/http"
)
func main(){

   http.HandleFunc("/"  , handleRoot )
   fmt.Println("server is running in the port 8080 ")
   http.ListenAndServe(":8080"  , nil )

}

func handleRoot(  w http.ResponseWriter , r *http.Request ){
  fmt.Fprintf( w , "hello wrold ")
  body, err := io.ReadAll(r.Body)
  if err != nil {
    fmt.Println("Error reading body:", err)
    return
 }
fmt.Println("Body:", string(body))
}