package api

import (
	"context"
	"fmt"
	"log"

	"github.com/eko/gocache/lib/v4/cache"
	"github.com/quessapp/core-go/configs"
	"github.com/quessapp/core-go/docs"
	"github.com/quessapp/core-go/internal/auth"
	"github.com/quessapp/core-go/internal/blocks"
	healthcheck "github.com/quessapp/core-go/internal/health-check"
	"github.com/quessapp/core-go/internal/middlewares"
	"github.com/quessapp/core-go/internal/questions"
	"github.com/quessapp/core-go/internal/queues"
	"github.com/quessapp/core-go/internal/reports"
	"github.com/quessapp/core-go/internal/settings"
	"github.com/quessapp/core-go/internal/users"

	"github.com/quessapp/toolkit/database"
	"github.com/quessapp/toolkit/queue"
	"github.com/quessapp/toolkit/s3"
	"github.com/redis/go-redis/v9"

	AWS_S3 "github.com/aws/aws-sdk-go/service/s3"
	"github.com/gofiber/fiber/v2"

	"github.com/streadway/amqp"
	"go.mongodb.org/mongo-driver/mongo"

	redis_store "github.com/eko/gocache/store/redis/v4"
)

func loadConfig() *configs.Conf {
	config, err := configs.LoadConfig(".")

	if err != nil {
		log.Fatalf("failed to load config: %s", err)
	}

	return config
}

func initDatabase(cfg *configs.Conf) *mongo.Database {
	connURI := fmt.Sprintf("%s:%s", cfg.DB.Host, cfg.DB.Port)

	db, err := database.Connect(connURI, cfg.DB.Name)

	if err != nil {
		log.Fatalf("failed to connect to database: %s", err)
	}

	return db
}

func initCache(cfg *configs.Conf) *configs.Cache {
	redisStore := redis_store.NewRedis(redis.NewClient(&redis.Options{
		Addr: cfg.Cache.URI,
	}))

	return cache.New[string](redisStore)
}

func initMessageBroker(cfg *configs.Conf) (*amqp.Connection, *amqp.Channel) {
	return queue.Connect(cfg.Queue.URI)
}

func initEmailsQueue(ch *amqp.Channel, queueName string) *amqp.Queue {
	q, err := queues.DeclareQueue(ch, queueName)

	if err != nil {
		log.Fatalf("failed to declare emails queue: %s", err)
	}

	return q
}

func initTrustedIPsQueue(ch *amqp.Channel, queueName string) *amqp.Queue {
	q, err := queues.DeclareQueue(ch, queueName)

	if err != nil {
		log.Fatalf("failed to declare trusted IPs queue: %s", err)
	}

	return q
}

func initS3(cfg *configs.Conf) *AWS_S3.S3 {
	S3Client, err := s3.Configure(&cfg.S3.Region, &s3.S3Credentials{
		AccessKey: cfg.S3.AccessKey,
		Secret:    cfg.S3.Secret,
		Token:     cfg.S3.Token,
	})

	if err != nil {
		log.Fatalf("failed to configure S3 client: %s", err)
	}

	return S3Client
}

func initRepositories(db *mongo.Database) (*auth.AuthRepository, *users.UsersRepository, *questions.QuestionsRepository, *blocks.BlocksRepository, *reports.ReportsRepository) {
	return auth.NewAuthRepository(db), users.NewRepository(db), questions.NewRepository(db), blocks.NewRepository(db), reports.NewRepository(db)
}

func initRoutes(appCtx *configs.AppCtx, authRepository *auth.AuthRepository, usersRepository *users.UsersRepository, questionsRepository *questions.QuestionsRepository, blocksRepository *blocks.BlocksRepository, reportsRepository *reports.ReportsRepository) {
	auth.LoadRoutes(appCtx, authRepository, usersRepository)
	questions.LoadRoutes(appCtx, usersRepository, questionsRepository, blocksRepository)
	blocks.LoadRoutes(appCtx, usersRepository, blocksRepository)
	users.LoadRoutes(appCtx, usersRepository)
	healthcheck.LoadRoutes(appCtx, authRepository, questionsRepository, usersRepository)
	settings.LoadRoutes(appCtx, usersRepository)
	reports.LoadRoutes(appCtx, questionsRepository, usersRepository, reportsRepository)
	docs.LoadRoutes(appCtx)
}

func initServer(cfg *configs.Conf, messageBrokerChannel *amqp.Channel, S3Client *AWS_S3.S3, db *mongo.Database) {
	app := fiber.New()

	AppCtx := &configs.AppCtx{
		App:             app,
		DB:              db,
		Cfg:             cfg,
		MessageQueueCh:  messageBrokerChannel,
		S3Client:        S3Client,
		EmailsQueue:     initEmailsQueue(messageBrokerChannel, cfg.Queue.SendEmailsQueueName),
		TrustedIPsQueue: initTrustedIPsQueue(messageBrokerChannel, cfg.Queue.CheckTrustedIPsQueueName),
		Cache:           initCache(cfg),
	}

	middlewares.ApplyMiddlewares(AppCtx.App, AppCtx.Cfg)

	authRepository, usersRepository, questionsRepository, blocksRepository, reportsRepository := initRepositories(db)
	initRoutes(AppCtx, authRepository, usersRepository, questionsRepository, blocksRepository, reportsRepository)

	log.Fatal(AppCtx.App.Listen(AppCtx.Cfg.App.ServerPort))
}

// Setup inits the application by loading the configuration, connecting to the database and message broker, and initializing the S3 client.
// It first calls the loadConfig function to read the configuration file and store it in the variable 'cfg'.
// Then, it uses the initDatabase and initMessageBroker functions to connect to the database and message broker, respectively, using the configuration stored in 'cfg'.
// It also calls the initS3 function to init the S3 client, which will be used to upload and download files.
// Finally, it calls the initServer function to start the HTTP server, passing the initd database, message broker, and S3 client as parameters.
func Setup() {
	cfg := loadConfig()
	db := initDatabase(cfg)

	defer func() {
		if err := db.Client().Disconnect(context.Background()); err != nil {
			panic(err)
		}
	}()

	conn, ch := initMessageBroker(cfg)
	defer conn.Close()

	s3 := initS3(cfg)

	initServer(cfg, ch, s3, db)
}
