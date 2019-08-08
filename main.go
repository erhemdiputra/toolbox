package main

import (
	"log"
	"time"

	"github.com/erhemdiputra/toolbox/httprequest"
)

func main() {
	tryHTTPRequestPackage()
}

func printSeparator() {
	log.Println("==============================================================================")
}

func tryHTTPRequestPackage() {
	log.Println("[tryHTTPRequestPackage]")

	type Data struct {
		UserID    int64  `json:"userId"`
		ID        int64  `json:"id"`
		Title     string `json:"title"`
		Completed bool   `json:"completed"`
	}

	var data Data

	httpRequestInst := httprequest.NewHTTPRequest(5 * time.Second)
	err := httpRequestInst.DoRequest("GET", "https://jsonplaceholder.typicode.com/todos/1", &data)
	log.Printf("Data: %+v, Error: %+v\n", data, err)

	printSeparator()
}
