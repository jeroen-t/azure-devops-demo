/*
Copyright © 2022 Jeroen Trimbach jeroentrimbach.com

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "add 'Key-Value' pairs to the config file",
	Long: `add 'Key-Value' pairs to the config file. For example:

<cmd> config add --key firstname --value "jeroen"`,
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		value, _ := cmd.Flags().GetString("value")
		addKeyValuePair(key, value)
	},
}

func init() {
	configCmd.AddCommand(addCmd)
}

func addKeyValuePair(key string, value interface{}) {
	viper.Set(key, value)
	err := viper.WriteConfig()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Wrote the %s pair.\n", key)
}
