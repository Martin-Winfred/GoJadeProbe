package main

import (
	"fmt"
	"log"
	"time"

	"github.com/Martin-Winfred/GoJadeProbe/pkg/api"
	"github.com/Martin-Winfred/GoJadeProbe/pkg/setting"
)

func main() {
	config, err := setting.LoadIni()
	if err != nil {
		log.Print("Failed to load config file: ", err)
		panic(err)
	}
	// print config information
	println("Remote Address", config.Address)
	println("Remote Port", config.Port)
	println("Node UUID", config.UUID)
	println("Report Interval", config.ReportInterval)
	println("Monitor Intrface", config.InterfaceName)

	// Start the probe
	ticker := time.NewTicker(time.Duration(config.ReportInterval) * time.Second)
	for range ticker.C {
		api.SentData(
			fmt.Sprintf("http://%s:%v/api", config.Address, config.Port),
			fmt.Sprintf("%v", config.InterfaceName),
			config.Password,
		)
	}

}
