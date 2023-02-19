package conf

type Database struct {
	Type        string `json:"type" env:"DB_TYPE"`
	Host        string `json:"host" env:"DB_HOST"`
	Port        int    `json:"port" env:"DB_PORT"`
	User        string `json:"user" env:"DB_USER"`
	Password    string `json:"password" env:"DB_PASS"`
	Name        string `json:"name" env:"DB_NAME"`
	DBFile      string `json:"db_file" env:"DB_FILE"`
	TablePrefix string `json:"table_prefix" env:"DB_TABLE_PREFIX"`
	SSLMode     string `json:"ssl_mode" env:"DB_SSL_MODE"`
}

type Config struct {
	SiteURL        string   `json:"site_url" env:"SITE_URL"`
	JwtSecret      string   `json:"jwt_secret" env:"JWT_SECRET"`
	TokenExpiresIn int      `json:"token_expires_in" env:"TOKEN_EXPIRES_IN"`
	Database       Database `json:"database"`
}
