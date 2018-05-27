package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var file string
var N int

var rootCmd = &cobra.Command{
	Use:   "aliens",
	Short: "Run simulation of a battle of aliens",
	Run: func(cmd *cobra.Command, args []string) {
		err = Init(file, N)
		if error != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	},
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run tests for the alien task",
	Run: func(cmd *cobra.Command, args []string) {
		// TODO
		// run tests N
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(testCmd)
	rootCmd.PersistentFlags().StringVarP(&N, "N", "N", "", "Number of aliens placed in the map")
	runCmd.Flags().StringVarP(&file, "file", "f", "example.txt", "Full path to the .txt file containing the map")
	runCmd.MarkFlagRequired("file")
	runCmd.MarkFlagRequired("N")
	viper.BindPFlag("file", runCmd.Flags().Lookup("file"))
	viper.BindPFlag("N", runCmd.Flags().Lookup("N"))
}
