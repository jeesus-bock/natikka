/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"log"
	"math/rand"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

// publishCmd represents the publish command
var publishCmd = &cobra.Command{
	Use:   "publish",
	Short: "Test JetStream publishing",
	Run:   publish,
}

func init() {
	rootCmd.AddCommand(publishCmd)
}

func publish(cmd *cobra.Command, args []string) {
	// Connect to NATS
	nc, err := nats.Connect("nats://foo:bar@127.0.0.1:4222")

	///nc, err := nats.Connect(nats.DefaultURL,)

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

	// Simple Stream Publisher
	go func() {
		for {
			// Sleep 1-5 seconds
			d := time.Duration(rand.Intn(5)) * time.Second
			time.Sleep(d)
			// Create test message
			msg := "Time " + time.Now().Format("15:04:05")

			// And publish it
			js.Publish("TEST_STREAM.subj", []byte(msg))
			log.Println("Published: " + msg)
		}
	}()

	// Keep running without stressing the cpu.
	ctx := context.Background()
	<-ctx.Done()

}
