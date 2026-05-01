package main

import (
	"fmt"
	"time"
)

func main() {
	now := time.Now()
	onlyAafter, err := time.Parse(time.RFC3339, "2026-5-1:19:41+00:00")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(now, onlyAafter)
	fmt.Println(now.After(onlyAafter))
	if now.After(onlyAafter) {
		fmt.Println("I'm DEADPOOL.....Executing actions!!!")
	} else {
		fmt.Println("Now is not the time yet jokerr!!!")
	}
}
