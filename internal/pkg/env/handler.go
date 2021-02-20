package env

import (
	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

func ReceiveEnvVars() (Conf, error) {
	var err error
	viper.SetConfigFile(".env")

	if err = viper.ReadInConfig(); err != nil {
		log.Error().Timestamp().Err(err).Msg("Error: loading env vars")
		return Conf{}, err
	}

	conf := Conf{
		DbUser: viper.Get("DB_USER").(string),
		DbPass: viper.Get("DB_PASS").(string),
		DbPort: viper.Get("DB_PORT").(string),
		DbName: viper.Get("DB_NAME").(string),
		DbCollAllCards: viper.Get("DB_COLLECTION_ALLCARDS").(string),
		DbCollMyCards: viper.Get("DB_COLLECTION_ALLCARDS").(string),
		DbCollSetImages: viper.Get("DB_COLLECTION_ALLCARDS").(string),
	}

	return conf, err
}