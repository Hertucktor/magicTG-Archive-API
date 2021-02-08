package env

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Conf struct {
	DbUser string
	DbPass string
	DbPort string
	DbName string
}

func ReceiveEnvVars() (Conf, error) {
	var err error
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: loading env vars")
		return Conf{}, err
	}

	conf := Conf{
		DbUser: viper.Get("DB_USER").(string),
		DbPass: viper.Get("DB_PASS").(string),
		DbPort: viper.Get("DB_PORT").(string),
		DbName: viper.Get("DB_NAME").(string),
	}

	return conf, err
}
