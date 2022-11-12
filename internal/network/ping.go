package network

import (
	"github.com/go-ping/ping"
)

func PingHost(hostname string) (*ping.Statistics, error) {
	pinger, err := ping.NewPinger(hostname)
	if err != nil {
		return nil, err
	}

	pinger.SetPrivileged(true)
	pinger.Count = 3

	err = pinger.Run()
	if err != nil {
		return nil, err
	}

	return pinger.Statistics(), nil
}

func HostReachable(hostname string) bool {
	res, err := PingHost(hostname)
	if err != nil {
		return false
	}

	return res.PacketsRecv > 1
}

func ScanAll(hosts []string, c func(available bool, host string)) {
	for _, h := range hosts {
		c(HostReachable(h), h)
	}
}
