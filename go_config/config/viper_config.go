package config

import (
	"log"
	"os"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type ViperConfig struct{}

func (v ViperConfig) Build() {
	//TODO implement me
	var configFilePath string
	viper.SetConfigFile("config")
	if configFilePath == "" {
		_, err := os.Stat(configFilePath)
		if err != nil {
			log.Fatalf("Errpr when reading stat : %s", err)
		}
		viper.AddConfigPath(configFilePath)
	}
	viper.AddConfigPath(".")
	viper.AddConfigPath("$HOME/.config")
	viper.AddConfigPath("/etc/config")
	viper.AutomaticEnv()
	viper.SetConfigType("yml")
	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file : %s", err)
		viper.OnConfigChange(func(in fsnotify.Event) {
			log.Printf("Config file changed : %s", in.Name)
		})
	}
	viper.WatchConfig()
}

func (v ViperConfig) IsSet(key string) bool {
	//TODO implement me
	return viper.IsSet(key)
}

func (v ViperConfig) GetString(key string) string {
	//TODO implement me
	return viper.GetString(key)
}

func (v ViperConfig) GetInt(key string) int {
	//TODO implement me
	return viper.GetInt(key)
}

func (v ViperConfig) GetBool(key string) bool {
	//TODO implement me
	return viper.GetBool(key)
}

func (v ViperConfig) GetFloat(key string) float64 {
	//TODO implement me
	return viper.GetFloat64(key)
}

func (v ViperConfig) GetStringSlice(key string) []string {
	//TODO implement me
	return viper.GetStringSlice(key)
}

func (v ViperConfig) GetIntSlice(key string) []int {
	//TODO implement me
	return viper.GetIntSlice(key)
}

func (v ViperConfig) GetStringMap(key string) map[string]interface{} {
	//TODO implement me
	return viper.GetStringMap(key)
}

func (v ViperConfig) GetStringMapString(key string) map[string]string {
	//TODO implement me
	return viper.GetStringMapString(key)
}
