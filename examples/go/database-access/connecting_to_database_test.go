package databaseaccess

import (
	"context"
	"database/sql"
	"fmt"
	"time"
)

// openDB is how a service opens its pool at startup.
func openDB(ctx context.Context, dsn string) (*sql.DB, error) {
	db, err := sql.Open("pgx", dsn) // validates arguments only; no connection yet
	if err != nil {
		return nil, err
	}
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()
	if err := db.PingContext(ctx); err != nil { // first real connection
		db.Close()
		return nil, fmt.Errorf("ping db: %w", err)
	}
	return db, nil
}

var _ = openDB // used by main in a real service
