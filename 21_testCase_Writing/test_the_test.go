package main

import (
	"fmt"
	"os"
	"testing"
)

func TestMain(  m  *testing.M){
  SetUpDB()
  
 

  code := m.Run()
  TeardownDB()
  os.Exit(code)
}

func TestQueryDatabase(t *testing.T ){
  rows, err := db.Query("SELECT * FROM users")
	if err != nil {
		t.Fatalf("Query failed: %v", err)
	}
	defer rows.Close()

	// Example: Print column names
	cols, _ := rows.Columns()
	fmt.Println("Columns:", cols)

}

