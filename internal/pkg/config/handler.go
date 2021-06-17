package config

import (
	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
	"io/ioutil"
)

type Config struct {
	DBUser string `yaml:"dbUser"`
	DBPass string `yaml:"dbPass"`
	DBPort string `yaml:"dbPort"`
	DBName string `yaml:"dbName"`
	DBCollectionAllcards string `yaml:"dbCollectionAllcards"`
	DBCollectionMycards string `yaml:"dbCollectionMycards"`
	DBCollectionSetimages string `yaml:"dbCollectionSetimages"`
	DBCollectionSetNames string `yaml:"dbCollectionSetNames"`
}

func GetConfig(configFile string) (Config, error) {
	var c Config

	buf, err := ioutil.ReadFile(configFile)
	if err != nil {
		log.Error().Err(err)
		return c, err
	}

	if err = yaml.Unmarshal(buf, &c); err != nil {
		log.Error().Err(err)
		return c, err
	}

	return c, err
}