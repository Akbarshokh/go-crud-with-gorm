package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/movie-app-crud-gorm/internal/domain"
	"github.com/movie-app-crud-gorm/internal/pkg/status"
	"net/http"
)

type AuthHandler struct {
	authUC domain.AuthUseCase
}

func NewAuthHandler(authUC domain.AuthUseCase) *AuthHandler {
	return &AuthHandler{authUC: authUC}
}

type SignUp struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required,min=6"`
}

type LoginRequest struct {
	Email    string `json:"email" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type JwtTokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

type userIdResponse struct {
	ID string `json:"id"`
}

// SignUp godoc
// @Summary Sign Up
// @Router /sign-up [POST]
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body SignUp true "User Data"
// @Success 200 {object} R{data=userIdResponse}
// @Failure 422 {object} R{data=userIdResponse}
// @Failure 500 {object} R{data=userIdResponse}
func (h *AuthHandler) SignUp(c *gin.Context) {
	var req SignUp
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, R{
			Status:    status.Failure,
			ErrorCode: status.ErrorBadRequest,
			ErrorNote: err.Error(),
		})
		return
	}

	userID, err := h.authUC.SignUp(c.Request.Context(), domain.User{Email: req.Email, Password: req.Password})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errView(status.ErrorInternalServerError, err.Error()))
		return
	}

	c.JSON(http.StatusOK, view(userIdResponse{
		ID: userID,
	}))
}

// Login godoc
// @Summary Log in
// @Router /login [POST]
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body LoginRequest true "User Data"
// @Success 200 {object} R{data=JwtTokens}
// @Failure 422 {object} R{data=JwtTokens}
// @Failure 500 {object} R{data=JwtTokens}
func (h *AuthHandler) Login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, R{
			Status:    status.Failure,
			ErrorCode: status.ErrorBadRequest,
			ErrorNote: err.Error(),
		})
		return
	}
	tokens, err := h.authUC.Login(c.Request.Context(), domain.User{Email: req.Email, Password: req.Password})
	if err != nil {
		c.JSON(http.StatusUnauthorized, errView(status.ErrorUnauthorizedAccess, err.Error()))
		return
	}

	c.JSON(http.StatusOK, view(JwtTokens{
		AccessToken:  tokens.AccessToken,
		RefreshToken: tokens.RefreshToken,
	}))
}
