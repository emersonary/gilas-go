package configs

import "github.com/spf13/viper"

var cfg *Conf

type Conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBUser        string `mapstructure:"DB_DBUSER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	DBName        string `mapstructure:"DB_NAME"`
	WebServerPort string `mapstructure:"DB_SERVER_PORT"`
}

func LoadConfig(path string) (*Conf, error) {

	viper.SetConfigName("server")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.AddConfigPath("gilasw.env")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	err = viper.Unmarshal(&cfg)
	if err != nil {
		return nil, err
	}

	return cfg, nil

}
