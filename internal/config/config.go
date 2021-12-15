package config

import (
	"github.com/spf13/viper"

	"github.com/xylonx/zapx"
	"go.uber.org/zap"
)

var Config *Setting

func Setup(cfgFile string) error {
	v := viper.New()

	v.SetConfigFile(cfgFile)

	setDefaultValue(v)

	if err := v.Unmarshal(Config); err != nil {
		zapx.Error("unmarshal config file failed", zap.Error(err))
	}

	return nil
}

func setDefaultValue(v *viper.Viper) {

}
