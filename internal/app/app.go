package app

import (
	"context"
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/smakasaki/shortener/internal/session"
	"github.com/smakasaki/shortener/internal/url"
	"github.com/smakasaki/shortener/internal/user"
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
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	dbURL := os.Getenv("DATABASE_URL")
	if dbURL == "" {
		return nil, fmt.Errorf("DATABASE_URL environment variable not set")
	}

	logger, err := newLogger(env)
	if err != nil {
		return nil, fmt.Errorf("failed to initialize logger: %v", err)
	}

	dbConn, err := connectDB(dbURL, logger)
	if err != nil {
		logger.Error("Failed to connect to database", zap.Error(err))
		return nil, fmt.Errorf("failed to connect to database: %v", err)
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
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowCredentials: true,
		AllowMethods:     []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	app := &App{
		Port:   port,
		DB:     dbConn,
		Echo:   e,
		Logger: logger,
		Env:    env,
	}

	logger.Info("App initialized", zap.String("env", env), zap.String("port", port))
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

	logger, err := cfg.Build()
	if err != nil {
		return nil, fmt.Errorf("failed to build logger: %w", err)
	}
	return logger, nil
}

func connectDB(url string, logger *zap.Logger) (*sql.DB, error) {
	db, err := sql.Open("postgres", url)
	if err != nil {
		logger.Error("Unable to open database connection", zap.Error(err))
		return nil, fmt.Errorf("unable to open database connection: %v", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	err = db.PingContext(ctx)
	if err != nil {
		logger.Error("Unable to connect to database", zap.Error(err))
		return nil, fmt.Errorf("unable to connect to database: %v", err)
	}

	logger.Info("Successfully connected to the database")
	return db, nil
}

func (a *App) Run() error {
	a.Logger.Info("Starting server", zap.String("port", a.Port), zap.String("env", a.Env))
	// Initialize session module
	userRepo := user.NewRepository(a.DB)
	sessionRepo := session.NewRepository(a.DB)
	sessionUseCase := session.NewUseCase(sessionRepo, userRepo)
	authMiddleware := session.NewAuthMiddleware(sessionRepo, userRepo)
	session.RegisterEndpoints(a.Echo, sessionUseCase, authMiddleware)

	// Initialize user module
	userUseCase := user.NewUseCase(userRepo)
	user.RegisterEndpoints(a.Echo, userUseCase, authMiddleware)

	// Initialize URL module
	urlRepo := url.NewRepository(a.DB)
	urlUseCase := url.NewUseCase(urlRepo)
	url.RegisterEndpoints(a.Echo, urlUseCase, authMiddleware)

	a.Echo.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	return a.Echo.Start(":" + a.Port)
}
