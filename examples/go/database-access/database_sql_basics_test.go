package databaseaccess

import (
	"database/sql"
	"fmt"
	"slices"

	_ "github.com/jackc/pgx/v5/stdlib" // registers the "pgx" driver
)

func Example_drivers() {
	fmt.Println(slices.Contains(sql.Drivers(), "pgx"))
	// Output: true
}
