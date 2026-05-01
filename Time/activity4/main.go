package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	onlyAfter, err := time.Parse(time.RFC3339, "2020-5-1:19:00:00:00")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(now, onlyAfter)
	fmt.Println(now.After(onlyAfter))
	if now.After(onlyAfter) {
		fmt.Println("I'm DEADPOOL.....Executing actions!!!")
	} else {
		fmt.Println("Now is not the time yet jokerr!!!")
	}
}
