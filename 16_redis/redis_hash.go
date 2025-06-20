package main 


import (
"context"
	"encoding/csv"
	"fmt"
	"log"
	"os"
"github.com/go-redis/redis/v8"
)

var ctx = context.Background()

func main(){


	 rdb := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", 
        DB:       0,  
    })

	fmt.Println(rdb)

	// 2. Open the CSV
	file, err := os.Open("pincode.csv") 
	if err != nil {
		log.Fatalf(" Failed to open CSV: %v", err)
	}
	defer file.Close()

	// 3. Read CSV
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf(" Failed to read CSV: %v", err)
	}

	// 4. Skip Header and insert into Redis

	for i, row := range records {
		if i == 0 {
			continue // Skip header
		}

		pincode := row[0]
		err := rdb.HSet(ctx, "pincode_data", pincode, "1").Err()
		if err != nil {
			log.Printf(" Failed to insert pincode %s: %v", pincode, err)
		}
	}
	fmt.Println("📦 All pincodes loaded into Redis Hash!")

	// 5. Example Query
	fmt.Println(GetPincode(rdb, "110000"))
	fmt.Println(SetPincode(rdb, "2101ee"))
    fmt.Println(checkPincode(rdb, "2101ee"))
}

func checkPincode(rdb *redis.Client, pincode string) (bool  , error ) {
	exists, err := rdb.HExists(ctx, "pincode_data", pincode).Result()
	if err != nil {
		log.Fatalf(" Redis error: %v", err)
		return  false , err
	}

	if exists != true  {
       return false , nil  
	}

	fmt.Printf("🔎 Does pincode %s exist? %v\n", pincode, exists)
	return true , nil 
}

func GetPincode(rdb *redis.Client, pincode string ) string {
	exists, err := rdb.HGet(ctx, "pincode_data", pincode).Result()
	if err != nil {
		log.Fatalf(" Redis error: %v", err)
	}
	fmt.Printf("🔎 Does pincode %s exist? %v\n", pincode, exists)
	return exists ;
}

func SetPincode(rdb *redis.Client, pincode string ) error  {
	err := rdb.HSet(ctx, "pincode_data", pincode, "1").Err()
	if err != nil {
		log.Fatalf(" Redis error: %v", err)
		return err 
	}
    return nil 
}

