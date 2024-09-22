package gtools

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// LoadConfig load config file
func LoadConfig(cf string) map[string]any {
	viper.SetConfigName("application")
	viper.SetConfigType("yaml")
	viper.AddConfigPath(cf)
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}
	cm := viper.AllSettings()
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		cm = viper.AllSettings()
	})
	return cm
}

func PrintConfigMap(config map[string]interface{}) {
	for key, value := range config {
		fmt.Printf("%s: %v\n", key, value)
	}
}
