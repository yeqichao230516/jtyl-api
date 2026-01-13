package initialize

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func Viper(logrus *logrus.Logger) *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(".")
	err := v.ReadInConfig()
	if err != nil {
		logrus.Fatalf("读取配置文件失败: %v", err)
	}
	return v
}
