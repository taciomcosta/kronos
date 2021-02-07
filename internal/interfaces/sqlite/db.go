package sqlite

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

func newDB(name string) {
	connectDB(name)
	createTables()
}

func connectDB(name string) {
	conn, err := sqlite3.Open(name)
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