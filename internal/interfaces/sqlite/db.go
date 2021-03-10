package sqlite

import (
	"fmt"
	"log"
	"time"

	"github.com/bvinc/go-sqlite-lite/sqlite3"
)

var db *sqlite3.Conn
var tablesStmts []string = []string{
	`CREATE TABLE IF NOT EXISTS job(name TEXT PRIMARY KEY, command TEXT, tick TEXT)`,
	`CREATE TABLE IF NOT EXISTS execution(
		job_name TEXT, date DATE, STATUS TEXT, 
		mem_usage INTEGER, cpu_time INTEGER
	)`,
}

func newDB(name string) {
	connectDB(name)
	createTables()
}

func connectDB(name string) {
	conn, err := sqlite3.Open(name)
	if err != nil {
		log.Fatalf("Cannot establish connection with sqlite db: %v", err)
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
