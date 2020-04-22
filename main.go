package main

import (
	"pigLatin/api"
	"pigLatin/db"
	service "pigLatin/services"

	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var configFile string

//initialize config
func init() {
	if configFile != "" {
		viper.SetConfigFile(configFile)
	} else {
		viper.SetConfigName("config")
		viper.AddConfigPath(".")
		viper.AddConfigPath("$HOME/.pigLatin")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		logrus.WithError(err).Warnf("unable to read config from file")
	}
}

func main() {

	db, err := db.NewDB(viper.GetString("DatabaseURI"))

	if err != nil {
		logrus.WithError(err).Warnf("cannot connect to database")
	}

	pigService := service.NewPiglatinService(db)
	router := mux.NewRouter()
	app := api.New(pigService, router)

	app.Start()

}
