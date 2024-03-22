package monitor

import (
	"errors"
	"net"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	net3 "github.com/shirou/gopsutil/v3/net"
)

// HostMonitor Structure
type Probe struct {
	CPULoad   []float64
	MemUsage  float64
	MemUsed   uint64
	MemTotal  uint64
	NetName   string
	BytesRecv uint64
	BytesSent uint64
	LocalIP   string
}

type HostInfo struct {
	Arch      string
	OSInfo    string
	Hostname  string
	KernelVer string
	OSVersion string
	Platform  string
	Family    string
}

// GetProbe returns the Probe struct
func GetProbe(targetInterface string) (probe Probe, err error) {
	probe.CPULoad, err = cpu.Percent(time.Second, false)
	if err != nil {
		return probe, errors.New("unable to get CPU load per sec")
	}

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return probe, errors.New("unable to get memory load per sec")
	}
	probe.MemUsage = memInfo.UsedPercent
	probe.MemUsed = memInfo.Used
	probe.MemTotal = memInfo.Total

	netCard, err := net3.IOCounters(true)
	if err != nil {
		return probe, errors.New("can't get Net Info")
	}
	for _, card := range netCard {
		if card.Name == targetInterface {
			probe.NetName = card.Name
			probe.BytesRecv = card.BytesRecv
			probe.BytesSent = card.BytesSent
			break
		}
	}

	probe.LocalIP, err = GetLocalIP()
	if err != nil {
		return probe, err
	}

	return probe, nil
}

// GetHostInfo returns the HostInfo struct
func GetHostInfo() (hostInfo HostInfo, err error) {
	hostInfo.Arch = runtime.GOARCH
	hostInfo.OSInfo = runtime.GOOS
	hostInfo.Hostname, err = os.Hostname()
	if err != nil {
		return hostInfo, errors.New("can't detect HostInfo")
	}

	hostInfo.KernelVer, err = host.KernelVersion()
	if err != nil {
		return hostInfo, errors.New("can not load kernelVerion")
	}

	hostInfo.Platform, hostInfo.Family, hostInfo.OSVersion, err = host.PlatformInformation()
	if err != nil {
		return hostInfo, errors.New("can't load platform version")
	}

	return hostInfo, nil
}

func GetLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}
