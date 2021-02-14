package app

import (
	"github.com/abmomin/users_api/controllers/ping"
	"github.com/abmomin/users_api/controllers/users"
)

func mapUrls() {
	router.GET("/ping", ping.Ping)

	router.POST("/users", users.Create)
	router.GET("/users/:userId", users.Get)
	router.PUT("/users/:userId", users.Update)
	router.PATCH("/users/:userId", users.Update)
	router.DELETE("/users/:userId", users.Delete)
	router.GET("/internal/users/search", users.Search)
}
