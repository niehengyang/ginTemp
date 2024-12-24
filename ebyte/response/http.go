package response

import "github.com/gin-gonic/gin"

//func Success(c *gin.Context, data any) {
//	c.JSON(200, data)
//}

func Success(c *gin.Context, data interface{}) {
	c.JSON(200, gin.H{
		"code": 200,
		"data": data,
	})
}

// 资源已存在
func Error409(c *gin.Context, message string) {
	Error(c, 409, "409", message)
}

func Error404(c *gin.Context, message string) {
	Error(c, 404, "404", message)
}

func Error500(c *gin.Context, message string) {
	Error(c, 500, "500", message)
}

func Error401(c *gin.Context, message string) {
	Error(c, 401, "401", message)
}

func Error400(c *gin.Context, message string) {
	Error(c, 400, "400", message)
}

func Error403(c *gin.Context, message string) {
	Error(c, 403, "403", message)
}

func Error(c *gin.Context, httpCode int, ret, message string) {
	c.JSON(httpCode, gin.H{
		"ret":     ret,
		"message": message,
	})
}
