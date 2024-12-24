package route

import (
	"ginTemp/apps/demo/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	asrRoute := r.Group("/demo")
	{
		asrRoute.GET("/index", controllers.Default{}.Index)
	}

}
