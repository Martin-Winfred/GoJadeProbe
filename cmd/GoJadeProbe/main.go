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

	//Sent Host Info
	if config.Encrypted {
		err := api.SentEncHostInfo(
			fmt.Sprintf("http://%s:%v/api/info", config.Address, config.Port),
			fmt.Sprintf("%v", config.InterfaceName),
			config.UUID, config.Password,
		)
		if err != nil {
			log.Print("Failed to send host info: ", err)
			panic(err)
		}
	} else {
		err := api.SentHostInfo(
			fmt.Sprintf("http://%s:%v/api/info", config.Address, config.Port),
			fmt.Sprintf("%v", config.InterfaceName),
			config.UUID,
		)
		if err != nil {
			log.Print("Failed to send host info: ", err)
			panic(err)
		}
	}

	//Sent Probe Data
	ticker := time.NewTicker(time.Duration(config.ReportInterval) * time.Second)
	for range ticker.C {
		if config.Encrypted {
			err := api.SentEncProbeData(
				fmt.Sprintf("http://%s:%v/api/probe", config.Address, config.Port),
				fmt.Sprintf("%v", config.InterfaceName),
				config.UUID, config.Password,
			)
			if err != nil {
				log.Print("Failed to send probe data: ", err)
				panic(err)
			}
		} else {
			err := api.SentProbeData(
				fmt.Sprintf("http://%s:%v/api/probe", config.Address, config.Port),
				fmt.Sprintf("%v", config.InterfaceName),
				config.UUID,
			)
			if err != nil {
				log.Print("Failed to send probe data: ", err)
				panic(err)
			}
		}
	}

}
