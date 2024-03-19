package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/Martin-Winfred/GoJadeProbe/pkg/monitor" // 替换为你的项目和包的实际路径
)

type HostMonitor struct {
	Arch      string    `json:"arch"`
	OSInfo    string    `json:"osInfo"`
	Hostname  string    `json:"hostname"`
	KernelVer string    `json:"kernelVer"`
	Version   string    `json:"version"`
	Platform  string    `json:"platform"`
	Family    string    `json:"family"`
	CPULoad   []float64 `json:"cpuLoad"`
	MemUsage  float64   `json:"memUsage"`
	MemUsed   uint64    `json:"memUsed"`
	MemTotal  uint64    `json:"memTotal"`
	NetName   string    `json:"netName"`
	BytesRecv uint64    `json:"bytesRecv"`
	BytesSent uint64    `json:"bytesSent"`
	LocalIP   string    `json:"localIP"`
}

func main() {
	// Copy data to new struct
	data, _ := monitor.GetHostMonitor("lo")
	newData := HostMonitor{
		Arch:      data.Arch,
		OSInfo:    data.OSInfo,
		Hostname:  data.Hostname,
		KernelVer: data.KernelVer,
		Version:   data.Version,
		Platform:  data.Platform,
		Family:    data.Family,
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
		fmt.Println(err)
		return
	}

	resp, err := http.Post("http://127.0.0.1:5000/api", "application/json", bytes.NewBuffer(jsonData))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	// Print the response status
	fmt.Println(resp.Status)
}
