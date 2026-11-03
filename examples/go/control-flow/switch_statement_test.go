package controlflow

import "fmt"

func dayKind(day string) string {
	switch day {
	case "sat", "sun":
		return "weekend"
	default:
		return "weekday"
	}
}

func Example_switch() {
	fmt.Println(dayKind("sun"), dayKind("mon"))

	x := 15
	switch { // no tag: same as switch true
	case x > 10:
		fmt.Println("big")
		fallthrough // runs the next case body without checking it
	case x > 100:
		fmt.Println("huge?")
	default:
		fmt.Println("small")
	}
	// Output:
	// weekend weekday
	// big
	// huge?
}
