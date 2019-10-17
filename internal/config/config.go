package config

import (
	"os"
	"time"

	"github.com/rs/zerolog/log"
	"gopkg.in/yaml.v2"
)

type Conf struct {
	PortName    string `yaml:"port"`
	Baud        int
	ReadTimeout time.Duration
	Size        byte
	Detector    []string `yaml:"detector"`
	Qualifiers  []string `yaml:"qualifiers"`
}

func NewConfig(filename string) Conf {

	conf := &Conf{
		Baud:        9600,
		ReadTimeout: 1,
		Size:        8,
	}

	f, err := os.Open(filename)
	if err != nil {
		log.Fatal().Err(err).Caller().Msg("")
	}

	if err = yaml.NewDecoder(f).Decode(conf); err != nil {
		log.Fatal().Err(err).Caller().Msg("")
	}

	return *conf
}
