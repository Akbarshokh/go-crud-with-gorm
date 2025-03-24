package domain

import "context"

// Movie — чистая бизнес-сущность (без GORM, JSON-тегов)
type Movie struct {
	ID       uint
	Title    string
	Director string
	Year     int
	Plot     string
}

// MovieRepository — интерфейс, реализуемый в drivers/dbstore/movie,
// используется в usecase и может быть замокан для тестов
type MovieRepository interface {
	Create(ctx context.Context, request Movie) (uint, error)
	GetAll(ctx context.Context) ([]Movie, error)
	GetByID(ctx context.Context, id uint) (Movie, error)
	Update(ctx context.Context, request Movie) error
	Delete(ctx context.Context, id uint) error
}

// MovieUseCase — интерфейс бизнес-логики, реализуется в usecases/movies
type MovieUseCase interface {
	Create(ctx context.Context, movie Movie) (uint, error)
	GetAll(ctx context.Context) ([]Movie, error)
	GetByID(ctx context.Context, id uint) (Movie, error)
	Update(ctx context.Context, movie Movie) error
	Delete(ctx context.Context, id uint) error
}
