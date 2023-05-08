package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetDBConfig() (dbms, dsn string, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		return "", "", err
	}

	dbms = viper.GetString("app.db.dbms")
	host := viper.GetString("app.db.host")
	port := viper.GetString("app.db.port")
	db := viper.GetString("app.db.db")
	user := viper.GetString("app.db.user")
	password := viper.GetString("app.db.password")

	return dbms, fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Seoul", host, user, password, db, port), nil
}

func GetServerConfig() (bool, int, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./conf")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		return false, 0, err
	}

	debug := viper.GetBool("app.server.gin-release")
	port := viper.GetInt("app.server.port")
	return debug, port, nil

}
