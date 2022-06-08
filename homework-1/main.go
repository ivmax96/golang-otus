package main

import (
	"fmt"
	"os"

	"github.com/beevik/ntp"
)

const host = "0.beevik-ntp.pool.ntp.org"

func main() {
	curTime, err := ntp.Time(host)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	fmt.Printf("Current time: %s\n", curTime.Format("15:04:05"))
	fmt.Printf("Current date: %s\n", curTime.Format("2006/01/02"))
}
