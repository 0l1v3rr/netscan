package network

import (
	"fmt"
	"log"
	"net"
)

func GetLocalIp() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func GetMaskBits(ip net.IP) int {
	counter := 0

	for _, ch := range ip.DefaultMask().String() {
		if ch == 'f' {
			counter++
		}
	}

	return counter * 4
}

func NetworkAddress(ip net.IP, bits int) string {
	ip, ipnet, _ := net.ParseCIDR(fmt.Sprintf("%v/%v", ip.String(), bits))
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		return ip.String()
	}
	return ""
}

func Hosts(cidr string) ([]string, error) {
	ip, ipnet, err := net.ParseCIDR(cidr)
	if err != nil {
		return nil, err
	}

	var ips []string
	for ip := ip.Mask(ipnet.Mask); ipnet.Contains(ip); inc(ip) {
		ips = append(ips, ip.String())
	}

	// remove network and broadcast address
	return ips[1 : len(ips)-1], nil
}

func inc(ip net.IP) {
	for j := len(ip) - 1; j >= 0; j-- {
		ip[j]++
		if ip[j] > 0 {
			break
		}
	}
}
