package standardlibrary

import (
	"fmt"
	"time"
)

func Example_time() {
	t := time.Date(2026, 11, 3, 14, 30, 0, 0, time.UTC)
	fmt.Println(t.Format("2006-01-02 15:04"), t.Format(time.RFC3339), t.Weekday())
	fmt.Println(t.Add(90 * time.Minute).Format(time.Kitchen))

	d, _ := time.ParseDuration("1h30m")
	xmas, _ := time.Parse("02/01/2006", "25/12/2026") // layout uses the reference date
	fmt.Println(d.Minutes(), xmas.Sub(t).Hours())
	// Output:
	// 2026-11-03 14:30 2026-11-03T14:30:00Z Tuesday
	// 4:00PM
	// 90 1233.5
}
