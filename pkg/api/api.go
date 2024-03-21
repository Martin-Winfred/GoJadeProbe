package api

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"encoding/json"
	"io"
	"net/http"

	"github.com/Martin-Winfred/GoJadeProbe/pkg/monitor"
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

// encrypt encrypts the data using AES.
func encrypt(data []byte, passphrase string) ([]byte, error) {
	block, err := aes.NewCipher([]byte(passphrase))
	if err != nil {
		return nil, err
	}

	// Create a new GCM block
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}

	// Create a nonce
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}

	// Encrypt the data
	ciphertext := gcm.Seal(nonce, nonce, data, nil)
	return ciphertext, nil
}

func SentData(remoteHost string, iface string, passphrase string) error {
	// Copy data to new struct
	data, _ := monitor.GetHostMonitor(iface)
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
		return err
	}

	hash := sha256.Sum256([]byte(passphrase))                // Generate hash from passphrase
	encryptedData, err := encrypt(jsonData, string(hash[:])) // Convert hash to string
	if err != nil {
		return err
	}

	resp, err := http.Post(remoteHost, "application/octet-stream", bytes.NewBuffer(encryptedData))
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	return nil
}
