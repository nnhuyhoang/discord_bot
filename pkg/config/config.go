package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	ServiceName      string
	BaseURL          string
	Port             string
	Env              string
	AllowedOrigins   string
	DiscordBotName   string
	DiscordBotToken  string
	DiscordChannelId string
}

// Loader load config from reader into Viper
type Loader interface {
	Load(viper.Viper) (*viper.Viper, error)
}

// generateConfigFromViper generate config from viper data
func generateConfigFromViper(v *viper.Viper) Config {

	return Config{
		Port:        v.GetString("PORT"),
		BaseURL:     v.GetString("BASE_URL"),
		ServiceName: v.GetString("SERVICE_NAME"),
		Env:         v.GetString("ENV"),

		AllowedOrigins: v.GetString("ALLOWED_ORIGINS"),

		DiscordBotName:   v.GetString("DISCORD_BOT_NAME"),
		DiscordBotToken:  v.GetString("DISCORD_BOT_TOKEN"),
		DiscordChannelId: v.GetString("CHANNEL_ID"),
	}
}

// DefaultConfigLoaders is default loader list
func DefaultConfigLoaders() []Loader {
	loaders := []Loader{}
	fileLoader := NewFileLoader(".env", ".")
	loaders = append(loaders, fileLoader)
	loaders = append(loaders, NewENVLoader())

	return loaders
}

// LoadConfig load config from loader list
func LoadConfig(loaders []Loader) Config {
	v := viper.New()
	v.SetDefault("PORT", "8080")
	v.SetDefault("ENV", "local")

	for idx := range loaders {
		newV, err := loaders[idx].Load(*v)

		if err == nil {
			v = newV
		}
	}
	return generateConfigFromViper(v)
}
