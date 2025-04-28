package configs

import (
	"os"

	"github.com/spf13/viper"
)

type conf struct {
	DBDriver          string `mapstructure:"DB_DRIVER"`
	DBHost            string `mapstructure:"DB_HOST"`
	DBPort            string `mapstructure:"DB_PORT"`
	DBUser            string `mapstructure:"DB_USER"`
	DBPassword        string `mapstructure:"DB_PASSWORD"`
	DBName            string `mapstructure:"DB_NAME"`
	HabbitMqHost      string `mapstructure:"HABBITMQ_HOST"`
	HabbitMqPort      string `mapstructure:"HABBITMQ_PORT"`
	HabbitMqUser      string `mapstructure:"HABBITMQ_USER"`
	HabbitMqPassword  string `mapstructure:"HABBITMQ_PASSWORD"`
	WebServerPort     string `mapstructure:"WEB_SERVER_PORT"`
	GRPCServerPort    string `mapstructure:"GRPC_SERVER_PORT"`
	GraphQLServerPort string `mapstructure:"GRAPHQL_SERVER_PORT"`
}

func LoadConfig(path string) (*conf, error) {
	var cfg *conf
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}

	if _, existe := os.LookupEnv("AMBIENTE"); !existe {
		cfg.DBHost = "localhost"
		cfg.HabbitMqHost = "localhost"
	}


	return cfg, err
}
