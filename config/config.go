package appconfig

import "github.com/spf13/viper"

type Config struct {
	ChronInterval string `mapstructure:"CHRON_INTERVAL"`
	SmtpHost      string `mapstructure:"SMTP_HOST"`
	SmtpPort      string `mapstructure:"SMTP_PORT"`
	SmtpUsername  string `mapstructure:"SMTP_USERNAME"`
	SmtpPassword  string `mapstructure:"SMTP_PASSWORD"`
	SmtpSender    string `mapstructure:"SMTP_SENDER"`
	APIVersion    string `mapstructure:"API_VERSION"`
}

func LoadConfig(path string) (config Config, err error) {
	viper.AddConfigPath(path)
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
