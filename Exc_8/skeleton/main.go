package main

import (
	"exc8/client"
	"exc8/server"
	"log"
	"time"
)

func main() {
	go func() {
		// todo start server
		log.Println("Starting gRPC server...")
		if err := server.StartGrpcServer(); err != nil {
			log.Fatalf("Server error: %v", err)
		}

	}()
	time.Sleep(1 * time.Second)
	// todo start client
	log.Println("Starting client...")
	c, err := client.NewGrpcClient()
	if err != nil {
		log.Fatalf("Failed to create client: %v", err)
	}

	if err := c.Run(); err != nil {
		log.Fatalf("Client error: %v", err)
	}

	println("Orders complete!")
}
