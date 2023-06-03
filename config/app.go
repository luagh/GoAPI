package config

import conf "GOHUB/pkg/config"

// 站点配置信息
func init() {
	conf.Add("app", func() map[string]interface{} {
		return map[string]interface{}{
			// 应用名称
			"name": conf.Env("APP_NAME", "Gohub"),
			"env":  conf.Env("APP_ENV", "production"),
			// 是否进入调试模式
			"debug": conf.Env("APP_DEBUG", false),
			// 应用服务端口
			"port": conf.Env("APP_PORT", "8081"),
			// 加密会话、JWT 加密
			"key": conf.Env("APP_KEY", "33446a9dcf9ea060a0a6532b166da32f304af0de"),
			// 用以生成链接
			"url": conf.Env("APP_URL", "http://localhost:8081"),
			// 设置时区，JWT 里会使用，日志记录里也会使用到
			"timezone": conf.Env("TIMEZONE", "Asia/Shanghai"),
		}
	})
}
