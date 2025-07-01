package main

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"time"
)

func main() {
	fmt.Println("Golang eureka client")
	// Listening for avaiable ports on the machine.
	listener, err := net.Listen("tcp", "0.0.0.0:0")
	if err != nil {
		log.Fatal(err)
	}

	addr := listener.Addr().(*net.TCPAddr)
	host := addr.IP.String()
	port := addr.Port

	// Running the registerService in a goroutine so that it didnt block the main thread.
	go func() {
		time.Sleep(4*time.Second)
		fmt.Println("Making request to the eureka server")
		registerService("go-service", addr.String(), host, port)
	}()


	http.Serve(listener, nil)
}
