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

//  note : in users tabel you will only see
// id , username , password only u will not see notes , creditcard
//  notes and credit card if you need
//  then you need DB.Preload("Notes").Preload("CreditCard").First(&user, userID)
//  here the relation is saved like this only
type User struct {
	gorm.Model
	ID         uint64 `gorm:"primaryKey"`
	Username   string `gorm:"size:64"`
	Password   string `gorm:"size:255"`
	Notes      []Note
	CreditCard *CreditCard
}

// user has 1 -> many with notes 
// user has 1 -> 1 with creditcard 

// if ypou add the * then that field can be nil , it saves memory 
// if you have the not added the feild can not be empty 


type Note struct {
	gorm.Model
	ID      uint64 `gorm:"primaryKey"`
	Name    string `gorm:"size:255"`
	Content string `gorm:"type:text"`
	UserID  uint64 `gorm:"index"`
	User    User
}

type CreditCard struct {
	gorm.Model
	Number string
	UserID uint64
	User   User
}

var DB *gorm.DB

func connectDatabase() {
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // Slow SQL threshold
			LogLevel:                  logger.Info, // Log level
			IgnoreRecordNotFoundError: true,        // Ignore ErrRecordNotFound error for logger
			Colorful:                  true,        // Disable color
		},
	)
	database, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/testdb1?charset=utf8&parseTime=true"), &gorm.Config{Logger: newLogger})

	if err != nil {
		panic("Failed to connect to databse!")
	}

	DB = database
}

func dbMigrate() {
	DB.AutoMigrate(&Note{}, &User{}, &CreditCard{})
}

func main() {
	connectDatabase()
	// dbMigrate()
    // insertMockData()

	var note Note 
    
	// when running the preload the flow is as flows 
    // 1->	SELECT * FROM `notes` WHERE `notes`.`deleted_at` IS NULL ORDER BY `notes`.`id` LIMIT 1
	// 2 ->  SELECT * FROM `users` WHERE `users`.`id` = 1 AND `users`.`deleted_at` IS NULL
	// in 2 the user id is brought from the notes.userid

	DB.Preload("User").First(&note)

	fmt.Printf("User from a note : %s \n " , note.User.Username)

    
   var user User 

   DB.Preload("Notes").Preload("CreditCard").Where("username = ?"  , "ayush").First(&user)

   fmt.Println(" NOtes are as follw ")

   for _ , note :=  range user.Notes{
	   fmt.Println(note.Content)
   }

   fmt.Printf("CreditCard from  user : %v \n" , user.CreditCard)

	
}

func insertMockData() {
	// Create Users
	users := []User{
		{Username: "ayush", Password: "secret123"},
		{Username: "elon", Password: "tesla42"},
		{Username: "steve", Password: "apple1984"},
	}

	for i := range users {
		if err := DB.Create(&users[i]).Error; err != nil {
			log.Println("Error creating user:", err)
		}
	}

	// Create Credit Cards (1 per user)
	cards := []CreditCard{
		{Number: "4111-1111-1111-1111", UserID: users[0].ID},
		{Number: "4222-2222-2222-2222", UserID: users[1].ID},
		{Number: "4333-3333-3333-3333", UserID: users[2].ID},
	}

	for i := range cards {
		if err := DB.Create(&cards[i]).Error; err != nil {
			log.Println("Error creating credit card:", err)
		}
	}

	// Create Notes for users
	notes := []Note{
		{Name: "Shopping List", Content: "Milk, Eggs, Bread", UserID: users[0].ID},
		{Name: "Goals", Content: "Learn Go, Build app", UserID: users[0].ID},
		{Name: "To-Do", Content: "Launch rocket", UserID: users[1].ID},
		{Name: "Ideas", Content: "New iPhone concept", UserID: users[2].ID},
	}

	for i := range notes {
		if err := DB.Create(&notes[i]).Error; err != nil {
			log.Println("Error creating note:", err)
		}
	}

	log.Println("✅ Mock data inserted.")
}
