package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
	"github.com/google/uuid"
)

type User struct {
	Id string 
	Name  string `json:"name"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

func main() {
    fmt.Println("Go Redis Tutorial")

    client := redis.NewClient(&redis.Options{
        Addr:     "localhost:6379",
        Password: "", // no password set
        DB:       0,  // use default DB
    })
	fmt.Println(client)

ping, err := client.Ping(context.Background()).Result()
if err != nil {
    fmt.Println(err.Error())
    return
}

    fmt.Println("Redis ping response:", ping)


	err = client.Set(context.Background(), "name", "Elliot", 0).Err()
if err != nil {
    fmt.Printf("Failed to set value in the redis instance: %s", err.Error())
    return
}

val, err := client.Get(context.Background(), "name").Result()
if err != nil {
    fmt.Printf("Failed to get value from redis: %s", err.Error())
}

fmt.Println(val )

     user := User{
	    Id : uuid.NewString(),
		Name:  "Ayush",
		Email: "ayush@example.com",
		Age:   25,
	}

	// Convert Go struct to JSON
	jsonBytes, err := json.Marshal(user)
	if err != nil {
		fmt.Println("Error marshaling:", err)
		return
	}

	fmt.Println(string(jsonBytes)) 



	err = client.Set(context.Background(), user.Id, jsonBytes, 0).Err()
if err != nil {
    fmt.Printf("Failed to set value in the redis instance: %s", err.Error())
    return
}

val, err = client.Get(context.Background(),  user.Id).Result()
if err != nil {
    fmt.Printf("Failed to get value from redis: %s", err.Error())
}
fmt.Println(val)


var ctx = context.Background()

// Add fields to a hash
	err = client.HSet(ctx, "user:1", "name", "Ayush", "age", "25", "city", "Bangalore").Err()
	if err != nil {
		panic(err)
	}

	// Get a single field
	name, err := client.HGet(ctx, "user:1", "name").Result()
	fmt.Println("Name:", name)
	if err != nil {
		panic(err)
	}
	// Get all fields
	allFields, err := client.HGetAll(ctx, "user:1").Result()
	fmt.Println("All fields:", allFields)
    if err != nil {
		panic(err)
	}
	// Check if field exists
	exists, _ := client.HExists(ctx, "user:1", "age").Result()
	fmt.Println("Age field exists?", exists)

}
