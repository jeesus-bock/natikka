/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"context"
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
	"github.com/spf13/cobra"
)

// consumeCmd represents the consume command
var consumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "Test JetStream consuming",
	Run: func(cmd *cobra.Command, args []string) {
		// Connect to NATS
		nc, _ := nats.Connect("nats://foo:bar@127.0.0.1:4222")

		// Create JetStream Context
		js, _ := nc.JetStream(nats.PublishAsyncMaxPending(256))

		// Simple Async Ephemeral Consumer
		_, err := js.Subscribe("TEST_STREAM.*", func(m *nats.Msg) {
			fmt.Printf("Received a JetStream message: %s - %s\n", string(m.Data), m.Subject)
			m.Ack()
		})
		if err != nil {
			log.Fatal(err)
		}

		// Keep running indefinitely
		ctx := context.Background()
		<-ctx.Done()
	},
}

func init() {
	rootCmd.AddCommand(consumeCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// consumeCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// consumeCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
