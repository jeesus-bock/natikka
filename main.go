package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connect to NATS
	nc, err := nats.Connect("nats://cobolt:bluish@0.0.0.0:4111")
	if err != nil {
		log.Fatal(err)
	}
	// Create JetStream Context
	js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "ORDERS",
		Subjects: []string{"ORDERS.*"},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Simple Async Stream Publisher
	go func() {
		for {
			js.PublishAsync("ORDERS.scratch", []byte("taasuus "+time.Now().UTC().Format("15:04")))

			select {
			case <-js.PublishAsyncComplete():
				fmt.Println("Published")
			case <-time.After(5 * time.Second):
				fmt.Println("Did not resolve in time")
			}
			time.Sleep(1 * time.Second)
		}
	}()
	ctx := context.Background()
outer:
	for {
		select {
		case <-ctx.Done():
			break outer
		}
	}
}
