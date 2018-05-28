package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var file string
var N int

// RootCmd is the basic command for Aliens
var RootCmd = &cobra.Command{
	Use:   "aliens",
	Short: "Run simulation of a battle of aliens",
	Run: func(cmd *cobra.Command, args []string) {
		err := Init(file, N)
		if err != nil {
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
	RootCmd.AddCommand(testCmd)
	RootCmd.PersistentFlags().IntVarP(&N, "N", "N", 10, "Number of aliens placed in the map")
	RootCmd.Flags().StringVarP(&file, "file", "f", "example.txt", "Full path to the .txt file containing the map")
	RootCmd.MarkFlagRequired("file")
	RootCmd.MarkFlagRequired("N")
	testCmd.MarkFlagRequired("N")
	viper.BindPFlag("file", RootCmd.Flags().Lookup("file"))
	viper.BindPFlag("N", RootCmd.Flags().Lookup("N"))
	viper.BindPFlag("N", testCmd.Flags().Lookup("N"))
}
