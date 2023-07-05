package commandline

import (
	"pius/databases"
	"pius/libs"

	"github.com/spf13/cobra"
)

var InitCommand = cobra.Command{
	Short: "on backend",
	Long: "online notes backend",
}

func init()  {
	InitCommand.AddCommand(libs.ServeCMD)
	InitCommand.AddCommand(databases.MigrateCMD)
}

func Run(args []string) error {
	InitCommand.SetArgs(args)

	return InitCommand.Execute()
}