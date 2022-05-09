/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"tracker/types"

	"github.com/spf13/cobra"
)

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		v, _ := handler.Handler(types.MESSAGE_TRANSACTION_KEEPER_GET_ALL_TRANSACTIONS, nil)
		store := v.(map[string]*types.Transaction)
		printTransactions(store)
	},
}

func printTransactions(t map[string]*types.Transaction) {
	for id, v := range t {
		fmt.Printf("%s -> Coin: %s, Value: %f \n", id, v.Coin, v.Value)
	}
}

func init() {
	transactionCmd.AddCommand(listCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// listCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// listCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
