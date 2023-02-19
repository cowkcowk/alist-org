package conf

import "net/url"

var (
	BuiltAt    string
	GoVersion  string
	GitAuthor  string
	GitCommit  string
	Version    string = "dev"
	WebVersion string
)

var (
	Conf *Config
	URL  *url.URL
)
