package config

import (
	"fmt"
	"time"

	cerror "github.com/BackAged/go-ddd-microservice/error"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

// PostGreSQL holds PostGreSQL config
type PostGreSQL struct {
	DBURL     string        `yaml:"db_url"`
	DBTimeOut time.Duration `yaml:"db_time_out"`
}

// RabbitMQ holds rabbitmq config
type RabbitMQ struct {
	RabbitURL string `yaml:"rabbit_url"`
}

// Application holds application configurations
type Application struct {
	Host       string     `yaml:"host"`
	Port       int        `yaml:"port"`
	PostGreSQL PostGreSQL `yaml:"postgresql"`
	RabbitMQ   RabbitMQ   `yaml:"rabbitmq"`
}

// IsValid checks if application configuration is valid
func (a Application) IsValid() (bool, error) {
	ve := cerror.ValidationError{}

	if a.Port == 0 {
		ve.Add("port", "is invalid")
	}
	if a.PostGreSQL.DBURL == "" {
		ve.Add("db_url", "is invalid")
	}
	if a.PostGreSQL.DBTimeOut <= 0 {
		ve.Add("db_timeout", "is invalid")
	}
	if a.RabbitMQ.RabbitURL == "" {
		ve.Add("rabbit_url", "is invalid")
	}
	if ve.HasErrors() {
		return false, ve
	}

	return true, nil
}

// GetApp returns application config
func GetApp() (*Application, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Println(".env not found, that's okay!")
	}

	viper.AutomaticEnv()

	cnf := &Application{
		Host: viper.GetString("HOST"),
		Port: viper.GetInt("PORT"),
		PostGreSQL: PostGreSQL{
			DBURL:     viper.GetString("DB_URL"),
			DBTimeOut: viper.GetDuration("DB_TIME_OUT") * time.Second,
		},
		RabbitMQ: RabbitMQ{
			RabbitURL: viper.GetString("RABBITMQ_URL"),
		},
	}
	if ok, err := cnf.IsValid(); !ok {
		return nil, err
	}

	return cnf, nil
}
