package config

import "github.com/spf13/viper"

type App struct {
	Name          string `json:"Name" mapstructure:"name"`
	Version       string `json:"Version" mapstructure:"version"`
	Description   string `json:"Description" mapstructure:"description"`
	Port          int    `json:"Port" mapstructure:"port"`
	LogLevel      string `json:"LogLevel" mapstructure:"log-level"`
	IsDevelopment bool   `json:"IsDevelopment" mapstructure:"is-development"`
}

type Postgres struct {
	Host         string `json:"Host" mapstructure:"host"`
	Port         int    `json:"Port" mapstructure:"port"`
	User         string `json:"User" mapstructure:"username"`
	Password     string `json:"Password" mapstructure:"password"`
	Database     string `json:"Database" mapstructure:"database"`
	Schema       string `json:"Schema" mapstructure:"schema"`
	Ssl          string `json:"Ssl" mapstructure:"ssl"`
	MaxIdleTime  int    `json:"MaxIdleTime" mapstructure:"max-idle-time"`
	MaxLifeTime  int    `json:"MaxLifeTime" mapstructure:"max-life-time"`
	MaxOpenConns int    `json:"MaxOpenConns" mapstructure:"max-open-conns"`
	MaxIdleConns int    `json:"MaxIdleConns" mapstructure:"max-idle-conns"`
}

type AppConfig struct {
	App      App      `json:"App" mapstructure:"app"`
	Postgres Postgres `json:"Postgres" mapstructure:"postgres"`
}

func AppConfigInit() (*AppConfig, error) {
	var appConfig AppConfig
	viper.SetConfigName("app.config")
	viper.SetConfigType("json")
	viper.AddConfigPath("file/configs")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return &appConfig, err
	}

	err = viper.Unmarshal(&appConfig)
	if err != nil {
		return &appConfig, err
	}

	return &appConfig, nil
}
