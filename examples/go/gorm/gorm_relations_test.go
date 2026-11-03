package gogorm

import (
	"fmt"

	"gorm.io/gorm"
)

type Customer struct {
	gorm.Model
	Name   string
	Orders []Order // has many
}

type Order struct {
	gorm.Model
	CustomerID uint // foreign key
	Customer   Customer
	Total      int
}

func Example_relations() {
	db := openTestDB()
	db.AutoMigrate(&Customer{}, &Order{})

	db.Create(&Customer{Name: "ann", Orders: []Order{{Total: 10}, {Total: 25}}}) // saves the orders too

	var c Customer
	db.Preload("Orders").First(&c, "name = ?", "ann") // load the association in a second query
	fmt.Println(c.Name, len(c.Orders))

	var o Order
	db.Preload("Customer").First(&o) // belongs to
	fmt.Println(o.Total, o.Customer.Name)
	// Output:
	// ann 2
	// 10 ann
}
