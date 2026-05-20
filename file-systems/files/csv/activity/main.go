package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"strings"
)
func main(){
	in := `firstname, lastname, age
	danny, jones, 12 
	erick, ndeto, 21
	jayden, mutiri, 12 
	emily, maitha, 16
       			`
          r := csv.NewReader(strings.NewReader(in))
          for  {
          record, err := r.Read()
          if  err == io.EOF{
          break
          }
          if err != nil{
          log.Fatal(err)
          }
          fmt.Println(record)
          }
}