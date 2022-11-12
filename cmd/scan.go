package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/0l1v3rr/netscan/internal/network"
	"github.com/0l1v3rr/netscan/internal/utils"
	"github.com/spf13/cobra"
)

var netscan = &cobra.Command{
	Use:   "netscan",
	Short: "The netscan command scans the network you're in to find reachable hosts.",
	Run:   netscanRun,
}

func netscanRun(cmd *cobra.Command, args []string) {
	localIp := network.GetLocalIp()
	maskBits := network.GetMaskBits(localIp)
	networkAddr := network.NetworkAddress(localIp, maskBits)
	cidr := fmt.Sprintf("%v/%v", networkAddr, maskBits)

	hosts, err := network.Hosts(cidr)
	if err != nil {
		utils.Error("An unknown error occurred while scanning your network.")
	}

	fmt.Printf("Network: %v\n", cidr)
	fmt.Println("Scanning the network...")
	start := time.Now()
	reachableHosts := 0

	fmt.Println("\nReachable hosts in your network: ")

	network.ScanAll(hosts, func(available bool, host string) {
		if available {
			if host == localIp.String() {
				utils.Information(fmt.Sprintf("%s (You)", host))
			} else {
				utils.Information(host)
			}
			reachableHosts++
		}
	})

	elapsed := time.Since(start)
	fmt.Println("\nThe network has been scanned.")
	fmt.Printf("Hosts scanned: %v\n", len(hosts))
	fmt.Printf("Reachable hosts: %v\n", reachableHosts)
	fmt.Printf("Elapsed time: %v\n", elapsed)
}

func Execute() {
	port.Flags().StringVarP(&Host, "host", "o", "", "the host you want to scan (required)")
	port.Flags().StringVarP(&ToScan, "ports", "p", "", "ports you want to scan directly, separated by commas")
	port.Flags().BoolVarP(&ShowClosed, "closed", "c", false, "show the closed ports as well")
	port.MarkFlagRequired("host")

	netscan.AddCommand(port)

	if err := netscan.Execute(); err != nil {
		os.Exit(1)
	}
}
