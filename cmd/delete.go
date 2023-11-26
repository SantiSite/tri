/*
Copyright © 2023 SantiSite santiagorv246@santisite.com
*/
package cmd

import (
	"fmt"
	"github.com/SantiSite/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"sort"
	"strconv"
)

func deleteRun(cmd *cobra.Command, args []string) {
	items, err := todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v\n", err)
	} else {
		if len(items) == 0 {
			log.Println("No items to delete")
		} else {
			var itemsToBeDeleted []int
			for _, arg := range args {
				argInInt, err := strconv.Atoi(arg)
				if err != nil {
					log.Printf("%v\n", err)
				} else {
					itemsToBeDeleted = append(itemsToBeDeleted, argInInt)
				}
			}

			if len(itemsToBeDeleted) > 0 {
				mp := make(map[int]bool)
				var argNoDup []int
				for _, arg := range itemsToBeDeleted {
					if _, value := mp[arg]; !value {
						mp[arg] = true
						argNoDup = append(argNoDup, arg)
					}
				}

				sort.Sort(todo.ByNumber(argNoDup))

				var deletedItems []int
				for _, arg := range argNoDup {
					if len(items) < arg {
						log.Println("Invalid item number. " + strconv.Itoa(arg) + " is not in the list")
						continue
					}
					deletedItems = append(deletedItems, arg)
					items = append(items[:arg-1], items[arg:]...)
				}
				err = todo.SaveItems(viper.GetString("datafile"), items)
				if err != nil {
					log.Printf("%v\n", err)
				} else {
					fmt.Printf("Items with id/s %v have been removed from the list\n", deletedItems)
				}
			}
		}
	}
}

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "Remove a todo item",
	Long:  `Remove a todo item.`,
	Run:   deleteRun,
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
