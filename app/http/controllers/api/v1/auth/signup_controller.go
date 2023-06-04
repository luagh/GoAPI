package auth

import (
	v1 "GOHUB/app/http/controllers/api/v1"
	"GOHUB/app/models/user"
	"GOHUB/app/requests"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 处理用户身份认证相关逻辑
type SingupController struct {
	v1.BaseAPIController
}

//检测手机是否被注册

func (sc *SingupController) IsPhoneExist(c *gin.Context) {

	//请求的对象

	request := requests.SignupPhoneExistRequest{}

	//解析json请求
	if err := c.ShouldBindJSON(&request); err != nil {

		//解析失败 返回一个状态码和错误信息
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		return
	}
	errs := requests.ValidateSignupPhoneExist(&request, c)
	if len(errs) > 0 {
		c.AbortWithStatusJSON(http.StatusUnprocessableEntity, gin.H{
			"errors": errs,
		})
		return
	}

	//检查数据库返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}
