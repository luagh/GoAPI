package routes

//注册网页相关路由
import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func RegisterAPIRoutes(r *gin.Engine) {

	v1 := r.Group("/v1")
	{
		v1.GET("/", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"hello": "wdf",
			})
		})
	}
}
