package database

import (
	"log"
	"os"
	"rsvp-system/config"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// SetupDatabase initializes the database connection
func SetupDatabase() (*gorm.DB, error) {
	dsn := config.LoadDatabaseConfig()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             200 * time.Millisecond, // log queries slower than this
			LogLevel:                  logger.Info,            // Log all queries (Info = all, Warn = slow queries, Error = only errors)
			IgnoreRecordNotFoundError: true,                   // ignore ErrRecordNotFound
			Colorful:                  true,                   // colorful output
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		return nil, err
	}

	return db, nil
}
