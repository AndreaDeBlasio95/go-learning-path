package db

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"
)

// DB is a global variable that holds the database connection
var DB *sql.DB


func InitDB() {
	var err error
	DB, err = sql.Open("sqlite3", "api.db") 

	if err != nil {
		panic("Could not open database: " + err.Error())
	}

	// Check if the database connection is actually working
	if err = DB.Ping(); err != nil {
		panic("Could not establish a connection to the database: " + err.Error())
	}

	DB.SetMaxOpenConns(10) // Set the maximum number of open connections to 10
	DB.SetMaxIdleConns(5)  // Set the maximum number of idle connections to 5

	createTables() // Call the function to create the tables
}


func createTables() {
	createEventsTable := `
	CREATE TABLE IF NOT EXISTS events (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL,
		description TEXT NOT NULL,
		location TEXT NOT NULL,
		datetime DATETIME NOT NULL,
		user_id INTEGER
		)
		`
	
	// Execute the SQL statement to create the events table
	_, err := DB.Exec(createEventsTable)

	if err != nil {
		panic("Could not create events table")
	}
}