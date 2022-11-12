package utils

import (
	"fmt"

	"github.com/0l1v3rr/netscan/internal/network"
	"github.com/fatih/color"
)

func Information(message string) {
	c := color.New(color.FgHiBlue)
	c.Print("[*]")
	fmt.Printf(" %s\n", message)
}

func Success(message string) {
	c := color.New(color.FgHiGreen)
	c.Print("[$]")
	fmt.Printf(" %s\n", message)
}

func Error(message string) {
	c := color.New(color.FgHiRed)
	c.Print("[#]")
	fmt.Printf(" %s\n", message)
}

func Port(port int, open bool) {
	if open {
		Information(fmt.Sprintf("%v \tOPEN \t%v", port, network.PortService(port)))
		return
	}

	Error(fmt.Sprintf("%v \tCLOSED \t%v", port, network.PortService(port)))
}
