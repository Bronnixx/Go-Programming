package main

import (
	"fmt"
	"time"
)

func main() {
	start := time.Now()
	fmt.Println("The script has started at: ", start)
	fmt.Println("I'm BATMAN...")
	time.Sleep(2 * time.Second)
	end := time.Now()
	fmt.Println("The script has completed at: ", end)
}
