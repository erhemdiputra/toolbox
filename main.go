package main

import (
	"log"
	"net/http"
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
	httpRequestInst := httprequest.NewHTTPRequest(5 * time.Second)

	// GET
	getResp := make(map[string]interface{})

	err := httpRequestInst.Get("https://jsonplaceholder.typicode.com/todos/1", http.Header{}, map[string]string{
		"hello":    "world",
		"kenpachi": "zaraki",
	}, &getResp)
	log.Printf("[GET] Data: %+v, Error: %+v\n", getResp, err)
	// EOF GET

	// POST
	postResp := make(map[string]interface{})

	err = httpRequestInst.Post("https://jsonplaceholder.typicode.com/posts",
		http.Header{}, map[string]string{
			"shinigami": "ichigo kurosaki, byakuya kuchiki, renji abarai!",
		}, nil, &postResp,
	)
	log.Printf("[POST][Form URL Encoded] Data: %+v, Error: %+v\n", postResp, err)

	err = httpRequestInst.Post("https://jsonplaceholder.typicode.com/posts",
		http.Header{}, nil, map[string]interface{}{
			"title":  "foo",
			"body":   "bar",
			"userId": 1,
		}, &postResp,
	)
	log.Printf("[POST][JSON] Data: %+v, Error: %+v\n", postResp, err)
	// EOF POST

	printSeparator()
}
