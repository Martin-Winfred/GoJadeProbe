package setting

import (
	"fmt"
	"os"

	"gopkg.in/ini.v1"
)

type Config struct {
	Address        string
	Port           int
	Password       string
	UUID           string
	ReportInterval int
}

func CheckAndCreateIni(path string) error {
	_, err := os.Stat(path)
	if os.IsNotExist(err) {
		cfg := ini.Empty()

		remote, _ := cfg.NewSection("remote")
		remote.NewKey("address", "example.com")
		remote.NewKey("port", "8080")

		credentials, _ := cfg.NewSection("credentials")
		credentials.NewKey("password", "your_password")
		credentials.NewKey("uuid", "your_uuid")

		settings, _ := cfg.NewSection("settings")
		settings.NewKey("report_interval", "60")

		err = cfg.SaveTo(path)
		if err != nil {
			return err
		}
	} else if err != nil {
		return err
	}

	return nil
}

func LoadConfig(path string) (*Config, error) {
	cfg, err := ini.Load(path)
	if err != nil {
		return nil, fmt.Errorf("fail to read file: %v", err)
	}

	config := &Config{
		Address:        cfg.Section("remote").Key("address").String(),
		Port:           cfg.Section("remote").Key("port").MustInt(),
		Password:       cfg.Section("credentials").Key("password").String(),
		UUID:           cfg.Section("credentials").Key("uuid").String(),
		ReportInterval: cfg.Section("settings").Key("report_interval").MustInt(),
	}

	return config, nil
}

func SaveConfig(path string, config *Config) error {
	cfg := ini.Empty()

	remoteSection, err := cfg.NewSection("remote")
	if err != nil {
		return err
	}
	remoteSection.NewKey("address", config.Address)
	remoteSection.NewKey("port", fmt.Sprint(config.Port))

	credentialsSection, err := cfg.NewSection("credentials")
	if err != nil {
		return err
	}
	credentialsSection.NewKey("password", config.Password)
	credentialsSection.NewKey("uuid", config.UUID)

	settingsSection, err := cfg.NewSection("settings")
	if err != nil {
		return err
	}
	settingsSection.NewKey("report_interval", fmt.Sprint(config.ReportInterval))

	return cfg.SaveTo(path)
}
