package api

import (
	"net/http"
	"shortLink/internal/service"

	"github.com/gin-gonic/gin"
)

// ShortenRequest represents the request body for creating a short link.
type ShortenRequest struct {
	URL string `json:"url" binding:"required"`
}

// ShortenResponse represents a successful short-link creation response.
type ShortenResponse struct {
	ShortURL string `json:"short_url"`
}

// ErrorResponse describes an error payload returned by the API.
type ErrorResponse struct {
	Error string `json:"error"`
}

func RegisterShortLinkRoutes(r *gin.Engine) {
	r.POST("/api/shorten", createShortLink)
	r.GET("/:code", redirect)
}

// createShortLink handles creation of short links.
// @Summary Create short link
// @Description Accept a long URL and return the shortened version.
// @Tags shortlink
// @Accept json
// @Produce json
// @Param data body ShortenRequest true "Request body"
// @Success 200 {object} ShortenResponse
// @Failure 400 {object} ErrorResponse
// @Failure 500 {object} ErrorResponse
// @Router /api/shorten [post]
func createShortLink(c *gin.Context) {
	var req ShortenRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, ErrorResponse{Error: err.Error()})
		return
	}
	code, err := service.CreateShortLink(req.URL)
	if err != nil {
		c.JSON(http.StatusInternalServerError, ErrorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, ShortenResponse{ShortURL: "http://localhost:8080/" + code})
}

// redirect handles jumping to the original URL by short code.
// @Summary Redirect to original URL
// @Description Redirect the client to the original URL identified by the short code.
// @Tags shortlink
// @Param code path string true "Short code"
// @Success 302 {string} string "Redirect"
// @Failure 404 {object} ErrorResponse
// @Router /{code} [get]
func redirect(c *gin.Context) {
	code := c.Param("code")
	url, err := service.GetOriginalURL(code)
	if err != nil {
		c.JSON(http.StatusNotFound, ErrorResponse{Error: "link not found"})
		return
	}

	c.Redirect(http.StatusFound, url)
}
