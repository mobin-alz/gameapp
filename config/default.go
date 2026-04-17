package config

var defaultConfig = map[string]interface{}{
	"auth.access_subject":  AccessTokenSubject,
	"auth.refresh_subject": RefreshTokenSubject,
}
