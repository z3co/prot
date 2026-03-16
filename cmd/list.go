package cmd

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/spf13/cobra"
	db "github.com/z3co/prot/db/gen"
)


var all bool

var listCmd = &cobra.Command{
	Use: "list",
	Short: "It lists uncompleted todos, unless all is passed then all todos are listed",
	RunE: func(cmd *cobra.Command, args []string) error {
		var todos []db.Todo
		todos, err := Instance.GetTodosByFolder(cmd.Context(), CurrentDir)
		if err != nil {
			return fmt.Errorf("Error while getting all todo: %s", err)
		}
		w := tabwriter.NewWriter(os.Stdout, 0, 0 ,4, ' ', 0)
		fmt.Fprintln(w, "ID\tDescription\tDone")
		for _, row := range todos {
			fmt.Fprintf(w, "%v\t%s\t%v\n", row.ID, row.Description, row.Done == 1)
		}
		w.Flush()
		return nil
	},
}

func init() {
	listCmd.Flags().BoolVarP(&all, "all", "a", false, "decides wether all todos or just the uncompleted is shown")
	rootCmd.AddCommand(listCmd)
}
