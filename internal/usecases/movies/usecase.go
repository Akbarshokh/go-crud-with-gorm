package movies

import (
	"context"
	"github.com/movie-app-crud-gorm/internal/domain"
	"github.com/movie-app-crud-gorm/internal/pkg/logger"
)

type UseCase struct {
	movieRepo domain.MovieRepository
	log       logger.Logger
}

func New(movieRepo domain.MovieRepository, log logger.Logger) *UseCase {
	return &UseCase{
		movieRepo: movieRepo,
		log:       log,
	}
}

func (uc *UseCase) Create(ctx context.Context, request domain.Movie) (uint, error) {
	var logMsg = "uc.Movies.Create "

	id, err := uc.movieRepo.Create(ctx, request)
	if err != nil {
		uc.log.Error(logMsg+"uc.movieRepo.Create failed", logger.Error(err))
		return 0, err
	}

	return id, nil
}

func (uc *UseCase) GetAll(ctx context.Context) ([]domain.Movie, error) {
	var logMsg = "uc.Movies.GetAll "

	movies, err := uc.movieRepo.GetAll(ctx)
	if err != nil {
		uc.log.Error(logMsg+"uc.movieRepo.GetAll failed", logger.Error(err))
		return nil, err
	}
	return movies, nil
}

func (uc *UseCase) GetByID(ctx context.Context, id uint) (domain.Movie, error) {
	var logMsg = "uc.Movies.GetByID "

	movie, err := uc.movieRepo.GetByID(ctx, id)
	if err != nil {
		uc.log.Error(logMsg+"uc.movieRepo.GetByID failed", logger.Error(err))
		return domain.Movie{}, err
	}
	return movie, nil
}

func (uc *UseCase) Update(ctx context.Context, request domain.Movie) error {
	var logMsg = "uc.Movies.Update "

	err := uc.movieRepo.Update(ctx, request)
	if err != nil {
		uc.log.Error(logMsg+"uc.movieRepo.Update failed", logger.Error(err))
		return err
	}

	return nil
}

func (uc *UseCase) Delete(ctx context.Context, id uint) error {
	var logMsg = "uc.Movies.Delete "

	err := uc.movieRepo.Delete(ctx, id)
	if err != nil {
		uc.log.Error(logMsg+"uc.movieRepo.Delete failed", logger.Error(err))
		return err
	}
	return nil
}
