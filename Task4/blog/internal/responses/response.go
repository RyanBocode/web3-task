package responses

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message,omitempty"`
}

func JSONError(c *gin.Context, status int, err string, msg string) {
	c.AbortWithStatusJSON(status, ErrorResponse{Error: err, Message: msg})
}

func JSONOK(c *gin.Context, data any) {
	c.JSON(http.StatusOK, data)
}
