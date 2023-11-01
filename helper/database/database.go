package database

import (
	"fmt"
	"time"

	"gibhub.com/raytr/simple-bank/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDatabase(cfg *config.DBConfig) (*gorm.DB, error) {
	return NewGormConnect(cfg)
}

func NewGormConnect(cfg *config.DBConfig) (*gorm.DB, error) {
	dialect := postgres.Open(fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host,
		cfg.Port,
		cfg.Username,
		cfg.Password,
		cfg.DBName,
		"disable",
	))

	db, err := gorm.Open(dialect, nil)
	if err != nil {
		return nil, err
	}

	//connection Pool
	err = connPool(db, cfg)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func connPool(m *gorm.DB, cfg *config.DBConfig) error {
	sqlDB, err := m.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConnection)
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConnection)
	sqlDB.SetConnMaxIdleTime(time.Second * cfg.ConnectionMaxIdleTime)
	sqlDB.SetConnMaxLifetime(time.Second * cfg.ConnectionMaxLifeTime)

	return nil
}
