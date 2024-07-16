package main

import (
	"fmt"
	"github.com/natealcedo/go-goose/models"
)

func main() {
	db, err := models.CreateClient()

	if err != nil {
		panic(err)
	}

	var testTables []models.TestTable
	result := db.DB.Table("test_tables").Find(&testTables) // Select all records from test_table
	if result.Error != nil {
		panic(result.Error)
	}

	// Print the results
	for _, testTable := range testTables {
		fmt.Printf("ID: %d, Name: %s\n", testTable.ID, testTable.Name)
	}

}
