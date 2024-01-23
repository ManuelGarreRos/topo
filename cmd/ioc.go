package main

import (
	"TOPO/appctr"
	"TOPO/cmd/middlewares"
	"TOPO/internal/handlers"
	"TOPO/internal/repositories"
	"TOPO/internal/services"
	"github.com/gin-gonic/gin"
)

func prepareIoC(r *gin.Engine) {
	rs := prepareRepositories()
	ss := prepareServices(rs)
	hs := prepareHandlers(ss)
	prepareRouter(hs, r)
}

type rs struct {
	userRepository *repositories.UserRepository
}

type ss struct {
	userService *services.UserService
}

type hs struct {
	userHandler *handlers.UserHandler
}

func prepareRepositories() *rs {
	db := appctr.DB()
	log := appctr.Log()

	return &rs{
		userRepository: repositories.NewUserRepository(db, log),
	}
}

func prepareServices(rs *rs) *ss {
	log := appctr.Log()

	return &ss{
		userService: services.NewUserService(rs.userRepository, log),
	}
}

func prepareHandlers(ss *ss) *hs {
	log := appctr.Log()

	return &hs{
		userHandler: handlers.NewUserHandler(ss.userService, log),
	}
}

func prepareRouter(hs *hs, r *gin.Engine) {
	prepareMiddlewares(r)

	// TODO refactor this into internal/router
	usersGroup := r.Group("/users")
	{
		usersGroup.GET("/:id", hs.userHandler.ByID)
		usersGroup.POST("/create", hs.userHandler.Create)
		usersGroup.GET("/list", hs.userHandler.PaginatedList)
	}

	port := appctr.Cfg().GetString("port")
	// TODO: Add trusted proxies for production
	err := r.SetTrustedProxies([]string{
		"127.0.0.1",
	})
	if err != nil {
		return
	}
	err = r.Run(port)
	if err != nil {
		return
	}
}

func prepareMiddlewares(r *gin.Engine) {
	r.Use(middlewares.TraceRequest)
	r.Use(middlewares.ValidateRequest)
	r.Use(middlewares.SecureRequest)
	r.Use(middlewares.AuthenticateRequest)
}
