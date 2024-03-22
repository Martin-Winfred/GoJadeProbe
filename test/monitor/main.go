package main

import (
	"fmt"

	"github.com/Martin-Winfred/GoJadeProbe/pkg/monitor"
)

func main() {
	//targetInterface := "lo"
	Probe, err := monitor.GetProbe("lo")
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("HostMonitor: %+v\n", Probe)
	fmt.Println("=======================")
	hostMonitor, err := monitor.GetHostInfo()
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return
	}
	fmt.Printf("HostMonitor: %+v\n", hostMonitor)
}
