package config

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"sync"

	"github.com/joho/godotenv"
)

type Config struct {
	ApiHost string
	ApiPort int
}

var (
	config Config
	once   sync.Once
)

func Get() *Config {
	once.Do(func() {
		var host string
		var portEnv string
		err := godotenv.Load()
		if err != nil {
			log.Printf("Error loading .env file, default config")
		}
		host = os.Getenv("API_HOST") // "ab" os.Getenv("MQTT_USER")
		portEnv = os.Getenv("API_PORT")

		port, err := strconv.Atoi(portEnv)
		if err != nil {
			port = 8000
		}

		config.ApiHost = host
		config.ApiPort = port

		configBytes, err := json.MarshalIndent(config, "", "  ")
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("Config:", string(configBytes))
	})
	return &config
}
