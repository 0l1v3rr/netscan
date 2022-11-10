package cmd

import (
	"os"

	"github.com/spf13/cobra"
)

var DialTime int

var netscan = &cobra.Command{
	Use:   "netscan",
	Short: "The netscan command scans the network you're in.",
	Run: func(cmd *cobra.Command, args []string) {

	},
}

func Execute() {
	netscan.Flags().IntVarP(&DialTime, "dialtime", "t", 5, "The dial timeout")

	if err := netscan.Execute(); err != nil {
		os.Exit(1)
	}
}
