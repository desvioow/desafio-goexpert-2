package cmd

import (
	"desafio-goexpert-2/internal/validation"
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "desafio-goexpert-2",
	Short: "A simple stress test cli tool",
	Long: `A simple stress test cli tool that can be used to test the performance of a web service and generate a report.
It achieves this by sending a number of requests to the web service and measuring the time it takes to receive a response.`,

	RunE: func(cmd *cobra.Command, args []string) error {

		url, err := cmd.Flags().GetString("url")
		if err != nil {
			return fmt.Errorf("failed to get URL flag: %w", err)
		}

		requests, err := cmd.Flags().GetInt("requests")
		if err != nil {
			return fmt.Errorf("failed to get requests flag: %w", err)
		}

		concurrency, err := cmd.Flags().GetInt("concurrency")
		if err != nil {
			return fmt.Errorf("failed to get concurrency flag: %w", err)
		}

		if err := validation.ValidateFlags(url, requests, concurrency); err != nil {
			return err
		}

		fmt.Printf("Testing URL: %s\n", url)
		fmt.Printf("Requests: %d\n", requests)
		fmt.Printf("Concurrency: %d\n", concurrency)

		return nil
	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.desafio-goexpert-2.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().StringP("url", "u", "", "URL to test")
	rootCmd.MarkFlagRequired("url")

	rootCmd.Flags().IntP("requests", "r", 0, "Amount of requests to send")
	rootCmd.MarkFlagRequired("requests")

	rootCmd.Flags().IntP("concurrency", "c", 0, "Concurrency level")
	rootCmd.MarkFlagRequired("concurrency")
}
