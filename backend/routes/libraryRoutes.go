package routes

import (
	"online-library-system/controllers"
	"online-library-system/middleware"

	"github.com/gin-gonic/gin"
)

func LibraryRoutes(router *gin.Engine) {
	router.POST("/libraries", middleware.RoleBasedAccessControl("Owner"), controllers.CreateLibrary)
	router.GET("/libraries", middleware.RoleBasedAccessControl("Owner"), controllers.GetLibraries)
	router.DELETE("/libraries/:id", controllers.DeleteLibrary)

}
