/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "A brief description of your command",
	Long:  `publish is testing grounds for NATS message publishing`,
	Run:   publish,
}

func init() {
	rootCmd.AddCommand(publishCmd)
}

func publish(cmd *cobra.Command, args []string) {
	// Connect to NATS
	//nc, err := nats.Connect("nats://test:test@0.0.0.0:4222")
	nc, err := nats.Connect(nats.DefaultURL)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to NATS")

	// Create JetStream Context
	js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))

	// Add the TEST_STREAM if it doesn't exist
	_, err = js.AddStream(&nats.StreamConfig{
		Name:     "TEST_STREAM",
		Subjects: []string{"TEST_STREAM.*"},
	})
	if err != nil {
		log.Fatal(err)
	}

	// Simple Async Stream Publisher
	go func() {
		for {
			// Sleep 1-5 seconds
			d := time.Duration(rand.Intn(5)) * time.Second
			time.Sleep(d)
			// Create test message.
			msg := "Time " + time.Now().UTC().Format("15:04:05")

			// And publish it.
			js.PublishAsync("TEST_STREAM.subj", []byte(msg))

			select {
			case <-js.PublishAsyncComplete():
				fmt.Println("Slept:", d.Seconds(), "Published: ", msg)
			case <-time.After(5 * time.Second):
				fmt.Println("Did not resolve in time")
			}
		}
	}()

	// Keep running without stressing the cpu.
	ctx := context.Background()
	<-ctx.Done()

}
