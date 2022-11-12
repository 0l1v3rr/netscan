package utils

import (
	"fmt"

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
