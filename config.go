package main

import (
	"time"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

type (
	Config struct {
		Db      Db                `json:"db" yaml:"db"`
		Tag     []string          `json:"tag" yaml:"tag"`
		TypeMap map[string]GoType `json:"typeMap" yaml:"typeMap"`
		Case    Case              `json:"case" yaml:"case"`
		Path    Path              `json:"path" yaml:"path"`
	}
	// Db 数据库信息
	Db struct {
		Username string `json:"username" yaml:"username"`
		Password string `json:"password" yaml:"password"`
		Host     string `json:"host" yaml:"host"`
	}
	Case struct {
		Old CaseType `json:"old" yaml:"old"`
		New CaseType `json:"new" yaml:"new"`
	}
	Path struct {
		Template string `json:"template" yaml:"template"`
	}
)

var (
	cfg = new(Config)
)

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

	if err := viper.Unmarshal(cfg); err != nil {
		log.Error("viper unmarshal error", zap.Error(err))
		return
	}

	if len(cfg.TypeMap) == 0 {
		cfg.TypeMap = GoTypeMap
	}

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
