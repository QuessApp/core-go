package configs

import (
	"log"

	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/go-chi/jwtauth"
	"github.com/gofiber/fiber/v2"
	"github.com/spf13/viper"
	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"
)

// TODO: separate in unique struct. eg: one struct for DB Config, one for Queues, etc.

// Conf is a model for app config. Like the app name, app port.
// Also it can initialize DB configs, JWT, etc.
type Conf struct {
	// App Config
	APPName    string `mapstructure:"APP_NAME"`
	ServerPort string `mapstructure:"SERVER_PORT"`
	ServerHost string `mapstructure:"SERVER_HOST"`
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
	MessageBrokerURI    string `mapstructure:"MESSAGE_BROKER_URI"`
	SendEmailsQueueName string `mapstructure:"SEND_EMAILS_QUEUE_NAME"`

	// Crypto
	CipherKey string `mapstructure:"CIPHER_KEY"`

	// S3 Uploader
	S3Region     string `mapstructure:"S3_REGION"`
	S3BucketName string `mapstructure:"S3_BUCKET_NAME"`
	S3AccessKey  string `mapstructure:"S3_ACCESS_KEY"`
	S3Secret     string `mapstructure:"S3_SECRET"`
	S3Token      string `mapstructure:"S3_TOKEN"`

	// Cloudfront
	CDN_URI string `mapstructure:"CDN_URI"`
}

var cfg *Conf

// AppCtx is a global model for app. It defines the router, db, config, repositories, etc.
// Use AppCtx to avoid long function params.
type AppCtx struct {
	App            *fiber.App
	DB             *mongo.Database
	Cfg            *Conf
	MessageQueueCh *amqp.Channel
	EmailsQueue    *amqp.Queue
	S3Client       *s3.S3
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

	log.Printf("======= %s ======= \n", cfg.APPName)
	log.Printf("PORT: %s", cfg.ServerPort)
	log.Printf("ENV: %s", cfg.Env)
	log.Println("===============================")

	cfg.JWTToken = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)

	return cfg, nil
}
