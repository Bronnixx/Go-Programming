package main

import (
	"fmt"
	"log"
	"os"
	"time"
)
func main(){
	message := []byte("imabronnix is the  text")	
	if err := os.WriteFile(createFile(), message, 0644); err!=nil{
		log.Panic(err)
	}
	fmt.Println("file has been created:", createFile())
	fmt.Println("Reading from the file in a moment")
	time.Sleep(5 * time.Second)
	fmt.Println(readFile(createFile()))
}
func createFile()string{
	fileName := os.Args[1]
	return fileName
  
}
func readFile(fileName string)string{
   content, err := os.ReadFile(fileName);
   if err != nil{
		log.Println(err)
	} 
	return string(content)
	
}