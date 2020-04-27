package cmd

import "github.com/spf13/cobra"

var (
	RootCmd = &cobra.Command{
		Use: "mygb",
		Short: "Run mygb",
		Long: "Run mygb",
		SilenceUsage: true,
	}
)

