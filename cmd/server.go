package cmd

import "github.com/spf13/cobra"

var (
	serverCmd = &cobra.Command{
		Use:		"server",
		Aliases: 	[]string{"s"},
		Short:		"run server",
		Long:		"Run Game Boy games on server",
		Run:		nil,
	}
)

func init() {

	RootCmd.AddCommand(serverCmd)
}