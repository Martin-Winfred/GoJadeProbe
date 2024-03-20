// Description: This file is used to test the api package. in pkg/api/api.go

package main

import (
	myapi "github.com/Martin-Winfred/GoJadeProbe/pkg/api"
)

func main() {
	myapi.SentData("http://127.0.0.1:5000/api", "eth0")
}
