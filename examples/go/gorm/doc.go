// Package gogorm holds the snippets for the "GORM" pages (Backend Practice)
// in src/content/docs/programming/go/gorm/.
// Tests use an in-memory SQLite database; production code would use gorm.io/driver/postgres.
package gogorm

import (
	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func openTestDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("file::memory:"), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})
	if err != nil {
		panic(err)
	}
	return db
}
