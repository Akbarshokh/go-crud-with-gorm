package movies

import (
	"github.com/movie-app-crud-gorm/internal/domain"
	"gorm.io/gorm"
)

type movie struct {
	gorm.Model
	Title    string `json:"title" binding:"required"`
	Director string `json:"director" binding:"required"`
	Year     int    `json:"year" binding:"required"`
	Plot     string `json:"plot" binding:"required"`
}

func FromDomain(m domain.Movie) movie {
	return movie{
		Model:    gorm.Model{ID: m.ID},
		Title:    m.Title,
		Director: m.Director,
		Year:     m.Year,
		Plot:     m.Plot,
	}
}

func ToDomain(m movie) domain.Movie {
	return domain.Movie{
		ID:       m.ID,
		Title:    m.Title,
		Director: m.Director,
		Year:     m.Year,
		Plot:     m.Plot,
	}
}
