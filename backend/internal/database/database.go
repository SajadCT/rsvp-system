package database

import (
	"log"
	"os"
	"rsvp-system/config"
	"rsvp-system/internal/models"
	"time"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func SetupDatabase() (*gorm.DB, error) {
	dsn := config.LoadDatabaseConfig()

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             200 * time.Millisecond,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  true,
		},
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: newLogger})
	if err != nil {
		return nil, err
	}

	log.Println("Running Database Migrations...")
	err = db.AutoMigrate(&models.User{}, &models.Event{}, &models.Guest{})
	if err != nil {
		log.Fatal("Migration Failed: ", err)
	}
	log.Println("Database Migrated Successfully!")

	return db, nil
}
