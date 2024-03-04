package main

import (
	"flag"
	"os"
	"time"

	"mlops/db"

	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type config struct {
	port string
	env  string
	db   struct {
		dsn string
	}
}
type application struct {
	config
	logger *zap.Logger
	router *chi.Mux
	sqlc   *db.Queries //sqlc generated queries
}

func main() {
	var cfg config
	flag.StringVar(&cfg.port, "port", "0.0.0.0:4000", "API server port")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("DB_DSN"), "PostgreSQL DSN")

	flag.Parse()

	config := zap.NewProductionConfig()

	config.EncoderConfig.EncodeTime = zapcore.TimeEncoderOfLayout(time.RFC3339)

	//logger, _ := config.Build()
	logger := logInit(true)

	logger.Info("Log writer has started successfully.") 
	dbPool, err := openDB(cfg)

	if err != nil {
		logger.Fatal("Failed to open DB", zap.Error(err))
	}

	defer dbPool.Close()
	defer logger.Sync()
	sugar := logger.Sugar()

	if err != nil {
		logger.Fatal("Failed to create mailer", zap.Error(err))
	}

	app := &application{
		config: cfg,
		sqlc:   db.New(dbPool),
		router: chi.NewRouter(),
		logger: logger,
	}
	sugar.Info("Database connection estabilished")
	app.Routes()
	if err := app.serve(); err != nil {
		logger.Fatal("Server failed to start", zap.Error(err))
	}
}
