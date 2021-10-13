package app

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
	"github.com/spf13/viper"
	"github.com/sw0ps/monitoring-app/internal/handler"
	"github.com/sw0ps/monitoring-app/internal/repository"
	"github.com/sw0ps/monitoring-app/internal/repository/postgres"
	"github.com/sw0ps/monitoring-app/internal/service"
	"log"
	"os"
)

func Run() {

	if err := godotenv.Load(); err != nil {
		log.Fatalf("error loading env variables: %s", err.Error())
	}

	if err := initConfig(); err != nil {
		log.Fatalf("error initializing configs: %s", err.Error())
	}

	db := DbConnect()

	repos := repository.NewRepository(db)
	services := service.NewService(repos)
	handlers := handler.NewHandler(services)

	app := handlers.InitRoutes(fiber.New())

	if err := handlers.FirstRun(); err != nil {
		log.Fatalf("error add data with first run: %s", err.Error())
	}

	log.Fatal(app.Listen(":" + viper.GetString("port")))
}

func DbConnect() *sqlx.DB {

	db, err := postgres.NewPostgresDB(postgres.Config{
		Host:     viper.GetString("db.host"),
		Port:     viper.GetString("db.port"),
		Username: viper.GetString("db.username"),
		DBName:   viper.GetString("db.dbname"),
		SSLMode:  viper.GetString("db.sslmode"),
		Password: os.Getenv("DB_PASSWORD"),
	})

	if err != nil {
		log.Fatalf("failed to initialize db: %s", err.Error())
	}

	return db
}

func initConfig() error {

	viper.AddConfigPath("configs")
	viper.SetConfigName("main")

	return viper.ReadInConfig()
}
