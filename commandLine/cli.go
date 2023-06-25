package commandline

import (
	"pius/libs"

	"github.com/spf13/cobra"
)

var InitCommand = cobra.Command{
	Short: "on backend",
	Long: "online notes backend",
}

func init()  {
	InitCommand.AddCommand(libs.ServeCMD)
}

func Run(args []string) error {
	InitCommand.SetArgs(args)

	return InitCommand.Execute()
}