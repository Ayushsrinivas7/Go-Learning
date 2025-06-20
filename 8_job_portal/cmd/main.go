package main

import (
	"fmt"
	"log"
	"job_portal/internal/repository"
	"github.com/gin-gonic/gin"
)
func main(){
    db , err := repository.InitDB()

	if err != nil {
         log.Fatal(err)
	}
	defer db.Close()

	r :=  gin.Default()

	fmt.Println("server satrted")
	r.Run(":8008")

}