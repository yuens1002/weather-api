package config

import (
	"github.com/spf13/viper"
	util "github.com/yuens1002/server-graphql-go-weather/utils"
)

// DBcfg postgres database config
type DBcfg struct {
	Port     int
	Host     string
	User     string
	Password string
	DBName   string
}

// DBconfig returns db settings
func DBconfig() DBcfg {
	h := viper.GetString("local-host")
	// maybe it's a windows thing
	// "HOST" is being return with quotes around the ["string"]
	h2 := util.Trim(viper.GetString("HOST"), "\"")
	var sfg *viper.Viper
	if h == h2 {
		sfg = viper.Sub("app-db-local")
	} else {
		sfg = viper.Sub("app-db-dev")
	}

	return dbSettings(sfg)
}

func dbSettings(cfg *viper.Viper) DBcfg {
	return DBcfg{
		Port:     cfg.GetInt("port"),
		Host:     cfg.GetString("host"),
		User:     cfg.GetString("user"),
		Password: cfg.GetString("password"),
		DBName:   cfg.GetString("dbName"),
	}
}
