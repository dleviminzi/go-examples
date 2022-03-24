package main

import (
	"fmt"
	"time"
)

func main() {
	t0 := time.Date(2021, time.August, 12, 0, 0, 0, 0, &time.Location{})
	t1 := time.Date(2021, time.August, 16, 0, 0, 0, 0, &time.Location{})

	for t0.Before(t1) || t0.Equal(t1) {
		fmt.Print(t0.Format("2006-01-02 15:04:05.999999999"))
		t0 = t0.AddDate(0, 0, 1)
	}
}
