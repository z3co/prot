package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

var completeCmd = &cobra.Command{
	Use: "complete [id]",
	Short: "Complete a todo with its id",
	Args: cobra.MinimumNArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		id, err := strconv.Atoi(args[0])
		if err != nil {
			return fmt.Errorf("invalid id, not a number")
		}
		todo, err := Instance.GetTodoById(cmd.Context(), int64(id))
		if err != nil { return fmt.Errorf("could not get todo info") }
		if todo.Done == 1 {
			err := Instance.UnCompleteTodo(cmd.Context(), int64(id))
			if err != nil { return fmt.Errorf("could not mark todo as uncompleted") }
			fmt.Println("Marked todo as uncompleted")
		} else {
			err := Instance.CompleteTodo(cmd.Context(), int64(id))
			if err != nil { return fmt.Errorf("could not mark todo completed") }
			fmt.Println("Marked todo as completed")
		}
		return nil
	},
}

func init() {
	rootCmd.AddCommand(completeCmd)
}
