package config

import "GOHUB/pkg/config"

func init() {
	config.Add("database", func() map[string]interface{} {
		return map[string]interface{}{
			//默认 数据库
			"connection": config.Env("DB_CONNECTION", "mysql"),

			"mysql": map[string]interface{}{
				//数据库链接信息
				"host":     config.Env("DB_HOST", "127.0.0.1"),
				"port":     config.Env("DB_PORT", "3306"),
				"database": config.Env("DB_DATABASE", "gohub"),
				"username": config.Env("DB_USERNAME", ""),
				"password": config.Env("DB_PASSWORD", ""),
				"charset":  config.Env("utf8mb4"),

				//链接池配置 (连接池的最大空闲连接数、最大打开连接数和连接生命周期)
				"max_idle_connections": config.Env("DB_MAX_TDLE_CONNECTIONS", 100),
				"max_open_connections": config.Env("DB_MAX_OPEN_CONNECTIONS", 25),
				"max_life_seconds":     config.Env("DB_MAX_LIFE_SECONDS", 5*60),
			},
			"sqlite": map[string]interface{}{
				"database": config.Env("DB_SQL_FTLE", "database/database.db"),
			},
		}
	})
}
