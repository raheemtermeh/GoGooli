package config

import "github.com/spf13/viper"

// Config یک ساختار برای نگهداری تمام تنظیمات برنامه ماست
// ساختار این struct دقیقاً با ساختار فایل config.yml مطابقت داره
type Config struct {
	Database struct {
		URL string `mapstructure:"url"`
	} `mapstructure:"database"`
	Server struct {
		Port string `mapstructure:"port"`
	} `mapstructure:"server"`
}

// LoadConfig تنظیمات رو از فایل config.yml می‌خونه و توی ساختار Config برمی‌گردونه
func LoadConfig() (config Config, err error) {
	// به Viper میگیم که دنبال فایل‌هایی به اسم config بگرده
	viper.SetConfigName("config")
	// و نوعشون هم yml هست
	viper.SetConfigType("yml")
	// و توی پوشه ریشه پروژه دنبالش بگرده
	viper.AddConfigPath(".")

	// سعی کن فایل رو بخونی
	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	// حالا تنظیمات خونده شده رو توی ساختار Config ما بریز
	err = viper.Unmarshal(&config)
	return
}