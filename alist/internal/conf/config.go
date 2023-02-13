package conf

type Config struct {
	SiteURL        string `json:"site_url" env:"SITE_URL"`
	TokenExpiresIn int    `json:"token_expires_in" env:"TOKEN_EXPIRES_IN"`
}
