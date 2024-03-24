package config

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)


type Config struct{
	PostgreString string
	RedisString string
	RedisPassword string
	RedisDB int
}



func (c *Config) InitConfig() (err error) {
	viper.SetConfigName(".env")
	viper.SetConfigType("env")
	viper.AddConfigPath("./pkg/common/env")
	err = viper.ReadInConfig()
	if err != nil {
		logrus.Fatalf("Error while reading config file %s", err)
		return err
	}

	c.PostgreString = viper.GetString("POSTGRE_STRING")
	c.RedisString = viper.GetString("REDIS_STRING")
	c.RedisPassword = viper.GetString("REDIS_PASSWORD")
	c.RedisDB = viper.GetInt("RREDIS_DB")
	

	
	return nil

}
