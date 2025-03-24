package movies_repo

import (
	"context"
	"github.com/movie-app-crud-gorm/internal/domain"
	"github.com/movie-app-crud-gorm/internal/errs"
	"github.com/movie-app-crud-gorm/internal/pkg/logger"
	"gorm.io/gorm"
)

type Repo struct {
	db  *gorm.DB
	log logger.Logger
}

func New(db *gorm.DB, logger logger.Logger) *Repo {
	return &Repo{
		db:  db,
		log: logger,
	}
}

func (repo *Repo) Create(ctx context.Context, request domain.Movie) (uint, error) {
	var logMsg = "repo.Movies.Create "

	data := Movies{
		Model:    gorm.Model{ID: request.ID},
		Title:    request.Title,
		Director: request.Director,
		Year:     request.Year,
		Plot:     request.Plot,
	}

	err := repo.db.WithContext(ctx).Create(&data).Error
	if err != nil {
		repo.log.Error(logMsg+" repo.db.Create failed", logger.Error(err))
		return 0, err
	}

	return data.ID, err
}

func (repo *Repo) GetAll(ctx context.Context) ([]domain.Movie, error) {
	var (
		logMsg   = "repo.Movies.GetAll "
		gormData []Movies
		result   []domain.Movie
	)

	if err := repo.db.WithContext(ctx).Find(&gormData).Error; err != nil {
		repo.log.Error(logMsg+" repo.db.GetAll failed", logger.Error(err))
		return nil, err
	}

	for _, val := range gormData {
		result = append(result, ToDomain(val))
	}

	if len(result) == 0 {
		repo.log.Error(logMsg + "no records found")
		return nil, errs.ErrNotFound
	}

	return result, nil
}

func (repo *Repo) GetByID(ctx context.Context, id uint) (domain.Movie, error) {
	var (
		logMsg   = "repo.Movies.GetByID "
		gormData Movies
	)

	if err := repo.db.WithContext(ctx).First(&gormData, id).Error; err != nil {
		repo.log.Error(logMsg+" repo.db.GetByID failed", logger.Error(err))
		return domain.Movie{}, err
	}

	return ToDomain(gormData), nil
}

func (repo *Repo) Update(ctx context.Context, request domain.Movie) error {
	var logMsg = "repo.Movies.Update "

	err := repo.db.WithContext(ctx).Save(&request).Error
	if err != nil {
		repo.log.Error(logMsg+" repo.db.Save failed", logger.Error(err))
		return err
	}

	return nil
}

func (repo *Repo) Delete(ctx context.Context, id uint) error {
	var logMsg = "repo.Movies.Delete "

	err := repo.db.WithContext(ctx).Delete(&Movies{}, id).Error
	if err != nil {
		repo.log.Error(logMsg+" repo.db.Delete failed", logger.Error(err))
		return err
	}
	return nil
}
