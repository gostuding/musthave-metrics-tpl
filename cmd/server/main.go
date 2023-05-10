package main

import (
	"flag"
	"fmt"
	"net/http"
	"os"
)

// начало работы
func main() {

	ip_address := flag.String("a", ":8080", "address and port to run server like address:port")
	flag.Parse()
	//-------------------------------------------------------------------------
	if address := os.Getenv("ADDRESS"); address != "" {
		ip_address = &address
	}
	//-------------------------------------------------------------------------
	fmt.Println("Run server at adress: ", *ip_address)
	err := http.ListenAndServe(*ip_address, GetRouter())
	if err != nil {
		panic(err)
	}
}
