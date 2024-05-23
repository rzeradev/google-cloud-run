package configs

import (
	"os"
	"path"

	"github.com/spf13/viper"
)

type Config struct {
	WeatherAPIKey string `mapstructure:"WEATHER_API_KEY"`
	CepAPIURL     string `mapstructure:"CEP_API_URL"`
	WeatherAPIURL string `mapstructure:"WEATHER_API_URL"`
	ServerPort    string `mapstructure:"SERVER_PORT"`
}

func defaultAndBindings() error {
	defaultConfigs := map[string]string{
		"SERVER_PORT":     "8080",
		"WEATHER_API_KEY": "",
		"WEATHER_API_URL": "https://api.weatherapi.com/v1/current.json?key=%s&q=%s",
		"CEP_API_URL":     "https://viacep.com.br/ws/%s/json/",
	}
	for envKey, envValue := range defaultConfigs {
		err := viper.BindEnv(envKey)
		if err != nil {
			return err
		}
		viper.SetDefault(envKey, envValue)
	}
	return nil

}
func LoadConfig(workdir string) (*Config, error) {
	var cfg *Config
	viper.SetConfigName("app_config")
	_, err := os.Stat(path.Join(workdir, ".env"))
	if err == nil {
		viper.SetConfigType("env")
		viper.AddConfigPath(workdir)
		viper.SetConfigFile(".env")
		err = viper.ReadInConfig()
		if err != nil {
			panic(err)
		}
	}
	viper.AutomaticEnv()
	err = defaultAndBindings()
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	return cfg, err
}
