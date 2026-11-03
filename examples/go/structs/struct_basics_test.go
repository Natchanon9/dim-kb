package structs

import "fmt"

type User struct {
	Name string
	Age  int
}

func Example_structBasics() {
	var u User // all fields at their zero value
	u.Name = "Ann"
	fmt.Println(u, u.Age)
	// Output: {Ann 0} 0
}
