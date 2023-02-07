package configs

import (
	"github.com/spf13/viper"
	"tgbot/dao"
)

func ParseConfig() (*dao.PostgresDB, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath(".")

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		return nil, err
	}

	var cfg *dao.PostgresDB
	err = viper.UnmarshalKey("db", &cfg)

	return cfg, nil
}
