package config

import "sync"

type Config interface {
	Build()
	IsSet(key string) bool
	GetString(key string) string
	GetInt(key string) int
	GetBool(key string) bool
	GetFloat(key string) float64
	GetStringSlice(key string) []string
	GetIntSlice(key string) []int
	GetStringMap(key string) map[string]interface{}
	GetStringMapString(key string) map[string]string
}

var once sync.Once
var instance Config

func Default() Config {
	once.Do(func() {
		instance = ViperConfig{}
		instance.Build()
	})
	return instance
}
