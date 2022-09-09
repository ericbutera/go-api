// TODO:
// rate limit
// - https://github.com/uber-go/ratelimit
// - https://github.com/axiaoxin-com/ratelimiter
// Api key
//
// swagger notes: https://github.com/swaggo/swag/blob/master/README.md#declarative-comments-format

package rest

import (
	"net/http"
	// "time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

const swaggerUri = "/swagger/index.html"

type Base struct {
	Message string `json:"message"`
}

type StartResponse struct {
	Base
	Id uuid.UUID `json:"id"`
}

type ErrorResponse struct {
	Base
	Code int
}

type HealthResponse struct {
	Base
}

func (app *App) Docs(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, swaggerUri)
}

// Health check
// @Summary Health check
// @Schemes
// @Description Health check
// @Accept json
// @Produce json
// @Success 200 {object} HealthResponse
// @Router /health [get]
func (app *App) Health(c *gin.Context) {
	c.JSON(http.StatusOK, HealthResponse{
		Base: Base{Message: "alive"},
	})
}

// Wait for 10 seconds
// @Summary Wait for 10 seconds
// @Schemes
// @Description Wait for 10 seconds
// @Accept json
// @Produce json
// @Success 200 {Base} Base
// @Router /wait [get]
func (app *App) Wait(c *gin.Context) {
	// time.Sleep(10 * time.Second)
	c.JSON(http.StatusOK, Base{
		Message: "Waited 10 seconds",
	})
}
