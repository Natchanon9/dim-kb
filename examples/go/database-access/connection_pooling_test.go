package databaseaccess

import (
	"database/sql"
	"fmt"
	"time"
)

func Example_pool() {
	db, err := sql.Open("pgx", "postgres://app:secret@localhost:5432/app") // no connection made yet
	if err != nil {
		panic(err)
	}
	defer db.Close()

	db.SetMaxOpenConns(10)                  // hard cap on connections
	db.SetMaxIdleConns(5)                   // kept warm between requests
	db.SetConnMaxLifetime(30 * time.Minute) // recycle long-lived connections
	db.SetConnMaxIdleTime(5 * time.Minute)

	fmt.Println(db.Stats().MaxOpenConnections, db.Stats().OpenConnections)
	// Output: 10 0
}
