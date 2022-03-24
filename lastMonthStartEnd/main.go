package main

import (
	"fmt"
	"time"
)

func main() {
	currTime := time.Now()

	lastMonth := currTime.AddDate(0, -1, 0)
	daysPassed := currTime.Day()

	firstDate := lastMonth.AddDate(0, 0, -daysPassed+1)
	lastDate := lastMonth.AddDate(0, 1, -daysPassed)

	firstTime := time.Date(firstDate.Year(), firstDate.Month(), firstDate.Day(), 0, 0, 0, 0, firstDate.UTC().Location())
	lastTime := time.Date(lastDate.Year(), lastDate.Month(), lastDate.Day(), 0, 0, 0, 0, lastDate.UTC().Location())
	fmt.Println(firstTime, lastTime)
}
