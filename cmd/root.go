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
	rootCmd.AddCommand(testCmd)
	rootCmd.PersistentFlags().IntVarP(&N, "N", "N", 10, "Number of aliens placed in the map")
	rootCmd.Flags().StringVarP(&file, "file", "f", "example.txt", "Full path to the .txt file containing the map")
	rootCmd.MarkFlagRequired("file")
	rootCmd.MarkFlagRequired("N")
	testCmd.MarkFlagRequired("N")
	viper.BindPFlag("file", rootCmd.Flags().Lookup("file"))
	viper.BindPFlag("N", rootCmd.Flags().Lookup("N"))
	viper.BindPFlag("N", testCmd.Flags().Lookup("N"))
}
