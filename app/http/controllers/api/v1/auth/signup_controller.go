package auth

import (
	v1 "GOHUB/app/http/controllers/api/v1"
	"GOHUB/app/models/user"
	"GOHUB/app/requests"
	"github.com/gin-gonic/gin"
	"net/http"
)

// 处理用户身份认证相关逻辑

// SignupController 注册控制器
type SingupController struct {
	v1.BaseAPIController
}

//检测手机是否被注册

func (sc *SingupController) IsPhoneExist(c *gin.Context) {

	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}

	//检查数据库返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsPhoneExist(request.Phone),
	})
}

// 检测邮箱是否被注册
func (sc *SingupController) IsEmailExist(c *gin.Context) {
	request := requests.SignupEmailExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupEmailExist); !ok {
		return
	}
	//检查数据库返回响应
	c.JSON(http.StatusOK, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})

}
