package goenv

import (
	"log"

	"github.com/spf13/viper"
)

type FileEnv struct {
	Path string
	Name string
	Extn string
}

/*
   @notice Load env file into the programme.
   @param _path Path to env file (ex: . or ../).
   @param _name Name of env file (ex: secret).
   @param _type Extension of env file (ex: env).
*/
func GetEnv(_f FileEnv, _result interface{}) (interface{}, error) {
	env, err := loadEnv(_f, _result)

	return env, err
}

/*
   @notice Load env file into the programme.
   @param _path Path to env file (ex: . or ../).
   @param _name Name of env file (ex: secret).
   @param _type Extension of env file (ex: env).
*/
func loadEnv(_f FileEnv, _result interface{}) (interface{}, error) {

	result := _result

	viper.AddConfigPath(_f.Path)
	viper.SetConfigName(_f.Name)
	viper.SetConfigType(_f.Extn)

	viper.AutomaticEnv()

	err := viper.ReadInConfig()

	if err != nil {
		log.Println(err.Error())
	}

	err = viper.Unmarshal(&result)

	if err != nil {
		log.Println(err.Error())
	}

	return result, err
}
