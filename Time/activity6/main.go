package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("The script started at: ", start)
	sum := 0
	for i := 1; i < 10000000000; i++ {
		sum += i
	}
	end := time.Now()
	duration := end.Sub(start)
	fmt.Println("The script completed at: ", end)
	fmt.Println("The task took", duration.Minutes(), "minutes(s) to complete!!!")
	fmt.Println("The task took", duration.Seconds(), "seconds(s) to complete!!!")
	fmt.Println("The task took", duration.Nanoseconds(), "nanoseconds(s) to complete!!!")
}
