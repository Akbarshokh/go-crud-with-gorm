package movies_repo

import (
	"github.com/movie-app-crud-gorm/internal/domain"
	"gorm.io/gorm"
)

type Movies struct {
	gorm.Model
	Title    string `json:"title" binding:"required"`
	Director string `json:"director" binding:"required"`
	Year     int    `json:"year" binding:"required"`
	Plot     string `json:"plot" binding:"required"`
}

func ToDomain(m Movies) domain.Movie {
	return domain.Movie{
		ID:       m.ID,
		Title:    m.Title,
		Director: m.Director,
		Year:     m.Year,
		Plot:     m.Plot,
	}
}
