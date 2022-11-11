package cmd

import (
	"fmt"
	"os"

	"github.com/0l1v3rr/netscan/internal/network"
	"github.com/0l1v3rr/netscan/internal/utils"
	"github.com/spf13/cobra"
)

var netscan = &cobra.Command{
	Use:   "netscan",
	Short: "The netscan command scans the network you're in to find reachable hosts.",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
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
	fmt.Println("\nReachable hosts in your network: ")

	network.ScanAll(hosts, func(available bool, host string) {
		if available {
			utils.Information(host)
		}
	})
}

func Execute() {
	if err := netscan.Execute(); err != nil {
		os.Exit(1)
	}
}
