package main

import (
	"log"
	"strconv"
	"time"
)

func main() {
	log.SetFlags(log.Ldate | log.Llongfile)
	appName := "HTTPCHECKER"
	action := "BASIC"
	date := time.Now()
	logFileName := appName + "_" + action + "_" + strconv.Itoa(date.Year()) + "_" + date.Month().String() + "_" + strconv.Itoa(date.Day()) + ".log"
	log.Println("The name of the logfile is", logFileName)
}
