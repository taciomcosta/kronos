package data

import (
	"fmt"
	"log"
	"time"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

var db *sqlite3.Conn
var tablesStmts []string = []string{
	`CREATE TABLE IF NOT EXISTS job(name TEXT PK, command TEXT, tick TEXT)`,
}

// New connects to database and creates needed tables.
func New() {
	connectDB()
	createTables()
}

func connectDB() {
	conn, err := sqlite3.Open("kronos.db")
	if err != nil {
		log.Fatal("Cannot establish connection with sqlite db")
	}
	conn.BusyTimeout(5 * time.Second)
	db = conn
	fmt.Println("Connection established to db")
}

func createTables() {
	for _, stmt := range tablesStmts {
		err := db.Exec(stmt)
		if err != nil {
			log.Fatalf("Failed to execute statement: %s, %v", stmt, err)
		}
	}
}
