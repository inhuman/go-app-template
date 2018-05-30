package config

import (
	configMerger "github.com/inhuman/config_merger"
)

var App *AppConfig

func Init(path string) error {

	confStruct := &AppConfig{}

	m := configMerger.NewMerger(confStruct)

	m.AddSource(&configMerger.JsonSource{
		Path: path,
	})

	err := m.Run()

	if err != nil {
		return err
	}
	return nil

}

type AppConfig struct {
	LogLevel string `json:"log_level"`
}
