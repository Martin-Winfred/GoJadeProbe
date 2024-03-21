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
type HostMonitor struct {
	Arch      string
	OSInfo    string
	Hostname  string
	KernelVer string
	Version   string
	Platform  string
	Family    string
	CPULoad   []float64
	MemUsage  float64
	MemUsed   uint64
	MemTotal  uint64
	NetName   string
	BytesRecv uint64
	BytesSent uint64
	LocalIP   string
}

// GetHostMonitor returns the HostMonitor struct
func GetHostMonitor(targetInterface string) (hostMonitor HostMonitor, err error) {
	hostMonitor.Arch = runtime.GOARCH
	hostMonitor.OSInfo = runtime.GOOS
	hostMonitor.Hostname, err = os.Hostname()
	if err != nil {
		return hostMonitor, errors.New("can't detect HostInfo")
	}

	hostMonitor.KernelVer, err = host.KernelVersion()
	if err != nil {
		return hostMonitor, errors.New("can not load kernelVerion")
	}

	hostMonitor.Platform, hostMonitor.Family, hostMonitor.Version, err = host.PlatformInformation()
	if err != nil {
		return hostMonitor, errors.New("can't load platform version")
	}

	hostMonitor.CPULoad, err = cpu.Percent(time.Second, false)
	if err != nil {
		return hostMonitor, errors.New("unable to get CPU load per sec")
	}

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		return hostMonitor, errors.New("unable to get memory load per sec")
	}
	hostMonitor.MemUsage = memInfo.UsedPercent
	hostMonitor.MemUsed = memInfo.Used
	hostMonitor.MemTotal = memInfo.Total

	netCard, err := net3.IOCounters(true)
	if err != nil {
		return hostMonitor, errors.New("can't get Net Info")
	}
	for _, card := range netCard {
		if card.Name == targetInterface {
			hostMonitor.NetName = card.Name
			hostMonitor.BytesRecv = card.BytesRecv
			hostMonitor.BytesSent = card.BytesSent
			break
		}
	}

	hostMonitor.LocalIP, err = getLocalIP()
	if err != nil {
		return hostMonitor, err
	}

	return hostMonitor, nil
}

func getLocalIP() (string, error) {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		return "", err
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String(), nil
}
