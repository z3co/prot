package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete [id]",
	Short: "Deletes a todo by id",
	Args:  cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("id is not an integer")
		}
		err = Instance.DeleteTodo(cmd.Context(), int64(id))
		if err != nil {
			return fmt.Errorf("could not delete todo: %s", err)
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)
}
