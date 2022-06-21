/*
Copyright © 2022 Jeroen Trimbach jeroentrimbach.com

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// updateCmd represents the update command
var updateCmd = &cobra.Command{
	Use:   "update",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		value, _ := cmd.Flags().GetString("value")
		updateKeyValuePair(key, value)
	},
}

func init() {
	configCmd.AddCommand(updateCmd)
}

func updateKeyValuePair(key string, value interface{}) {
	viper.Set(key, value)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Wrote the %s pair.\n", key)
}
