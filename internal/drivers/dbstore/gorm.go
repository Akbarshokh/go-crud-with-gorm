package dbstore

import (
	"fmt"
	"github.com/movie-app-crud-gorm/internal/config"
	movies_repo "github.com/movie-app-crud-gorm/internal/drivers/dbstore/movies"
	"github.com/movie-app-crud-gorm/internal/drivers/dbstore/user"
	"github.com/movie-app-crud-gorm/internal/pkg/logger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewGormDB(cfg config.Config, log logger.Logger) (*gorm.DB, error) {
	var logMsg = "gorm.Init"
	dsn := cfg.Postgres.DSN()

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Error(logMsg+"gorm.Open: Failed to connect to DB", logger.Error(err))
		return nil, fmt.Errorf("gorm.Open: %w", err)
	}

	err = db.AutoMigrate(
		&user.Users{},
		&movies_repo.Movies{},
	)
	if err != nil {
		log.Error(logMsg+"gorm.AutoMigrate: Failed to migrate DB", logger.Error(err))
		return nil, fmt.Errorf("auto migration failed: %w", err)
	}

	log.Info("GORM Connected to DB")

	return db, nil
}
