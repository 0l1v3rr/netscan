package cmd

import (
	"fmt"
	"time"

	"github.com/go-ping/ping"
	"github.com/spf13/cobra"
)

var cPing = &cobra.Command{
	Use:   "ping",
	Short: "The ping command acts just like a UNIX ping command.",
	Run:   pingRun,
}

var (
	PCount    int
	PTimeout  time.Duration
	PInterval time.Duration
	PSize     int
	PTtl      int
)

func pingRun(cmd *cobra.Command, args []string) {
	if len(args) < 1 {
		fmt.Println("Error: Please provide a host to ping.")
		return
	}

	pinger, err := ping.NewPinger(args[0])
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	pinger.OnRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v\n", pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl)
	}
	pinger.OnDuplicateRecv = func(pkt *ping.Packet) {
		fmt.Printf("%d bytes from %s: icmp_seq=%d time=%v ttl=%v (DUP!)\n", pkt.Nbytes, pkt.IPAddr, pkt.Seq, pkt.Rtt, pkt.Ttl)
	}
	pinger.OnFinish = func(stats *ping.Statistics) {
		fmt.Printf("\n--- %s ping statistics ---\n", stats.Addr)
		fmt.Printf("%d packets transmitted, %d packets received, %d duplicates, %v%% packet loss\n", stats.PacketsSent, stats.PacketsRecv, stats.PacketsRecvDuplicates, stats.PacketLoss)
		fmt.Printf("round-trip min/avg/max/stddev = %v/%v/%v/%v\n", stats.MinRtt, stats.AvgRtt, stats.MaxRtt, stats.StdDevRtt)
	}

	pinger.Count = PCount
	pinger.Timeout = PTimeout
	pinger.Interval = PInterval
	pinger.Size = PSize
	pinger.TTL = PTtl
	pinger.SetPrivileged(true)

	fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())

	err = pinger.Run()
	if err != nil {
		fmt.Println("Failed to ping target host:", err)
	}
}
