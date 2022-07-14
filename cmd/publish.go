/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"log"
	"math/rand"
	"time"

	"natikka/structs"

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
	if err != nil {
		log.Fatal(err)
	}
	enc, err := nats.NewEncodedConn(nc, nats.JSON_ENCODER)
	if err != nil {
		log.Fatal(err)
	}
	///nc, err := nats.Connect(nats.DefaultURL,)

	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to NATS")

	// Simple Encoded Publisher
	go func() {
		for {
			// Sleep 1-5 seconds
			d := time.Duration(rand.Intn(5)) * time.Second
			time.Sleep(d)
			// Create test message
			msg := structs.Data{Time: time.Now(), Msg: "Foo Bar Baz"}

			// And publish it
			enc.Publish("subj", msg)
			log.Printf("Published: %+v\n", msg)
		}
	}()

	// Keep running without stressing the cpu.
	ctx := context.Background()
	<-ctx.Done()

}
