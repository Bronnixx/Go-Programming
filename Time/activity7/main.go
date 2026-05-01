package main

import (
	"fmt"
	"time"
)

func main() {
	deadlineSeconds := time.Duration((600 * 10) * time.Millisecond)
	start := time.Now()
	fmt.Println("Deadline for the transaction is", deadlineSeconds)
	sum := 0
	for i := 1; i < 25000000000; i++ {
		sum += i
	}
	end := time.Now()
	duration := end.Sub(start)
	transactionTime := time.Duration(duration.Nanoseconds()) * time.Nanosecond
	fmt.Println("The transaction has completed at:", end, duration)
	if transactionTime <= deadlineSeconds {
		fmt.Println("Performance is OK transaction completed in", transactionTime)
	} else {
		fmt.Println("Performance problem, transaction completed in", transactionTime, "second(s)!!!")
	}
}
