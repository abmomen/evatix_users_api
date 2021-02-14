package ping

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//Ping responds to the /ping(GET)
func Ping(c *gin.Context) {
	c.String(http.StatusOK, "pong")
}
