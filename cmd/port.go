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
	ToScan     string
	Dialtime   int
)

var port = &cobra.Command{
	Use:   "port",
	Short: "Searches for open ports",
	Run:   portRun,
}

func portRun(cmd *cobra.Command, args []string) {
	var ports []int
	if ToScan != "" {
		var err error
		ports, err = network.ParsePortString(ToScan)

		if err != nil {
			fmt.Println("Error: please provide valid ports. Example -p usage: 21,23,80,443")
			return
		}
	} else {
		ports = network.Ports
	}

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

	network.ScanPorts(ports, Host, Dialtime, func(port int, isOpen bool) {
		if isOpen {
			utils.Port(port, isOpen)
			openPorts++
			return
		}

		if ShowClosed {
			utils.Port(port, isOpen)
		}
	})

	elapsed := time.Since(start)
	fmt.Println("\nThe host has been scanned.")
	fmt.Printf("Ports checked: %v\n", len(ports))
	fmt.Printf("Open  ports: %v\n", openPorts)
	fmt.Printf("Elapsed time: %v\n", elapsed)
}
