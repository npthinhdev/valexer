package env

import (
	"log"

	"github.com/kelseyhightower/envconfig"
)

var envPrefix = ""

// Load loads the environment variables into the provided struct
func Load(t interface{}) {
	err := envconfig.Process(envPrefix, t)
	if err != nil {
		log.Printf("config: unable to load config for %T: %s", t, err)
	}
}
