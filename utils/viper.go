package utils

import "github.com/spf13/viper"

// ViperInt starts the viper instance
func ViperInt() {
	viper.SetConfigName("config")
	viper.AddConfigPath(".")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	Check(err, "viper.ReadInConfig")
}
