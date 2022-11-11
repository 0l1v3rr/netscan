package network

import (
	"github.com/go-ping/ping"
)

func PingHost(hostname string) bool {
	pinger, err := ping.NewPinger(hostname)
	if err != nil {
		return false
	}

	pinger.SetPrivileged(true)
	pinger.Count = 3

	err = pinger.Run()
	if err != nil {
		return false
	}

	return pinger.Statistics().PacketsRecv > 1
}

func ScanAll(hosts []string, c func(available bool, host string)) {
	for _, h := range hosts {
		isOpen := PingHost(h)
		c(isOpen, h)
	}
}
