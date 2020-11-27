package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
	"strings"
)

func InitYAMLConfig(cfgFile string) {
	viper.SetConfigType("yaml")
	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigName("sites")
	}

	err := viper.ReadInConfig()
	if err != nil {
		panic(errors.Wrapf(err, "Error reading config file"))
	}

	// read config from environment too
	viper.SetEnvPrefix("sites")
	replacer := strings.NewReplacer("-", "_")
	viper.SetEnvKeyReplacer(replacer)
	viper.AutomaticEnv()
}
