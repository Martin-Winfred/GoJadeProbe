package main

import (
	myapi "github.com/Martin-Winfred/GoJadeProbe/pkg/api"
)

func main() {
	myapi.SentData("http://127.0.0.1:5000/api", "eth0")
}
