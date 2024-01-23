package appctr

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
)

// WhiteOriginList List of origins that a CORS request can be executed from.
// Never can be * in production
// TODO CHANGE THIS TO YOUR OWN ORIGINS
const (
	WhiteOriginList = "*"
)

func UseMiddlewares(r *gin.Engine) {
	// Logger middleware
	r.Use(gin.Logger())
	// Recovery middleware
	r.Use(gin.Recovery())

	allowedOrigins := WhiteOriginList
	if Env() == EnvDev {
		allowedOrigins = "*"
	}

	// CORS middleware
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{allowedOrigins},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Authorization", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
	}))

	// Content Type Negotiation middleware
	r.Use(static.Serve("/", static.LocalFile("./public", true)))
}
