package routers

import (
	"github.com/gin-gonic/gin"
	"supremine/lib/setting"
	"supremine/routers/api"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	gin.SetMode(setting.RunMode)

	apiv1 := r.Group("/api")
	{
		apiv1.GET("/tags", api.GetTags)
		apiv1.POST("/tags", api.AddTag)
		apiv1.PUT("/tags/:id", api.EditTag)
		apiv1.DELETE("/tags/:id", api.DeleteTag)
	}

	return r
}