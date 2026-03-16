package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	db "github.com/z3co/prot/db/gen"
)

var addCmd = &cobra.Command{
	Use:   "add [todo]",
	Short: "Add a new todo",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		description := strings.Join(args, " ")
		folder, err := os.Getwd()
		if err != nil {
			return err
		}
		_, err = Instance.CreateTodo(cmd.Context(), db.CreateTodoParams{Description: description, Folder: folder})
		if err != nil {
			return fmt.Errorf("could not create todo: %s", err)
		}
		fmt.Println("Succesfully created your todo")
		return nil
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
