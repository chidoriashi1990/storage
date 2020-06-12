package config

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

// EnviromentVariable is a structure for storing environment variables
type EnviromentVariable struct {
	AccessControlAllowOrigin string   `required:"true"  split_words:"true"`
	BaseDir                  string   `required:"true"  split_words:"true"`
	Exclude                  []string `required:"false" split_words:"true"`
	IsDebugMode              bool     `required:"true"  split_words:"true" default:"false"`
}

var (
	// Config is a configuration value read from an environment variable
	Config EnviromentVariable
)

func init() {
	err := envconfig.Process("", &Config)
	if err != nil {
		log.Fatal(err)
	}
}
