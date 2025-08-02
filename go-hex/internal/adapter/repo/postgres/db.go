package postgres

import (
	"go-hex/internal/config"
	"log/slog"
	"os"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func setup(db *gorm.DB, env string) error {
	_db, err := db.DB()
	if err != nil {
		return err
	}

	if env == config.EnvStaging {
		_db.SetMaxIdleConns(10)
		_db.SetMaxOpenConns(100)
		_db.SetConnMaxLifetime(time.Hour)        // 0 means no limit
		_db.SetConnMaxIdleTime(time.Minute * 30) // 0 means no limit
	}

	if env == config.EnvProduction {
		_db.SetMaxIdleConns(20)                  // More idle connections
		_db.SetMaxOpenConns(200)                 // Higher max connections
		_db.SetConnMaxLifetime(time.Hour * 2)    // Longer lifetime
		_db.SetConnMaxIdleTime(time.Minute * 15) // Shorter idle time
	}

	return nil
}

func getGormConfig(env string) *gorm.Config {
	var slogger *slog.Logger
	if env == config.EnvDevelopment {
		slogger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level:     slog.LevelDebug,
			AddSource: true,
		}))

		return &gorm.Config{
			// === LOGGING CONFIGURATION ===
			Logger: logger.New(
				slog.NewLogLogger(slogger.Handler(), slog.LevelInfo),
				logger.Config{
					SlowThreshold:             time.Second, // Log queries slower than 1s
					LogLevel:                  logger.Warn, // Log level (Silent/Error/Warn/Info)
					IgnoreRecordNotFoundError: true,        // Don't log ErrRecordNotFound
					ParameterizedQueries:      false,       // Log SQL with parameters (set true in production)
					Colorful:                  true,        // Enable color (disable in production)
				},
			),

			// === PERFORMANCE SETTINGS ===
			NowFunc: func() time.Time {
				return time.Now().UTC() // Always use UTC
			},

			// === QUERY OPTIMIZATION ===
			PrepareStmt:                              true,  // Use prepared statements (better performance + security)
			DisableForeignKeyConstraintWhenMigrating: false, // Keep FK constraints

			// === MIGRATION SETTINGS ===
			DisableAutomaticPing: false, // Ping database on connect

			// === TRANSACTION SETTINGS ===
			SkipDefaultTransaction: false, // Keep transactions for safety

			// === PERFORMANCE OPTIMIZATIONS ===
			CreateBatchSize: 1000, // Batch size for bulk operations
		}
	}

	slogger = slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: false,
	})).With(
		slog.String("source", "database"),
	)
	return &gorm.Config{
		Logger: logger.New(
			slog.NewLogLogger(slogger.Handler(), slog.LevelWarn),
			logger.Config{
				SlowThreshold:             2 * time.Second, // Only log very slow queries
				LogLevel:                  logger.Error,    // Only log errors
				IgnoreRecordNotFoundError: true,
				ParameterizedQueries:      true,  // Hide parameters in logs
				Colorful:                  false, // No colors in production
			},
		),
		NowFunc: func() time.Time {
			return time.Now().UTC()
		},
		PrepareStmt:            true,
		SkipDefaultTransaction: false, // Keep for data integrity
		CreateBatchSize:        1000,
	}
}

func NewDatabase(dsn string) (*gorm.DB, error) {
	cfg, _ := config.Load()

	db, err := gorm.Open(
		postgres.New(
			postgres.Config{
				DSN:                  dsn,
				PreferSimpleProtocol: false,
			}),
		getGormConfig(cfg.Env),
	)
	if err != nil {
		return nil, err
	}

	if err := setup(db, cfg.Env); err != nil {
		return nil, err
	}

	return db, nil
}
