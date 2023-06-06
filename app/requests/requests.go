package requests

import (
	"GOHUB/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/thedevsaddam/govalidator"
)

type ValidatorFunc func(interface{}, *gin.Context) map[string][]string

func Validate(c *gin.Context, obj interface{}, handler ValidatorFunc) bool {
	//1解析请求 支持json数据，表单请求和URL query
	if err := c.ShouldBind(obj); err != nil {
		response.BadRequest(c, err,
			"请求解析错误，请确认请求格式是否正确。上传文件请使用 multipart 标头，参数请使用 JSON 格式。")

		return false

	}
	//2.表单验证
	errs := handler(obj, c)

	//3判断验证是否通过
	if len(errs) > 0 {
		response.ValidationError(c, errs)
		return false
	}
	return true
}

func validate(data interface{}, rules govalidator.MapData, messages govalidator.MapData) map[string][]string {
	//配置初始化
	opts := govalidator.Options{
		Data:          data,
		Rules:         rules,
		Messages:      messages,
		TagIdentifier: "valid", // 模型中的 SignupEmailExistRequest Struct 标签标识符
	}
	return govalidator.New(opts).ValidateStruct()
}
