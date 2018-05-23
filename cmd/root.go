package main

import (
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "aliens",
	Short: "Simulation of a battle of aliens",
	Long:  "Aliens is a simulation of a battle of aliens built by @fedekunze and written in Go",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run main programm of the alien task",
	Long:  "Run main programm of the Cosmos' alien coding task",
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Run tests for the alien task",
	Long:  "Run logic and types tests for the Cosmos' alien coding task",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
	},
}

func init() {
	cobra.OnInitialize(initConfig)
	// rootCmd.Flags().StringVarP(&file, "file", "f", "", "Full path to the .txt file containing the map")
	rootCmd.MarkFlagRequired("file")
	viper.BindPFlag("file", rootCmd.PersistentFlags().Lookup("file"))
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(testCmd)
}

func initConfig() {
	// // Don't forget to read config either from cfgFile or from home directory!
	// if cfgFile != "" {
	// 	// Use config file from the flag.
	// 	viper.SetConfigFile(cfgFile)
	// } else {
	// 	// Find home directory.
	// 	home, err := homedir.Dir()
	// 	if err != nil {
	// 		fmt.Println(err)
	// 		os.Exit(1)
	// 	}
	//
	// 	// Search config in home directory with name ".cobra" (without extension).
	// 	viper.AddConfigPath(home)
	// 	viper.SetConfigName(".cobra")
	// }
	//
	// if err := viper.ReadInConfig(); err != nil {
	// 	fmt.Println("Can't read config:", err)
	// 	os.Exit(1)
	// }
}
