package helper

import (
	"github.com/caijunduo/go-scaffold/internal/pkg/log"
	"github.com/spf13/viper"
)

func GetLogOption() *log.Option {
	return &log.Option{
		DisableCaller:     viper.GetBool("log.disable-caller"),
		DisableStacktrace: viper.GetBool("log.disable-stacktrace"),
		Level:             viper.GetString("log.level"),
		Format:            viper.GetString("log.format"),
		OutputPaths:       viper.GetStringSlice("log.output-paths"),
	}
}
