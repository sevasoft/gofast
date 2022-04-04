package goenv

import (
	"log"

	"github.com/spf13/viper"
)

/*
   @notice Load env file into the programme.
   @param _path Path to env file (ex: . or ../).
   @param _name Name of env file (ex: secret).
   @param _type Extension of env file (ex: env).
*/
func GetEnv(_path, _name, _type string) (interface{}, error) {
	env, err := loadEnv(_path, _name, _type)

	return env, err
}

/*
   @notice Load env file into the programme.
   @param _path Path to env file (ex: . or ../).
   @param _name Name of env file (ex: secret).
   @param _type Extension of env file (ex: env).
*/
func loadEnv(_path, _name, _type string) (env interface{}, err error) {
	viper.AddConfigPath(_path)
	viper.SetConfigName(_name)
	viper.SetConfigType(_type)

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err.Error())
	}

	err = viper.Unmarshal(&env)
	return
}
