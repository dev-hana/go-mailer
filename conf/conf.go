package conf

import (
	"fmt"

	"github.com/spf13/viper"
)

func GetDBConfig() (dbms, dsn string, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
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

func GetServerConfig() (port int, mode bool, dbInit bool, err error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()

	err = viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		return 0, false, false, err
	}

	port = viper.GetInt("app.server.port")
	mode = viper.GetBool("app.server.gin-release")
	dbInit = viper.GetBool("app.server.db-init")
	return port, mode, dbInit, nil

}

type SMTPConfig struct {
	Host         string
	Port         int
	User         string
	Password     string
	Verification bool
}

func GetSMTPConfig() (*SMTPConfig, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		return nil, err
	}

	host := viper.GetString("app.smtp.host")
	port := viper.GetInt("app.smtp.port")
	user := viper.GetString("app.smtp.user")
	password := viper.GetString("app.smtp.password")
	vertification := viper.GetBool("app.smtp.vertification")

	return &SMTPConfig{
		Host:         host,
		Port:         port,
		User:         user,
		Password:     password,
		Verification: vertification,
	}, nil
}

func GetSchedulerConfig() (int, error) {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("fatal error config file: default \n", err)
		return 0, err
	}
	return viper.GetInt("app.scheduler.time-second"), nil
}
