package config

import (
	"gorm.io/gorm"
)

var (
	db *gorm.DB
)

func Initialize() error {
	return nil
}

func GetLogger(p string) *Logger {
	logger := NewLogger(p)
	return logger
}