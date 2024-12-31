package config

import (
	"os"

	"github.com/sbrenomartins/gopportunities/schemas"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func InitializeSQLite() (*gorm.DB, error) {
	logger := GetLogger("sqlite")
	dbPath := "./db/gopportunities.db"

	_, err := os.Stat(dbPath)

	if os.IsNotExist(err) {
		logger.Info("Creating SQLite database...")
		err = os.MkdirAll("./db", os.ModePerm)

		if err != nil {
			logger.Errorf("Error creating directory: %v", err)
			return nil, err
		}

		file, err := os.Create(dbPath)

		if err != nil {
			logger.Errorf("Error creating SQLite file: %v", err)
			return nil, err
		}

		file.Close()
	}

	db, err :=gorm.Open(sqlite.Open(dbPath), &gorm.Config{})

	if err != nil {
		logger.Errorf("Error initializing SQLite: %v", err)
		return nil, err
	}

	err = db.AutoMigrate(&schemas.Opening{})

	if err != nil {
		logger.Errorf("Error migrating SQLite: %v", err)
		return nil, err
	}

	return db, nil
}