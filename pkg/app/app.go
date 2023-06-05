package app

import "GOHUB/pkg/config"

func IsLocal() bool {
	return config.Get("app.env") == "local"

}
func IsProduction() bool {
	return config.Get("app.env") == "production"
}

func IsTesting() bool {
	return config.Get("app.env") == "testing"
}