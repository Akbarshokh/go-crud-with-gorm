package user

import (
	"context"
	"github.com/google/uuid"
	"github.com/movie-app-crud-gorm/internal/domain"
	"github.com/movie-app-crud-gorm/internal/pkg/logger"
	"gorm.io/gorm"
)

type Repo struct {
	db  *gorm.DB
	log logger.Logger
}

func New(db *gorm.DB, log logger.Logger) *Repo {
	return &Repo{
		db:  db,
		log: log,
	}
}

type Users struct {
	ID       string `gorm:"primaryKey"`
	Email    string `gorm:"uniqueIndex;not null"`
	Password string `gorm:"not null"`
}

func (r *Repo) CreateUser(ctx context.Context, user domain.User) (string, error) {
	var logMsg = "repo.user.CreateUser "
	id := uuid.NewString()
	dbUser := Users{
		ID:       id,
		Email:    user.Email,
		Password: user.Password,
	}
	if err := r.db.WithContext(ctx).Create(&dbUser).Error; err != nil {
		r.log.Error(logMsg+"CreateUser failed", logger.Error(err))
		return "", err
	}
	return id, nil
}

func (r *Repo) GetByID(ctx context.Context, userID string) (domain.User, error) {
	var (
		logMsg = "repo.user.GetByID "
		user   Users
	)
	if err := r.db.WithContext(ctx).First(&user, "id = ?", userID).Error; err != nil {
		r.log.Error(logMsg+"GetByEmail failed", logger.Error(err))
		return domain.User{}, err
	}
	return domain.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}

func (r *Repo) GetByEmail(ctx context.Context, email string) (domain.User, error) {
	var (
		logMsg = "repo.user.GetByEmail "
		user   Users
	)
	if err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error; err != nil {
		r.log.Error(logMsg+"GetByEmail failed", logger.Error(err))
		return domain.User{}, err
	}
	return domain.User{
		ID:       user.ID,
		Email:    user.Email,
		Password: user.Password,
	}, nil
}
