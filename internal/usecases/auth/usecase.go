package auth

import (
	"context"
	"github.com/movie-app-crud-gorm/internal/config"
	"github.com/movie-app-crud-gorm/internal/domain"
	jwtutil "github.com/movie-app-crud-gorm/internal/pkg/jwt"
	"github.com/movie-app-crud-gorm/internal/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"time"
)

type UseCase struct {
	userRepo domain.UserRepository
	log      logger.Logger
	cfg      config.Config
}

func New(
	userRepo domain.UserRepository,
	log logger.Logger,
	cfg config.Config) *UseCase {
	return &UseCase{
		userRepo: userRepo,
		log:      log,
		cfg:      cfg,
	}
}

func (uc *UseCase) SignUp(ctx context.Context, req domain.User) (string, error) {
	var logMsg = " uc.user.SignUp "
	hashed, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		uc.log.Error(logMsg+"bcrypt.GenerateFromPassword failed", logger.Error(err))
		return "", err
	}

	userID, err := uc.userRepo.CreateUser(ctx, domain.User{
		Email:    req.Email,
		Password: string(hashed),
	})
	if err != nil {
		uc.log.Error(logMsg+"uc.userRepo.CreateUser failed", logger.Error(err))
		return "", err
	}

	return userID, nil
}

func (uc *UseCase) Login(ctx context.Context, req domain.User) (domain.JwtTokens, error) {
	var logMsg = " uc.user.Login "

	user, err := uc.userRepo.GetByEmail(ctx, req.Email)
	if err != nil {
		uc.log.Error(logMsg+"uc.userRepo.GetByEmail failed", logger.Error(err))
		return domain.JwtTokens{}, err
	}

	if errG := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); errG != nil {
		uc.log.Error(logMsg+"bcrypt.CompareHashAndPassword failed", logger.Error(errG))
		return domain.JwtTokens{}, errG
	}

	accessTTL := time.Duration(uc.cfg.Jwt.AccessTokenTTL) * time.Minute
	refreshTTL := time.Duration(uc.cfg.Jwt.RefreshTokenTTL) * time.Hour

	access, errA := jwtutil.GenerateToken(user.ID, uc.cfg.Jwt.SecretKey, accessTTL)
	if errA != nil {
		uc.log.Error(logMsg+"jwtutil.GenerateToken ACT failed", logger.Error(errA))
		return domain.JwtTokens{}, errA
	}

	refresh, errR := jwtutil.GenerateToken(user.ID, uc.cfg.Jwt.SecretKey, refreshTTL)
	if errR != nil {
		uc.log.Error(logMsg+"jwtutil.GenerateToken RT failed", logger.Error(errR))
		return domain.JwtTokens{}, errR
	}

	return domain.JwtTokens{
		AccessToken:  access,
		RefreshToken: refresh,
	}, nil
}
