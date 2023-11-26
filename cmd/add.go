/*
Copyright © 2023 SantiSite santiagorv246@santisite.com
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/SantiSite/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func addRun(cmd *cobra.Command, args []string) {
	var (
		err   error
		items []todo.Item
	)
	items, err = todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v", err)
	}
	for _, arg := range args {
		item := todo.Item{Text: arg}
		item.SetPriority(priority)
		items = append(items, item)
	}
	err = todo.SaveItems(viper.GetString("datafile"), items)
	if err != nil {
		_ = fmt.Errorf("%v", err)
	}
}

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add a new todo",
	Long:  `Add will create a new todo item to the list.`,
	Run:   addRun,
}

func init() {
	rootCmd.AddCommand(addCmd)

	addCmd.Flags().IntVarP(&priority, "priority", "p", 2, "Priority:1,2,3")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
