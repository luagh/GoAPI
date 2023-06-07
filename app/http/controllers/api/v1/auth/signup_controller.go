package auth

import (
	v1 "GOHUB/app/http/controllers/api/v1"
	"GOHUB/app/models/user"
	"GOHUB/app/requests"
	"GOHUB/pkg/jwt"
	"GOHUB/pkg/response"
	"github.com/gin-gonic/gin"
)

// 处理用户身份认证相关逻辑

// SignupController 注册控制器
type SingupController struct {
	v1.BaseAPIController
}

//检测手机是否被注册

func (sc *SingupController) IsPhoneExist(c *gin.Context) {
	// 获取请求参数，并做表单验证
	request := requests.SignupPhoneExistRequest{}
	if ok := requests.Validate(c, &request, requests.SignupPhoneExist); !ok {
		return
	}

	//检查数据库返回响应
	response.JSON(c, gin.H{
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
	response.JSON(c, gin.H{
		"exist": user.IsEmailExist(request.Email),
	})

}

// SignupUsingPhone 使用手机和验证码进行注册
func (sc *SingupController) SignupUsingPhone(c *gin.Context) {
	// 1. 验证表单
	request := requests.SignupUsingPhoneRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingPhone); !ok {
		return
	}
	userModel := user.User{
		Name:     request.Name,
		Phone:    request.Phone,
		Password: request.Password,
	}
	userModel.Create()

	if userModel.ID > 0 {
		token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name)
		response.CreatedJSON(c, gin.H{
			"token": token,
			"data":  userModel,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试")
	}

}

// SignupUsingEmail 使用 Email + 验证码进行注册
func (sc *SingupController) SignupUsingEmail(c *gin.Context) {

	request := requests.SignupUsingEmailRequest{}
	if ok := requests.Validate(c, &request, requests.SignupUsingEmail); !ok {
		return
	}

	// 2. 验证成功，创建数据
	userModel := user.User{
		Name:     request.Name,
		Email:    request.Email,
		Password: request.Password,
	}
	userModel.Create()

	if userModel.ID > 0 {
		token := jwt.NewJWT().IssueToken(userModel.GetStringID(), userModel.Name)
		response.CreatedJSON(c, gin.H{
			"token": token,
			"data":  userModel,
		})
	} else {
		response.Abort500(c, "创建用户失败，请稍后尝试~")
	}
}
