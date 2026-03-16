package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	_ "embed"

	"github.com/spf13/cobra"
	db "github.com/z3co/prot/db/gen"
	logic "github.com/z3co/prot/db/sql"
	_ "modernc.org/sqlite"
)

var dbPath string
var Instance *db.Queries
var CurrentDir string

var rootCmd = &cobra.Command{
	Use: "prot",
	Short: "A project based todolist",
	Long: `Prot is a project based todolist. 
	That means that each todolist is connected to the folder of the project it was created in.`,
	PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
		var err error
		CurrentDir, err = os.Getwd()
		if err != nil { return err }
		dbConn, err := logic.InitDB(dbPath)
		if err != nil { return err }
		Instance = db.New(dbConn)
		return nil
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	configDir, _ := os.UserConfigDir()
	appConfigDir := filepath.Join(configDir, "prot")
	defaultPath := filepath.Join(appConfigDir, "databse.sqlite")
	rootCmd.PersistentFlags().StringVarP(&dbPath, "database", "d", defaultPath, "database file location")
}
