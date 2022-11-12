package cmd

import (
	"fmt"
	"time"

	"github.com/0l1v3rr/netscan/internal/network"
	"github.com/0l1v3rr/netscan/internal/utils"
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

	fmt.Printf("Host: %v\n", Host)
	fmt.Println("Scanning for open ports...")
	fmt.Println()
	start := time.Now()
	openPorts := 0

	network.ScanMostKnownPorts(Host, 5, func(port int, isOpen bool) {
		if isOpen {
			utils.Information(fmt.Sprintf("%v \tOPEN \t%v", port, network.PortService(port)))
			openPorts++
			return
		}

		if ShowClosed {
			utils.Error(fmt.Sprintf("%v \tCLOSED \t%v", port, network.PortService(port)))
		}
	})

	elapsed := time.Since(start)
	fmt.Println("\nThe host has been scanned.")
	fmt.Printf("Ports checked: %v\n", len(network.Ports))
	fmt.Printf("Open  ports: %v\n", openPorts)
	fmt.Printf("Elapsed time: %v\n", elapsed)
}
