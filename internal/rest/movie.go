package rest

import (
	"github.com/gin-gonic/gin"
	"github.com/movie-app-crud-gorm/internal/domain"
	"github.com/movie-app-crud-gorm/internal/pkg/status"
	"net/http"
	"strconv"
)

type MovieHandler struct {
	uc domain.MovieUseCase
}

func NewMovieHandler(uc domain.MovieUseCase) *MovieHandler {
	return &MovieHandler{
		uc: uc,
	}
}

type MovieRequest struct {
	Title    string `json:"title" binding:"required"`
	Director string `json:"director" binding:"required"`
	Year     int    `json:"year" binding:"required"`
	Plot     string `json:"plot"`
}

type MovieResponse struct {
	ID       uint   `json:"id"`
	Title    string `json:"title"`
	Director string `json:"director"`
	Year     int    `json:"year"`
	Plot     string `json:"plot"`
}

type result struct {
	Result bool `json:"result"`
}

type idResponse struct {
	ID uint `json:"id"`
}

func toView(m domain.Movie) MovieResponse {
	return MovieResponse{
		ID:       m.ID,
		Title:    m.Title,
		Director: m.Director,
		Year:     m.Year,
		Plot:     m.Plot,
	}
}

func fromDomainToView(m []domain.Movie) []MovieResponse {
	var result = make([]MovieResponse, len(m))

	for i := range m {
		result[i] = MovieResponse{
			ID:       m[i].ID,
			Title:    m[i].Title,
			Director: m[i].Director,
			Year:     m[i].Year,
			Plot:     m[i].Plot,
		}
	}
	return result
}

// Create Movie godoc
// @Summary Create a new movie record in db
// @Router /movies [POST]
// @Tags Movie
// @Accept json
// @Produce json
// @Security Bearer
// @Param request body MovieRequest true "Movie Data"
// @Success 200 {object} R{data=idResponse}
// @Failure 422 {object} R{data=idResponse}
// @Failure 500 {object} R{data=idResponse}
func (h *MovieHandler) Create(c *gin.Context) {
	var req MovieRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, R{
			Status:    status.Failure,
			ErrorCode: status.ErrorBindFailed,
			ErrorNote: err.Error(),
		})
		return
	}

	id, err := h.uc.Create(c.Request.Context(), domain.Movie{
		Title:    req.Title,
		Director: req.Director,
		Year:     req.Year,
		Plot:     req.Plot,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errView(status.ErrorCreateFailed, err.Error()))
		return
	}

	c.JSON(http.StatusCreated, view(idResponse{ID: id}))
}

// GetAll Movies godoc
// @Summary Retrieve all movies from db
// @Router /movies [GET]
// @Tags Movie
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} R{data=[]MovieResponse}
// @Failure 422 {object} R{data=[]MovieResponse}
// @Failure 500 {object} R{data=[]MovieResponse}
func (h *MovieHandler) GetAll(c *gin.Context) {
	movies, err := h.uc.GetAll(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, errView(status.ErrorNotFound, err.Error()))
		return
	}
	c.JSON(http.StatusOK, view(fromDomainToView(movies)))
}

// GetByID Movie godoc
// @Summary Retrieve a specific movie by ID
// @Router /movies/{id}  [GET]
// @Tags Movie
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Movie ID"
// @Success 200 {object} R{data=MovieResponse}
// @Failure 422 {object} R{data=MovieResponse}
// @Failure 500 {object} R{data=MovieResponse}
func (h *MovieHandler) GetByID(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, R{
			Status:    status.Failure,
			ErrorCode: status.ErrorBadRequest,
			ErrorNote: "invalid id format",
		})
		return
	}

	if id < 1 {
		c.JSON(http.StatusBadRequest, R{
			Status:    status.Failure,
			ErrorCode: status.ErrorCodeValidation,
			ErrorNote: "id must be greater than 0",
		})
		return
	}

	movie, errG := h.uc.GetByID(c.Request.Context(), uint(id))
	if errG != nil {
		c.JSON(http.StatusNotFound, errView(status.ErrorNotFound, errG.Error()))
		return
	}

	c.JSON(http.StatusOK, view(toView(movie)))
}

// Update Movie godoc
// @Summary Update an existing movie record
// @Router /movies/{id}  [PUT]
// @Tags Movie
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Movie ID"
// @Param request body MovieRequest true "Movie Data"
// @Success 200 {object} R{data=result}
// @Failure 422 {object} R{data=result}
// @Failure 500 {object} R{data=result}
func (h *MovieHandler) Update(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, R{
			Status:    status.Failure,
			ErrorCode: status.ErrorBadRequest,
			ErrorNote: "invalid id format",
		})
		return
	}

	if id < 1 {
		c.JSON(http.StatusBadRequest, R{
			Status:    status.Failure,
			ErrorCode: status.ErrorCodeValidation,
			ErrorNote: "id must be greater than 0",
		})
		return
	}

	var req MovieRequest
	if errS := c.ShouldBindJSON(&req); errS != nil {
		c.JSON(http.StatusBadRequest, R{
			Status:    status.Failure,
			ErrorCode: status.ErrorBindFailed,
			ErrorNote: errS.Error(),
		})
		return
	}

	err = h.uc.Update(c.Request.Context(), domain.Movie{
		ID:       uint(id),
		Title:    req.Title,
		Director: req.Director,
		Year:     req.Year,
		Plot:     req.Plot,
	})
	if err != nil {
		c.JSON(http.StatusInternalServerError, errView(status.ErrorNotFound, err.Error()))
		return
	}

	c.JSON(http.StatusOK, view(result{Result: true}))
}

// Delete Movie godoc
// @Summary Delete a movie record by ID
// @Router /movies/{id}  [DELETE]
// @Tags Movie
// @Accept json
// @Produce json
// @Security Bearer
// @Param id path int true "Movie ID"
// @Success 200 {object} R{data=result}
// @Failure 422 {object} R{data=result}
// @Failure 500 {object} R{data=result}
func (h *MovieHandler) Delete(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, R{
			Status:    status.Failure,
			ErrorCode: status.ErrorBadRequest,
			ErrorNote: "invalid id format",
		})
		return
	}

	if id < 1 {
		c.JSON(http.StatusBadRequest, R{
			Status:    status.Failure,
			ErrorCode: status.ErrorCodeValidation,
			ErrorNote: "id must be greater than 0",
		})
		return
	}

	err = h.uc.Delete(c.Request.Context(), uint(id))
	if err != nil {
		c.JSON(http.StatusInternalServerError, errView(status.ErrorNotFound, err.Error()))
		return
	}

	c.JSON(http.StatusOK, view(result{Result: true}))
}
