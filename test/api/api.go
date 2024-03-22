package api

import (
	"bytes"
	"encoding/json"
	"net/http"

	"github.com/Martin-Winfred/GoJadeProbe/pkg/monitor"
)

type Probe struct {
	CPULoad   []float64 `json:"cpu_load"`
	MemUsage  float64   `json:"mem_usage"`
	MemUsed   uint64    `json:"mem_used"`
	MemTotal  uint64    `json:"mem_total"`
	NetName   string    `json:"net_name"`
	BytesRecv uint64    `json:"bytes_recv"`
	BytesSent uint64    `json:"bytes_sent"`
	LocalIP   string    `json:"local_ip"`
}

type HostInfo struct {
	Arch      string `json:"arch"`
	OSInfo    string `json:"os_info"`
	Hostname  string `json:"hostname"`
	KernelVer string `json:"kernel_ver"`
	OSVersion string `json:"os_version"`
	Platform  string `json:"platform"`
	Family    string `json:"family"`
}

func SentProbeData(remoteHost string, iface string) error {
	// Copy data to new struct
	data, _ := monitor.GetProbe(iface)
	newData := Probe{
		CPULoad:   data.CPULoad,
		MemUsage:  data.MemUsage,
		MemUsed:   data.MemUsed,
		MemTotal:  data.MemTotal,
		NetName:   data.NetName,
		BytesRecv: data.BytesRecv,
		BytesSent: data.BytesSent,
		LocalIP:   data.LocalIP,
	}

	jsonData, err := json.Marshal(newData)
	if err != nil {
		return err
	}

	resp, err := http.Post(remoteHost, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}

func SentHostInfo(remoteHost string, iface string) error {
	// Copy data to new struct
	data, _ := monitor.GetHostInfo()
	newData := HostInfo{
		Arch:      data.Arch,
		OSInfo:    data.OSInfo,
		Hostname:  data.Hostname,
		KernelVer: data.KernelVer,
		OSVersion: data.OSVersion,
		Platform:  data.Platform,
		Family:    data.Family,
	}

	jsonData, err := json.Marshal(newData)
	if err != nil {
		return err
	}

	resp, err := http.Post(remoteHost, "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
