/*
Copyright © 2022 Jeroen Trimbach jeroentrimbach.com

*/
package cmd

import (
	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "configuration management",
	Long:  "configuration management",
	// Run: func(cmd *cobra.Command, args []string) {
	// 	fmt.Println("config called")
	// },
}

func init() {
	rootCmd.AddCommand(configCmd)

	configCmd.PersistentFlags().StringP("key", "k", "", "The key for the key value set to add to the configuration.")
	configCmd.PersistentFlags().StringP("value", "v", "", "The value for the key value set to add to the configuration.")
}
