// Package databaseaccess holds the snippets for the "Database Access" pages
// in src/content/docs/programming/go/database-access/.
// Functions that need a live PostgreSQL are compiled but not run.
package databaseaccess

import "time"

type User struct {
	ID        int64
	Name      string
	Email     string
	CreatedAt time.Time
}
