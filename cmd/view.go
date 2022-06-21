/*
Copyright © 2022 Jeroen Trimbach jeroentrimbach.com

*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// viewCmd represents the view command
var viewCmd = &cobra.Command{
	Use:   "view",
	Short: "A brief description of the view command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("%s\n\n", viper.AllKeys())
		settings := viper.AllSettings()
		for i, v := range settings {
			fmt.Printf("%v: %v\n", i, v)
		}
	},
}

func init() {
	configCmd.AddCommand(viewCmd)
}
