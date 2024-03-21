package main

import (
	"fmt"
	"log"

	"github.com/Martin-Winfred/GoJadeProbe/pkg/setting"
)

func main() {
	config, err := setting.LoadIni()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Address:", config.Address)
	fmt.Println("Port:", config.Port)
	fmt.Println("Password:", config.Password)
	fmt.Println("UUID:", config.UUID)
	fmt.Println("Report Interval:", config.ReportInterval)
}
