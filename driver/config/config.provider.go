package config

type ConfigProvider interface {
	Logger() Logger
	ListenPort() int64
	ListenHost() string
	DSN() string

	ElasticsearchEndpoint() string
	ElasticsearchAuthUsername() string
	ElasticsearchAuthPassword() string

	MinioEndpoint() string
	MinioAuthID() string
	MinioAuthKey() string
}
