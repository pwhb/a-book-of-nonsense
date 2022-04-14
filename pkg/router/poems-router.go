package router

import (
	"github.com/gin-gonic/gin"
	"github.com/pwhb/lear/pkg/controller"
)

func ConfigurePoemsRoutes(r *gin.Engine) {
	r.GET("/api/poems", controller.GetAllPoems)
	r.GET("/api/poems/random", controller.GetRandomPoem)
	r.GET("/api/poems/:id", controller.GetOnePoem)
	r.GET("/", controller.RenderHTML)
}
