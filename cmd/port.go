package cmd

import (
	"fmt"

	"github.com/0l1v3rr/netscan/internal/network"
	"github.com/spf13/cobra"
)

var (
	Host       string
	ShowClosed bool
)

var port = &cobra.Command{
	Use:   "port",
	Short: "The port command searches for open ports..",
	Run:   portRun,
}

func portRun(cmd *cobra.Command, args []string) {
	reachable := network.HostReachable(Host)
	if !reachable {
		fmt.Printf("Error: host \"%v\" is not reachable\n", Host)
		return
	}

	fmt.Println("Port command")
	fmt.Println(Host)
	fmt.Println(ShowClosed)
}
