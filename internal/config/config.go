package config

import "github.com/spf13/viper"

type Config struct {
	Database struct {
		URL string `mapstructure:"url"`
	} `mapstructure:"database"`
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
	Security struct {
        JWTSecret string `mapstructure:"jwt_secret"`
    } `mapstructure:"security"`

}

func LoadConfig() (config Config, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&config)
	return
}