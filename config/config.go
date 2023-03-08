package config

import (
	"time"

	"github.com/pingcap/log"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type (
	Config struct {
		Db Db `json:"db" yaml:"db"`
	}
	// Db 数据库信息
	Db struct {
		Username string `json:"username" yaml:"username"`
		Password string `json:"password" yaml:"password"`
		Host     string `json:"host" yaml:"host"`
	}
)

var Cfg = new(Config)

var (
	defaultTimezone = "Asia/Shanghai"
)

func InitConfig() {
	viper.SetConfigType("yaml")
	viper.AddConfigPath(".")
	viper.AddConfigPath("$GOPATH/bin")
	if err := viper.ReadInConfig(); err != nil {
		log.Error("read config error", zap.Error(err))
		return
	}

	if err := viper.Unmarshal(Cfg); err != nil {
		log.Error("viper unmarshal error", zap.Error(err))
		return
	}
	//TODO 必须的参数为空时，载入默认配置
	//加载时区
	//_ = initTimezone()
}

// 初始化时区
func initTimezone() error {
	var err error
	time.Local, err = time.LoadLocation(defaultTimezone)
	if err != nil {
		log.Error("load timezone error", zap.Error(err))
		return err
	}
	return nil
}
