package databaseaccess

import (
	"database/sql"
	"fmt"
)

func Example_null() {
	var phone sql.NullString
	phone.Scan(nil) // what rows.Scan does for a NULL column
	fmt.Println(phone.Valid, phone.String == "")

	var age sql.Null[int64] // generic version, Go 1.22+
	age.Scan(int64(30))
	fmt.Println(age.Valid, age.V)

	var nick *string // a pointer also works: nil means NULL
	fmt.Println(nick == nil)
	// Output:
	// false true
	// true 30
	// true
}
