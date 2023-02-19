package config

import (
	"github.com/pelletier/go-toml/v2"
	"github.com/sirupsen/logrus"
	"os"
)

type animeLoadConfig struct {
	Name   string `toml:"Name"`
	Detail struct {
		AnimeNameExclusion []string `toml:"AnimeName_exclusion"`
		AnimeName          string   `toml:"AnimeName"`
		DefaultSeason      string   `toml:"DefaultSeason"`
		CategoryExclusion  []string `toml:"Category_exclusion"`
		Category           string   `toml:"Category"`
	} `toml:"Detail"`
}

var AnimeLoadConfig animeLoadConfig

func initAnimeLoadConfig() {
	configDir := FlagConfig.ConfigDir

	if configDir == "" {
		logrus.Panic("Config must be set.")
	}
	var data []byte
	if tempData, err := os.ReadFile(configDir); err != nil {
		logrus.Panic(err)
	} else {
		data = tempData
	}
	if err := toml.Unmarshal(data, &AnimeLoadConfig); err != nil {
		logrus.Panic(err)
	}
}
