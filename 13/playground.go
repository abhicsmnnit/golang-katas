package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())

	d := 2*time.Hour + 45*time.Minute + 100*time.Microsecond
	fmt.Println(d)
	fmt.Println(d.String())

	d1, _ := time.ParseDuration("25h50m10s100ms100us100ns")
	fmt.Println(d1.Hours())
	fmt.Println(d1.Round(time.Hour))
	fmt.Println(d1.Seconds())
	fmt.Println(d1.Round(time.Second))

	t, _ := time.Parse("2006-01-02 15:04:05 -0700", "2023-03-13 00:00:00 +0530")
	fmt.Println(t.Format("January 2, 2006 at 3:04:05PM MST"))

	time.AfterFunc(2*time.Second, func() { fmt.Println("Yo, people!") })
	time.Sleep(3 * time.Second)
}
