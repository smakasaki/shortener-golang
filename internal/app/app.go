package app

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	userdelivery "github.com/smakasaki/typing-trainer/internal/user/delivery"
	userrepo "github.com/smakasaki/typing-trainer/internal/user/repository"
	useruc "github.com/smakasaki/typing-trainer/internal/user/usecase"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type App struct {
	Port   string
	DB     *sql.DB
	Echo   *echo.Echo
	Logger *zap.Logger
	Env    string
}

func New() (*App, error) {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "production"
		log.Println("APP_ENV environment variable not set, defaulting to production")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Println("PORT environment variable not set, defaulting to 8080")
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable not set")
	}

	dbConn, err := connectDB(dbURL)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %v", err)
	}

	logger, err := newLogger(env)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %v", err)
	}

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(middleware.RequestLoggerWithConfig(middleware.RequestLoggerConfig{
		LogURI:     true,
		LogStatus:  true,
		LogMethod:  true,
		LogLatency: true,
		LogError:   true,
		LogValuesFunc: func(c echo.Context, v middleware.RequestLoggerValues) error {
			logger.Info("Request",
				zap.String("method", v.Method),
				zap.String("uri", v.URI),
				zap.Int("status", v.Status),
				zap.Duration("latency", v.Latency),
			)
			return nil
		},
	}))

	app := &App{
		Port:   port,
		DB:     dbConn,
		Echo:   e,
		Logger: logger,
		Env:    env,
	}

	return app, nil
}

func newLogger(env string) (*zap.Logger, error) {
	var cfg zap.Config
	if env == "development" {
		cfg = zap.NewDevelopmentConfig()
		cfg.EncoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
	} else {
		cfg = zap.NewProductionConfig()
	}

	return cfg.Build()
}

func connectDB(url string) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		return nil, fmt.Errorf("unable to open database connection: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	log.Println("Successfully connected to the database")
	return db, nil
}

func (a *App) Run() error {
	a.Logger.Info("Starting server", zap.String("port", a.Port), zap.String("env", a.Env))

	userRepo := userrepo.NewUserRepository(a.DB)
	userUseCase := useruc.NewUserUseCase(userRepo)
	userdelivery.RegisterEndpoints(a.Echo, userUseCase)

	a.Echo.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	return a.Echo.Start(":" + a.Port)
}
