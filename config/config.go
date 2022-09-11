package appconfig

import (
	"github.com/spf13/viper"
	"log"
	"os"
)

type Config struct {
	NotificationTime string `mapstructure:"NOTIFICATION_TIME"`
	SmtpHost         string `mapstructure:"SMTP_HOST"`
	SmtpPort         string `mapstructure:"SMTP_PORT"`
	SmtpUsername     string `mapstructure:"SMTP_USERNAME"`
	SmtpPassword     string `mapstructure:"SMTP_PASSWORD"`
	SmtpSender       string `mapstructure:"SMTP_SENDER"`
	APIVersion       string `mapstructure:"API_VERSION"`
	LocalLocation    string `mapstructure:"LOCAL_LOCATION"`
}

func LoadConfig() (config Config, err error) {
	pwd, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	viper.AddConfigPath(pwd)
	viper.SetConfigName("app")
	viper.SetConfigType("env")

	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}
