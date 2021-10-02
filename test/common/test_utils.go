package common

import (
	"encoding/json"
	"os"

	"github.com/pdkovacs/igo-repo/config"
)

func CloneConfig(conf config.Options) config.Options {
	var clone config.Options
	var err error

	var configAsJSON []byte
	configAsJSON, err = json.Marshal(conf)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(configAsJSON), &clone)
	if err != nil {
		panic(err)
	}

	return clone
}

func GetTestConfig() config.Options {
	oldDBHostEnvVar := os.Getenv("DB_HOST")
	dbHostEnvVar := os.Getenv("ICONREPO_DB_HOST")
	if dbHostEnvVar != "" {
		os.Setenv("DB_HOST", dbHostEnvVar)
	}
	config := CloneConfig(config.GetDefaultConfiguration())
	os.Setenv("DB_HOST", oldDBHostEnvVar)
	return config
}
