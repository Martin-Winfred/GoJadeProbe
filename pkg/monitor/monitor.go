package monitor

import (
	"errors"
	"log"
	"net"
	"os"
	"runtime"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	net3 "github.com/shirou/gopsutil/v3/net"
)

// HostMonitor Stracture
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

// GetHostMonitor| Rerturn the struct of HostMonitor
func GetHostMonitor(targetInterface string) (hostMonitor HostMonitor, err error) {
	hostMonitor.Arch = runtime.GOARCH
	hostMonitor.OSInfo = runtime.GOOS
	hostMonitor.Hostname, err = os.Hostname()
	if err != nil {
		err = errors.New("can't detect HostInfo")
		return
	}

	hostMonitor.KernelVer, err = host.KernelVersion()
	if err != nil {
		err = errors.New("can not load kernelVerion")
		return
	}

	hostMonitor.Platform, hostMonitor.Family, hostMonitor.Version, err = host.PlatformInformation()
	if err != nil {
		err = errors.New("can't load platform version")
		return
	}

	hostMonitor.CPULoad, err = cpu.Percent(time.Second, false)
	if err != nil {
		err = errors.New("unable to get CPU load per sec")
		return
	}

	memInfo, err := mem.VirtualMemory()
	if err != nil {
		err = errors.New("unable to get memory load per sec")
		return
	}
	hostMonitor.MemUsage = memInfo.UsedPercent
	hostMonitor.MemUsed = memInfo.Used
	hostMonitor.MemTotal = memInfo.Total

	netCard, err := net3.IOCounters(true)
	if err != nil {
		err = errors.New("can't get Net Info")
		return
	}
	for _, card := range netCard {
		if card.Name == targetInterface { // 修改为你需要的网络接口名称
			hostMonitor.NetName = card.Name
			hostMonitor.BytesRecv = card.BytesRecv
			hostMonitor.BytesSent = card.BytesSent
			break
		}
	}

	hostMonitor.LocalIP = getLocalIP()

	return hostMonitor, nil
}

func getLocalIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}
