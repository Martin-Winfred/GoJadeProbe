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
	InterfaceName  string
	Encrypted      bool
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

	encryted, err := sec.Key("encrypted").Bool()

	sec = cfg.Section("report")
	if sec == nil {
		return nil, fmt.Errorf("missing 'report' section in config file")
	}

	reportInterval, err := sec.Key("interval").Int()
	if err != nil {
		return nil, fmt.Errorf("invalid 'interval' value in config file: %v", err)
	}

	sec = cfg.Section("local")
	if sec == nil {
		return nil, fmt.Errorf("missing 'local' section in config file")
	}

	uuid := sec.Key("uuid").String()
	if uuid == "" {
		uuid = guuid.New().String()
		sec.Key("uuid").SetValue(uuid)
		err = cfg.SaveTo("./config.ini")
		if err != nil {
			return nil, fmt.Errorf("failed to write UUID to config file: %v", err)
		}
	}

	interfaceName := sec.Key("interfaceName").String()
	if interfaceName == "" {
		return nil, fmt.Errorf("missing 'InterfaceName' value in config file")
	}

	config := &Myconfig{
		Address:        address,
		Port:           port,
		Password:       password,
		UUID:           uuid,
		ReportInterval: reportInterval,
		InterfaceName:  interfaceName,
		Encrypted:      encryted,
	}

	return config, nil
}
