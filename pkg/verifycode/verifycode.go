package verifycode

import (
	"GOHUB/pkg/app"
	"GOHUB/pkg/config"
	"GOHUB/pkg/helpers"
	"GOHUB/pkg/logger"
	"GOHUB/pkg/mail"
	"GOHUB/pkg/redis"
	"GOHUB/pkg/sms"
	"fmt"
	"strings"
	"sync"
)

// 这个页面代码用以发送手机验证码和邮箱验证码

type VerifyCode struct {
	Store Store
}

var once sync.Once
var internalVerifyCode *VerifyCode

// NewVerifyCode 单例模式获取
func NewVerifyCode() *VerifyCode {
	once.Do(func() {
		internalVerifyCode = &VerifyCode{
			Store: &RedisStore{
				RedisClient: redis.Redis,
				KeyPrefix:   config.GetString("app.name") + ":verifycode",
			},
		}
	})
	return internalVerifyCode
}

// SendSMS 发送短信验证码
func (vc *VerifyCode) SendSMS(phone string) bool {
	// 生成验证码
	code := vc.generateVerifyCode(phone)
	// 方便本地和 API 自动测试
	//if !app.IsProduction() && phone == "13017173106" {
	//	return true
	//}
	if !app.IsProduction() && strings.HasPrefix(phone, config.GetString("app.tep")) {
		return true
	}

	return sms.NewSMS().Send(phone, sms.Message{
		Template: config.GetString("sms.aliyun.template_code"),
		Data:     map[string]string{"code": code},
	})
}

// SendEmail 发送邮件验证码
func (vc *VerifyCode) SendEmail(email string) error {

	// 生成验证码
	code := vc.generateVerifyCode(email)

	// 方便本地和 API 自动测试
	if !app.IsProduction() && strings.HasSuffix(email, config.GetString("verifycode.debug_email_suffix")) {
		return nil
	}
	content := fmt.Sprintf("<h1>您的 Email 验证码是 %v </h1>", code)
	// 发送邮件
	mail.NewMailer().Send(mail.Email{
		From: mail.From{
			Address: config.GetString("mail.from.address"),
			Name:    config.GetString("mail.from.name"),
		},
		To:      []string{email},
		Subject: "Email 验证码",
		HTML:    []byte(content),
	})

	return nil
}

// CheckAnswer 检查用户提交的验证码是否正确，key 可以是手机号或者 Email
func (vc *VerifyCode) CheckAnswer(key string, answer string) bool {
	logger.DebugJSON("验证码", "检查验证码", map[string]string{key: answer})
	// 方便开发，在非生产环境下，具备特殊前缀的手机号和 Email后缀，会直接验证成功
	if !app.IsProduction() &&
		(strings.HasSuffix(key, config.GetString("verifycode.debug_email_suffix")) ||
			strings.HasPrefix(key, config.GetString("app.tep"))) {
		return true
	}
	return vc.Store.Verify(key, answer, false)
}

// 生成验证码，并放置于 Redis 中
func (vc *VerifyCode) generateVerifyCode(key string) string {

	code := helpers.RandomNumber(config.GetInt("verifycode.code_length"))
	// 为方便开发，本地环境使用固定验证码
	if app.IsLocal() {
		code = config.GetString("verifycode.debug_code")
	}
	logger.DebugJSON("验证码", "生成验证码", map[string]string{key: code})
	// 将验证码及 KEY（邮箱或手机号）存放到 Redis 中并设置过期时间
	vc.Store.Set(key, code)
	return code
}
