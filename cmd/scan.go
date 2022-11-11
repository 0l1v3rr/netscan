package cmd

import (
	"fmt"
	"os"

	"github.com/0l1v3rr/netscan/internal/network"
	"github.com/spf13/cobra"
)

var DialTime int

var netscan = &cobra.Command{
	Use:   "netscan",
	Short: "The netscan command scans the network you're in.",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	localIp := network.GetLocalIp()
	maskBits := network.GetMaskBits(localIp)
	networkAddr := network.NetworkAddress(localIp, maskBits)
	cidr := fmt.Sprintf("%v/%v", networkAddr, maskBits)

	fmt.Printf("Network: %v\n", cidr)
	fmt.Println("Scanning the network...")
}

func Execute() {
	netscan.Flags().IntVarP(&DialTime, "dialtime", "t", 5, "The dial timeout")

	if err := netscan.Execute(); err != nil {
		os.Exit(1)
	}
}
