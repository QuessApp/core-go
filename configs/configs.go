package configs

import (
	"log"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

// Cache is a type alias for cache.Cache[string].
type Cache = redis.Client

// AppConfig holds the application configuration.
type AppConfig struct {
	APPName    string `mapstructure:"APP_NAME"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	ServerHost string `mapstructure:"SERVER_HOST"`
	Env        string `mapstructure:"ENV"`
	APIKey     string `mapstructure:"API_KEY"`
}

// DBConfig holds the database configuration.
type DBConfig struct {
	Driver   string `mapstructure:"DB_DRIVER"`
	Host     string `mapstructure:"DB_HOST"`
	Port     string `mapstructure:"DB_PORT"`
	User     string `mapstructure:"DB_USER"`
	Password string `mapstructure:"DB_PASSWORD"`
	Name     string `mapstructure:"DB_NAME"`
}

// CORSConfig holds the CORS configuration.
type CORSConfig struct {
	Origins string `mapstructure:"ALLOW_ORIGINS"`
}

type JWTConfig struct {
	Secret    string `mapstructure:"JWT_SECRET"`
	ExpiresIn int    `mapstructure:"JWT_EXPIRES_IN"`
}

// QueueConfig holds the message broker configuration.
type QueueConfig struct {
	URI                      string `mapstructure:"MESSAGE_BROKER_URI"`
	SendEmailsQueueName      string `mapstructure:"SEND_EMAILS_QUEUE_NAME"`
	CheckTrustedIPsQueueName string `mapstructure:"CHECK_TRUSTED_IPS_QUEUE_NAME"`
}

// CryptoConfig holds the crypto configuration.
type CryptoConfig struct {
	Key string `mapstructure:"CIPHER_KEY"`
}

// S3Config holds the S3 configuration.
type S3Config struct {
	Region     string `mapstructure:"S3_REGION"`
	BucketName string `mapstructure:"S3_BUCKET_NAME"`
	AccessKey  string `mapstructure:"S3_ACCESS_KEY"`
	Secret     string `mapstructure:"S3_SECRET"`
	Token      string `mapstructure:"S3_TOKEN"`
}

// CDNConfig holds the CDN configuration.
type CDNConfig struct {
	URI string `mapstructure:"CDN_URI"`
}

// CacheConfig holds the cache configuration.
type CacheConfig struct {
	URI string `mapstructure:"CACHE_URI"`
}

// Conf is a model for app config. Like the app name, app port.
// Also it can initialize DB configs, JWT, etc.
type Conf struct {
	App    AppConfig    `mapstructure:",squash"`
	DB     DBConfig     `mapstructure:",squash"`
	CORS   CORSConfig   `mapstructure:",squash"`
	JWT    JWTConfig    `mapstructure:",squash"`
	Queue  QueueConfig  `mapstructure:",squash"`
	Crypto CryptoConfig `mapstructure:",squash"`
	S3     S3Config     `mapstructure:",squash"`
	CDN    CDNConfig    `mapstructure:",squash"`
	Cache  CacheConfig  `mapstructure:",squash"`
}

var cfg *Conf

// AppCtx is a global model for app. It defines the router, db, config, repositories, etc.
// Use AppCtx to avoid long function params.
type AppCtx struct {
	App             *fiber.App
	DB              *mongo.Database
	Cfg             *Conf
	MessageQueueCh  *amqp.Channel
	EmailsQueue     *amqp.Queue
	TrustedIPsQueue *amqp.Queue
	S3Client        *s3.S3
	Cache           *Cache
}

// HandlersCtx is a global model for handlers. It defines the fiber context, app context, etc.
// Use HandlersCtx to avoid long function params.
type HandlersCtx struct {
	// Context from fiber.
	C *fiber.Ctx
	// App config.
	AppCtx
}

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

	log.Printf("======= %s ======= \n", cfg.App.APPName)
	log.Printf("PORT: %s", cfg.App.ServerPort)
	log.Printf("ENV: %s", cfg.App.Env)
	log.Println("===============================")

	return cfg, nil
}
