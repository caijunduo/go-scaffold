package config

import (
	"fmt"
	"github.com/spf13/viper"
	"strings"
)

type Option struct {
	File       string
	EnvPrefix  string
	ConfigName string
}

func New(opt Option) {
	if opt.File != "" {
		viper.SetConfigFile(opt.File)
	} else {
		viper.AddConfigPath(".")
		viper.SetConfigType("yaml")
		viper.SetConfigName(opt.ConfigName)
	}

	viper.AutomaticEnv()
	viper.SetEnvPrefix(opt.EnvPrefix)
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_", "-", "_"))

	if err := viper.ReadInConfig(); err != nil {
		fmt.Println("Failed to read viper configuration file, err: ", err)
	}
}
