package main

import (
	"os"       // for file creation, opening, closing
	"net/http"
	"path/filepath"

	"github.com/gin-gonic/gin"
)


func main(){
  
	r := gin.Default()
    r.GET("/hello" , helloHandler )
	r.GET("/user"  , userHandler )
	r.GET("/items"  , getItemHandler )
	r.GET("/file"  , fileHandler )
    // file name is path parameter 
	r.GET("/downloadfile/:filename"  , downloadFileHandler )
	// we can also get the query parameer 
	//  in the post or get 
	r.GET("/search" , queryHandler )
	r.Run(":8082")

}

// returning to string 


func helloHandler( c *gin.Context){
   c.String(http.StatusOK , "hello wolrd")
}


// returning to json 
func userHandler( c *gin.Context){
	user := gin.H{
		"name" : "jhon" ,
		"age"  : "22" ,
	}
	c.JSON(http.StatusOK , user )
}

// returning the array of objects 

func getItemHandler(c *gin.Context) {
	
    items := []gin.H{
        {"id": 1, "name": "Item 1"},
        {"id": 2, "name": "Item 2"},
        {"id": 3, "name": "Item 3"},
    }

    c.JSON(http.StatusOK , items )

}


func fileHandler(c *gin.Context ){
    
	c.File("./text.txt")
	
}

// func downloadFileHandler(c *gin.Context) {
//     filename := c.Param("filename")
//     baseDir := "./allFiles"
//     filePath := filepath.Join(baseDir, filename)

//     // Check if file exists
//     if _, err := os.Stat(filePath); os.IsNotExist(err) {
//         c.String(http.StatusNotFound, "File not found")
//         return
//     }

//     // Set the correct content disposition header
//     c.Header("Content-Disposition", "attachment; filename=" + filename)
//     c.Header("Content-Type", "application/octet-stream") // force download for all types

//     // Serve the file
//     c.File(filePath)
// }

func downloadFileHandler(c *gin.Context) {
    filename := c.Param("filename")
    baseDir := "./allFiles"
    filePath := filepath.Join(baseDir, filename)

    if _, err := os.Stat(filePath); os.IsNotExist(err) {
        c.String(http.StatusNotFound, "File not found")
        return
    }

    c.FileAttachment(filePath, filename)
}


func  queryHandler(c *gin.Context ){
	name := c.DefaultQuery("name"  , "Guest")
	c.String( http.StatusOK , "Hello %s!" , name  )

}



