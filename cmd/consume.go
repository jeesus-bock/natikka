/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// consumeCmd represents the consume command
var consumeCmd = &cobra.Command{
	Use:   "consume",
	Short: "A brief description of your command",
	Long:  `consume is for testing message consumers`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("consume called")
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
