package env

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Conf struct {
	dbUser string
	dbPass string
	dbPort string
	dbName string
}

func ReceiveEnvVars() (Conf, error) {
	var err error
	viper.SetConfigFile(".env")

	if err := viper.ReadInConfig(); err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: loading env vars")
		return Conf{}, err
	}

	conf := Conf{
		dbUser: viper.Get("DB_USER").(string),
		dbPass: viper.Get("DB_PASS").(string),
		dbPort: viper.Get("DB_PORT").(string),
		dbName: viper.Get("DB_NAME").(string),
	}

	return conf, err
}
