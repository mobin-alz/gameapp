package config

var defaultConfig = map[string]interface{}{
	"auth.access_subject":          AccessTokenSubject,
	"auth.refresh_subject":         RefreshTokenSubject,
	"auth.refresh_expiration_time": RefreshTokenExpireDuration,
	"auth.access_expiration_time":  AccessTokenExpireDuration,
}
