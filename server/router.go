package server

import (
	"mithril/api"
	"mithril/middleware"
	//	"os"

	"github.com/gin-gonic/gin"
)

// NewRouter 路由配置
func NewRouter() *gin.Engine {
	r := gin.Default()

	// 中间件, 顺序不能改
	r.Use(middleware.Cors())
	var authMiddleware = middleware.GinJWTMiddlewareInit(middleware.AllUserAuthorizator)
	// 路由
	v1 := r.Group("/api/v1")
	{
		v1.POST("ping", api.Ping)

		// 用户注册
		v1.POST("user/register", api.UserRegister)

		// 用户登录
		v1.POST("user/login", authMiddleware.LoginHandler)

		// 需要登录保护的
		auth := v1.Group("")

		auth.Use(authMiddleware.MiddlewareFunc())
		{
			// User Routing
			auth.GET("/refresh_token", authMiddleware.RefreshHandler)

			auth.GET("user/me", api.UserMe)

			auth.PUT("video/:id", api.UpdateVideo)
			auth.DELETE("video/:id", api.DeleteVideo)
			auth.POST("videos", api.CreateVideo)
		}
		v1.GET("video/:id", api.ShowVideo)
		v1.GET("videos", api.ListVideo)
		// 排行榜
		v1.GET("rank/daily", api.DailyRank)
		// 其他
		v1.POST("upload/token", api.UploadToken)
	}
	return r
}
