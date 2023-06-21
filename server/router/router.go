package router

import (
	"server/internal/user"
	"server/internal/ws"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var r *gin.Engine

func InitRouter(userHandler *user.Handler, wsHandler *ws.Handler) {
	r = gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"GET", "POST"},
		AllowHeaders:     []string{"Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "http://localhost:3000"
		},
		MaxAge: 12 * time.Hour,
	}))

	r.POST("/api/signup", userHandler.CreateUser)
	r.POST("/api/login", userHandler.Login)
	r.GET("/api/logout", userHandler.Logout)

	r.POST("/api/ws/createRoom", wsHandler.CreateRoom)
	r.GET("/api/ws/joinRoom/:roomId", wsHandler.JoinRoom)
	r.GET("/api/ws/getRooms", wsHandler.GetRooms)
	r.GET("/api/ws/getClients/:roomId", wsHandler.GetClients)
}

func Start(addr string) error {
	return r.Run(addr)
}
