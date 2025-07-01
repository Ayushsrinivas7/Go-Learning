// examples of the scope 

// 🧪 Example 1: Basic Filter

// func ActiveUsers(db *gorm.DB) *gorm.DB {
// 	return db.Where("status = ?", "active")
// }

// var users []User
// DB.Scopes(ActiveUsers).Find(&users)


// 🧪 Example 2: Pagination

// func Paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		offset := (page - 1) * pageSize
// 		return db.Offset(offset).Limit(pageSize)
// 	}
// }


// var users []User
// DB.Scopes(Paginate(2, 10)).Find(&users) // Get page 2, 10 users per page


// 🧪 Example 3: Combining Multiple Scopes

// func ByStatus(status string) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		return db.Where("status = ?", status)
// 	}
// }

// func SortBy(field string) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		return db.Order(field)
// 	}
// }

// DB.Scopes(ByStatus("active"), SortBy("created_at desc")).Find(&users)

//  example > range query 

// func CreatedBetween(start, end time.Time) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		return db.Where("created_at BETWEEN ? AND ?", start, end)
// 	}
// }

// start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
// end := time.Now()
// var movies []Movie
// DB.Scopes(CreatedBetween(start, end)).Find(&movies)

// Example -> name alike 

// func SearchByName(keyword string) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		return db.Where("name LIKE ?", "%"+keyword+"%")
// 	}
// }

// var movies []Movie
// DB.Scopes(SearchByName("Avengers")).Find(&movies)


//  VVVIP  - DYNAMIC FILTER 

// type MovieFilter struct {
// 	Keyword    string     // optional
// 	StartDate  *time.Time // optional
// 	EndDate    *time.Time // optional
// 	WithActors bool       // optional
// 	OrderBy    string     // optional: e.g., "created_at desc"
// }

// func ApplyMovieFilters(filter MovieFilter) func(db *gorm.DB) *gorm.DB {
// 	return func(db *gorm.DB) *gorm.DB {
// 		if filter.Keyword != "" {
// 			db = db.Where("name LIKE ?", "%"+filter.Keyword+"%")
// 		}
// 		if filter.StartDate != nil && filter.EndDate != nil {
// 			db = db.Where("created_at BETWEEN ? AND ?", filter.StartDate, filter.EndDate)
// 		}
// 		if filter.WithActors {
// 			db = db.Preload("Actors")
// 		}
// 		if filter.OrderBy != "" {
// 			db = db.Order(filter.OrderBy)
// 		}
// 		return db
// 	}
// }


// func queryWithDynamicFilter() {
// 	start := time.Date(2020, 1, 1, 0, 0, 0, 0, time.UTC)
// 	end := time.Now()

// 	filter := MovieFilter{
// 		Keyword:    "Avengers",
// 		StartDate:  &start,
// 		EndDate:    &end,
// 		WithActors: true,
// 		OrderBy:    "created_at desc",
// 	}

// 	var movies []Movie
// 	err := DB.Scopes(ApplyMovieFilters(filter)).Find(&movies).Error
// 	if err != nil {
// 		log.Println("Query error:", err)
// 		return
// 	}

// 	for _, m := range movies {
// 		fmt.Println("🎬", m.Name)
// 		if filter.WithActors {
// 			for _, a := range m.Actors {
// 				fmt.Println("  👤", a.Name)
// 			}
// 		}
// 	}
// }

// -> AAACTUAL EXAMPLE 

package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type User struct {
	gorm.Model
	Name   string
	Email  string
	Orders []Order
}

type Order struct {
	gorm.Model
	UserId      int64
	OrderTime   time.Time
	PaymentMode string // Card or Cash
	Price       int
	User        User
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
	database, err := gorm.Open(mysql.Open("root:@tcp(127.0.0.1:3306)/my_db_name?charset=utf8&parseTime=true"), &gorm.Config{Logger: newLogger})

	if err != nil {
		panic("Failed to connect to databse!")
	}

	DB = database
}

func dbMigrate() {
	DB.AutoMigrate(&User{}, &Order{})
}

func CardOrders(db *gorm.DB) *gorm.DB {
	return db.Where("payment_mode = ?", "card")
}

func PriceGreaterThan30(db *gorm.DB) *gorm.DB {
	return db.Where("price > ?", 30)
}

func UsersFromDomain(domain string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		return db.Where("email like ?", "%"+domain)
	}
}

func main() {
	connectDatabase()
	// dbMigrate()
    // seedMockData()

	// var orders []Order
	// DB.Scopes(CardOrders, PriceGreaterThan30).Find(&orders)

	// fmt.Println("orders:")
	// for _, order := range orders {
	// 	fmt.Printf("Price: %d, Payment Type: %s\n", order.Price, order.PaymentMode)
	// }

	var users []User
	DB.Scopes(UsersFromDomain("example.com")).Preload("Orders", CardOrders).Find(&users)

	fmt.Printf("Users: \n")
	for _, user := range users {
		fmt.Printf("User email: %s\n", user.Email)
	}

	fmt.Printf("Orders from a user (%s): \n", users[0].Email)
	for _, order := range users[0].Orders {
		fmt.Printf("Price: %d, Payment Type: %s\n", order.Price, order.PaymentMode)
	}
}

func seedMockData() {
	
	// Create mock users
	users := []User{
		{Name: "Alice", Email: "alice@example.com"},
		{Name: "Bob", Email: "bob@example.com"},
		{Name: "Charlie", Email: "charlie@gmail.com"},
	}

	for i := range users {
		DB.Create(&users[i])
	}

	// Create mock orders for each user
	orders := []Order{
		{UserId: int64(users[0].ID), OrderTime: time.Now().Add(-48 * time.Hour), PaymentMode: "card", Price: 100},
		{UserId: int64(users[0].ID), OrderTime: time.Now().Add(-24 * time.Hour), PaymentMode: "cash", Price: 20},
		{UserId: int64(users[1].ID), OrderTime: time.Now().Add(-72 * time.Hour), PaymentMode: "card", Price: 50},
		{UserId: int64(users[2].ID), OrderTime: time.Now().Add(-96 * time.Hour), PaymentMode: "cash", Price: 200},
	}

	for i := range orders {
		DB.Create(&orders[i])
	}

	fmt.Println("Mock data seeded successfully.")
}
