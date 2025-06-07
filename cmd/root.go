package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var (
	region  string
	profile string
)

var rootCmd = &cobra.Command{
	Use:   "sm",
	Short: "CLI tool for managing AWS Session Manager sessions.",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&region, "region", "us-east-1", "AWS region")
	rootCmd.PersistentFlags().StringVar(&profile, "profile", "default", "AWS CLI profile")
	rootCmd.AddCommand(listAndConnectCmd())
}
