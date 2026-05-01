package main

import (
	"fmt"
	"time"
)

func main() {
	day := time.Now().Weekday()
	hour := time.Now().Hour()
	fmt.Println("Day: ", day, "Hour: ", hour)
	if day.String() == "Monday" {
		fmt.Println("Performing full blown test!")
	} else {
		fmt.Println("Performing hit-n-run test!!!")
	}
}
