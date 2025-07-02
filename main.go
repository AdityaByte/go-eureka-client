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
	serviceName := "go-service"

	// Running the registerService in a goroutine so that it didnt block the main thread.
	go func() {
		time.Sleep(4*time.Second)
		fmt.Println("Making request to the eureka server")
		registerService(fmt.Sprintf("localhost:%s:%d", serviceName, port), serviceName, addr.String(), host, port)
	}()

	ticker := time.NewTicker(30*time.Second)

	done := make(chan bool)
	go func() {
		// Delaying this goroutine for 6 seconds
		time.Sleep(6*time.Second)
		for {
			select {
			case <- ticker.C :
				err := sendHeartBeats("go-service", fmt.Sprintf("localhost:go-service:%d", port))
				if err != nil {
					ticker.Stop()
					log.Fatal(err)
					return
				}
			case <- done:
				ticker.Stop();
				return
			}
		}
	} ()


	http.Serve(listener, nil)
}
