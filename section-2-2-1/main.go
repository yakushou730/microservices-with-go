package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// client
func main() {
	resp, err := http.Get("http://localhost:8080/example")
	if err != nil {
		log.Fatal(err.Error())
	}
	statusCode := resp.StatusCode
	if statusCode > 200 {
		log.Fatalf("Error or Request, StatusCode: %d", statusCode)
	}
	contentType := resp.Header.Get("Content-type")
	fmt.Println("Content-type: " + contentType)

	proto := resp.Proto
	fmt.Println("Proto: " + proto)

	contLen := resp.ContentLength
	fmt.Printf("Content Length: %d\n", contLen)

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Error Reading Body")
	}
	defer resp.Body.Close()

	fmt.Println("Client1:: Received::")
	fmt.Println(string(body))
}
