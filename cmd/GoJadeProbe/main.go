package main

import (
	"fmt"

	"github.com/Martin-Winfred/GoJadeProbe/pkg/monitor"
)

func main() {
	//targetInterface := "lo"
	hostMonitor, err := monitor.GetHostMonitor("lo")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("HostMonitor: %+v\n", hostMonitor)
}
