package users

import (
	"context"
	"go-service/internal/server"
	"net/http"

	"github.com/gin-gonic/gin"
)

func userRoutes(u UserService) server.RouteFunc {
	return func(c context.Context, r *gin.Engine) {
		r.GET("/", func(c *gin.Context) {
			r := u.Get()
			c.JSON(http.StatusOK, gin.H{"message": r})
		})
	}
}

func Routes(u UserService) []server.RouteFunc {
	var r []server.RouteFunc = []server.RouteFunc{
		userRoutes(u),
	}

	return r
}
