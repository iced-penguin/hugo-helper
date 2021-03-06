package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var rootCmd = &cobra.Command{
	Use:   "hugo-helper",
	Short: "cli tool for blogs generated by HUGO",
	Long:  "cli tool for blogs generated by HUGO",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("hello helper")
	},
}

var toggle bool

func init() {
	cobra.OnInitialize(initConfig)
}

func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func initConfig() {
	viper.SetConfigName("hugo-helper")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")
	if err := viper.ReadInConfig(); err != nil {
		panic(fmt.Errorf("failed to read config file: %w", err))
	}
}
