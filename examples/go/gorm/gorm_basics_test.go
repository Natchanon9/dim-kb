package gogorm

import (
	"errors"
	"fmt"

	"gorm.io/gorm"
)

type Product struct {
	gorm.Model        // ID, CreatedAt, UpdatedAt, DeletedAt
	Code       string `gorm:"uniqueIndex"`
	Price      int
}

func Example_basics() {
	db := openTestDB()
	db.AutoMigrate(&Product{}) // create or alter the table to match the struct

	db.Create(&Product{Code: "P1", Price: 100})

	var p Product
	db.First(&p, "code = ?", "P1")
	fmt.Println(p.ID, p.Code, p.Price)

	db.Model(&p).Update("Price", 200)
	db.First(&p, p.ID)
	fmt.Println(p.Price)

	db.Delete(&p) // soft delete: sets DeletedAt
	err := db.First(&Product{}, p.ID).Error
	fmt.Println(errors.Is(err, gorm.ErrRecordNotFound))
	// Output:
	// 1 P1 100
	// 200
	// true
}
