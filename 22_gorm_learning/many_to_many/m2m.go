package main

import (
	// "fmt"
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

/* Sample Data:
   <Movie>: 					<List of Actors>
   "Iron Man": 					Robert Downey Jr.
   "Avengers": 					Robert Downey Jr., Chris Evans, Scarlett Johansson
   "Black Panther": 			Chadwick Boseman
   "Avengers Infinity War": 	Robert Downey Jr., Chris Evans, Scarlett Johansson, and Chadwick Boseman
   "Sherlock Holmes": 			Robert Downey Jr.
   "Lost in Translation": 		Scarlett Johansson
   "Marriage Story": 			Scarlett Johansson
*/

type Movie struct {
	gorm.Model
	Name   string
	Actors []Actor `gorm:"many2many:filmography;"`
}
//  note here the fimilography tabel is created byt he goem whivch is 
// is used to maintain the id of the actor and id of the movie so that they can be maintained for  mahy to many relation 
//  we write the join query using the filimography tabel 

type Actor struct {
	gorm.Model
	Name   string
	Movies []Movie `gorm:"many2many:filmography;"`
}

var DB *gorm.DB

func connectDatabase() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Microsecond, // Slow SQL threshold
			LogLevel:                  logger.Info,      // Log level
			IgnoreRecordNotFoundError: true,             // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,             // Disable color
		},
	)
	database, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/testdb1?charset=utf8&parseTime=true"), &gorm.Config{Logger: newLogger})

	if err != nil {
		panic("Failed to connect to databse!")
	}

	DB = database
}

func dbMigrate() {
	DB.AutoMigrate(&Movie{}, &Actor{})
}

func main() {
	connectDatabase()
	// dbMigrate()
    // insertSampleData()

	var movie Movie
	// note the gorm will first bring move then actirs 
	DB.Where("name = ? " , "Avengers Infinity War").Preload("Actors").First(&movie)
    
	fmt.Println(movie.Name)

	for _ , actor := range movie.Actors {
		fmt.Printf(" --> %v \n ", actor.Name )
	}


	var actor Actor 

	DB.Where("name = ?" , "Robert Downey Jr.").Preload("Movies").First(&actor)
    
	fmt.Println(actor.Name)

	for _ , mov := range actor.Movies{
    fmt.Printf(" --> %v \n ", mov.Name )
	}

}

func insertSampleData() {
	// Define actors
	rdj := Actor{Name: "Robert Downey Jr."}
	evans := Actor{Name: "Chris Evans"}
	scarlett := Actor{Name: "Scarlett Johansson"}
	chadwick := Actor{Name: "Chadwick Boseman"}

	// Save actors first (to reuse)
	DB.Create(&rdj)
	DB.Create(&evans)
	DB.Create(&scarlett)
	DB.Create(&chadwick)

	// Define movies and associate actors
	movies := []Movie{
		{Name: "Iron Man", Actors: []Actor{rdj}},
		{Name: "Avengers", Actors: []Actor{rdj, evans, scarlett}},
		{Name: "Black Panther", Actors: []Actor{chadwick}},
		{Name: "Avengers Infinity War", Actors: []Actor{rdj, evans, scarlett, chadwick}},
		{Name: "Sherlock Holmes", Actors: []Actor{rdj}},
		{Name: "Lost in Translation", Actors: []Actor{scarlett}},
		{Name: "Marriage Story", Actors: []Actor{scarlett}},
	}

	// Save movies with associations
	for _, movie := range movies {
		err := DB.Create(&movie).Error
		if err != nil {
			log.Println("Error inserting movie:", movie.Name, err)
		}
	}
}
