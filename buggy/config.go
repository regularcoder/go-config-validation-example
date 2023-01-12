package buggy

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/go-playground/validator/v10"
	"log"
	"os"
)

type config struct {
	ServerUrl string `json:"server_url" validate:"required"`
	AppPort   int    `json:"app_port" validate:"required,numeric,gte=8080,lte=8085"`
}

type Config interface {
	GetServerURL() string
	GetAppPort() int
	GetPrivateKey() string
}

func (c config) GetServerURL() string {
	return c.ServerUrl
}

func (c config) GetAppPort() int {
	return c.AppPort
}

func (c config) GetPrivateKey() string {
	return os.Getenv("PRIVATE_KEY")
}

func LoadConfig() (Config, error) {
	configFile, err := os.Open("configuration/config.json")
	if err != nil {
		log.Fatal("Could not open config file : ", err.Error())
	}

	decoder := json.NewDecoder(configFile)
	config := config{}

	decodeErr := decoder.Decode(&config)
	if decodeErr != nil {
		log.Fatal("Could not decode config file : ", decodeErr.Error())
	}

	if !Validate(config) {
		return nil, errors.New("invalid config")
	}
	return config, nil
}

func Validate(config Config) bool {
	validate := validator.New()
	err := validate.Struct(config)
	if err != nil {
		fmt.Println("Invalid config !")
		for _, validationErr := range err.(validator.ValidationErrors) {
			fmt.Println(validationErr.StructNamespace() + " violated " + validationErr.Tag() + " validation.")
		}
		return false
	}
	return true
}
