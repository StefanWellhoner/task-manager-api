package config

import (
	"fmt"

	"github.com/StefanWellhoner/task-manager-api/internal/mode"
	"github.com/jinzhu/configor"
)

type Configuration struct {
	Server struct {
		Port                  int    `default:"8080"`
		ListenAddr            string `default:""`
		KeepAlivePeriodSecond int    `default:"0"`
		SSL                   struct {
			Enabled         bool   `default:"false"`
			RedirectToHTTPS bool   `default:"false"`
			ListenAddr      string `default:""`
			Port            int    `default:"443"`
			CertFile        string `default:""`
			KeyFile         string `default:""`
			LetsEncrypt     struct {
				Enabled   bool     `default:"false"`
				Hosts     []string `default:""`
				AcceptTOS bool     `default:"false"`
				CacheDir  string   `default:"data/certs"`
			}
		}
	}
	Database struct {
		Host     string `default:"localhost"`
		Port     int    `default:"5432"`
		User     string `default:"postgres"`
		Password string `default:"postgres"`
		Database string `default:"postgres"`
		SSLMode  string `default:"disable"`
	}
	Secrets struct {
		Jwt string `default:"secret"`
	}

	PassStrength int `default:"10"`
}

func configFiles() string {
	switch mode.GetEnv() {
	case mode.Prod:
		return "config.yml"
	case mode.Dev:
		return "config.dev.yml"
	case mode.Test:
		return "config.test.yml"
	default:
		panic("Unknown environment")
	}
}

func Get() *Configuration {
	conf := new(Configuration)
	err := configor.Load(conf, configFiles())
	if err != nil {
		panic(err)
	}
	fmt.Printf("Config loaded from %s\n", configFiles())
	return conf
}
