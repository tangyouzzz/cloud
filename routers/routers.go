package routers

import (
	"github.com/gin-gonic/gin"
	"main/controller"
	"net/http"
	"main/middleware"

)

func SetupRouter() *gin.Engine{
	r :=gin.Default()
	r.Static("/static","./static")
	r.LoadHTMLFiles("./templates/home.html")
	//user group
	user :=r.Group("/user")
	{
		user.POST("/auth", controller.AuthHandler)
		user.GET("/index", middleware.JWTAuthMiddleware(), controller.HomeHandler)
	
	}
	//shouye
	r.GET("/home",func (c *gin.Context)  {
		c.HTML(http.StatusOK,"home.html",nil)
	})
	//file group
	file :=r.Group("/file",middleware.JWTAuthMiddleware())
	{
		file.GET("/list",controller.Filefind)
		file.POST("/upload",controller.Fileupload)
		file.GET("/download",controller.Filedownload)
		file.GET("/delete",controller.Filedelete)	
	}

	return r
}