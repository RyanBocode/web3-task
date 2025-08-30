package router

import (
	"blog/internal/config"
	"blog/internal/controllers"
	"blog/internal/middleware"

	"github.com/gin-gonic/gin"
)

func Setup(cfg *config.Config) *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())

	// 静态文件服务
	r.Static("/static", "./static")

	// 首页路由 - 直接返回HTML文件
	r.GET("/", func(c *gin.Context) {
		c.File("./static/index.html")
	})

	v1 := r.Group("/api/v1")
	{
		authCtl := controllers.NewAuthController(cfg)
		postCtl := controllers.NewPostController()
		cmtCtl := controllers.NewCommentController()

		// Auth
		auth := v1.Group("/auth")
		{
			auth.POST("/register", authCtl.Register)
			auth.POST("/login", authCtl.Login)
		}

		v1.GET("/posts", postCtl.List)
		v1.GET("/posts/:id", postCtl.Get)
		v1.GET("/posts/:id/comments", cmtCtl.ListByPost)

		// Protected routes
		pr := v1.Group("")
		pr.Use(middleware.AuthRequired(cfg))
		{
			pr.GET("/me", authCtl.Me)
			pr.POST("/posts", postCtl.Create)
			pr.PUT("/posts/:id", postCtl.Update)
			pr.DELETE("/posts/:id", postCtl.Delete)

			pr.POST("/posts/:id/comments", cmtCtl.Create)
		}
	}
	return r
}
