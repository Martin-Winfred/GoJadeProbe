// Description: This file is used to test the api package. in pkg/api/api.go

package main

import (
	myapi "github.com/Martin-Winfred/GoJadeProbe/pkg/api"
)

func main() {
	myapi.SentHostInfo("http://127.0.0.1:5000/api/info", "eth0", "123456")
	myapi.SentProbeData("http://127.0.0.1:5000/api/probe", "eth0", "123456")
	myapi.SentEncHostInfo("http://127.0.0.1:5000/api/info", "eth0", "123456", "dhrydhfjfufj#$6&")
	myapi.SentEncProbeData("http://127.0.0.1:5000/api/probe", "eth0", "123456", "asdf1234asdf1234")

}
