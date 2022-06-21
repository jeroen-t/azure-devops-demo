/*
Copyright © 2022 Jeroen Trimbach jeroentrimbach.com

*/
package cmd

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of the delete command",
	Long: `A longer description for delete that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		key, _ := cmd.Flags().GetString("key")
		err := deleteKeyValuePair(key)
		if err != nil {
			fmt.Println(err)
		}
	},
}

func init() {
	configCmd.AddCommand(deleteCmd)
}

func deleteKeyValuePair(key string) error {
	configMap := viper.AllSettings()
	delete(configMap, key)
	encodedConfig, _ := json.MarshalIndent(configMap, "", " ")
	err := viper.ReadConfig(bytes.NewReader(encodedConfig))
	if err != nil {
		return err
	}
	return viper.WriteConfig()
}
