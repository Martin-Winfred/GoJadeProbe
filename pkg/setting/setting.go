package setting

import (
	"fmt"

	guuid "github.com/google/uuid"

	iniv1 "gopkg.in/ini.v1"
)

type Myconfig struct {
	Address        string
	Port           int
	Password       string
	UUID           string
	ReportInterval int
}

func LoadIni() (*Myconfig, error) {
	cfg, err := iniv1.Load("./config.ini")
	if err != nil {
		return nil, fmt.Errorf("failed to load config file: %v", err)
	}

	sec := cfg.Section("server")
	if sec == nil {
		return nil, fmt.Errorf("missing 'server' section in config file")
	}

	address := sec.Key("address").String()
	port, err := sec.Key("port").Int()
	if err != nil {
		return nil, fmt.Errorf("invalid 'port' value in config file: %v", err)
	}

	sec = cfg.Section("auth")
	if sec == nil {
		return nil, fmt.Errorf("missing 'auth' section in config file")
	}

	password := sec.Key("password").String()

	uuid := sec.Key("uuid").String()
	if uuid == "" {
		uuid = guuid.New().String()
		sec.Key("uuid").SetValue(uuid)
		err = cfg.SaveTo("./config.ini")
		if err != nil {
			return nil, fmt.Errorf("failed to write UUID to config file: %v", err)
		}
	}

	sec = cfg.Section("report")
	if sec == nil {
		return nil, fmt.Errorf("missing 'report' section in config file")
	}

	reportInterval, err := sec.Key("interval").Int()
	if err != nil {
		return nil, fmt.Errorf("invalid 'interval' value in config file: %v", err)
	}

	config := &Myconfig{
		Address:        address,
		Port:           port,
		Password:       password,
		UUID:           uuid,
		ReportInterval: reportInterval,
	}

	return config, nil
}
