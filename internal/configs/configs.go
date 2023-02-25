package configs

import (
	"log"

	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

// Conf is a model for app config. Like the app name, app port.
// Also it can initialize DB configs, JWT, etc.
type Conf struct {
	// App Config
	APPName    string `mapstructure:"APP_NAME"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	Env        string `mapstructure:"ENV"`
	APIKey     string `mapstructure:"API_KEY"`

	// Database Config
	DBDriver   string `mapstructure:"DB_DRIVER"`
	DBHost     string `mapstructure:"DB_HOST"`
	DBPort     string `mapstructure:"DB_PORT"`
	DBUser     string `mapstructure:"DB_USER"`
	DBPassword string `mapstructure:"DB_PASSWORD"`
	DBName     string `mapstructure:"DB_NAME"`

	// CORS
	CORSOrigins string `mapstructure:"ALLOW_ORIGINS"`

	// JWT config
	JWTSecret    string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn int    `mapstructure:"JWT_EXPIRES_IN"`
	JWTToken     *jwtauth.JWTAuth

	// Queues
	MessageQueueURI string `mapstructure:"MessageQueueURI"`
}

var cfg *Conf

// LoadConfig loads config from .env. It handles JWT config, db config, server config, etc.
func LoadConfig(path string) (*Conf, error) {
	viper.SetConfigName("app_config")
	viper.SetConfigType("env")
	viper.AddConfigPath(path)
	viper.SetConfigFile(".env")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		panic(err)
	}

	log.Printf("======= %s ======= \n", cfg.APPName)
	log.Printf("PORT: %s", cfg.ServerPort)
	log.Printf("ENV: %s", cfg.Env)
	log.Println("===============================")

	cfg.JWTToken = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	return cfg, nil
}
