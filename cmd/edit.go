/*
Copyright © 2023 SantiSite santiagorv246@santisite.com
*/
package cmd

import (
	"errors"
	"fmt"
	"github.com/SantiSite/tri/todo"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"strconv"
)

func editRun(cmd *cobra.Command, args []string) {
	var (
		err   error
		items []todo.Item
		idx   int
	)
	items, err = todo.ReadItems(viper.GetString("datafile"))
	if err != nil {
		log.Printf("%v\n", err)
	}
	if len(args) == 1 {
		for i := 0; i < len(args); i++ {
			idx, _ = strconv.Atoi(args[i])
		}
	}
	var item = &items[idx-1]
	var newTitle = args[1]
	fmt.Println("newTitle:", newTitle)
	if newTitle != "" {
		item.Text = newTitle
	}
	if priority > 0 {
		item.Priority = priority
	}
	err = todo.SaveItems(viper.GetString("datafile"), items)
	if err != nil {
		log.Printf("%v\n", err)
	}
}

// editCmd represents the edit command
var editCmd = &cobra.Command{
	Use:   "edit id [id | text]",
	Short: "Edit the content of an item",
	Long:  `Edit will update the content of an existing item in the list.`,
	Example: `> tri edit 1 "New content"
> tri edit 1 -p2
> tri edit 2 --due 05/13/15
> tri edit 1 2 3 -p1
> tri edit 2 3 --due 05/12/15`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("requires at least one arg")
		}
		if _, err := strconv.Atoi(args[0]); err != nil {
			return errors.New("el primer argumento debe de ser un entero, el cual hace referencia al id de la tarea")
		}
		return nil
	},
	Run: editRun,
}

func init() {
	rootCmd.AddCommand(editCmd)

	editCmd.Flags().IntVarP(&priority, "priority", "p", 0, "Priority:1,2,3")
	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// editCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// editCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
