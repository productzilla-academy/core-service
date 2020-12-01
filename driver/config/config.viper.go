package config

import (
	"errors"

	"github.com/ory/viper"
	"github.com/sirupsen/logrus"
)

const (
	viperKeyListenPort = "listen.port"
	viperKeyListenHost = "listen.host"

	viperKeyDSN = "dsn"

	viperKeyElasticsearchEndpoint     = "elasticsarch.endpoint"
	viperKeyElasticsearchAuthUsername = "elasticsearch.auth.username"
	viperKeyElasticsearchAuthPassword = "elasticsearch.auth.password"

	viperKeyMinioEndpoint     = "minio.endpoint"
	viperKeyMinioAuthUsername = "minio.auth.username"
	viperKeyMinioAuthPassword = "minio.auth.password"
)

func GetString(key, defaultValue string) string {
	viper.SetDefault(key, defaultValue)
	return viper.GetString(key)
}

func GetInt(key string, defaultValue int64) int64 {
	viper.SetDefault(key, defaultValue)
	return viper.GetInt64(key)
}

func GetBool(key string, defaultValue bool) bool {
	viper.SetDefault(key, defaultValue)
	return viper.GetBool(key)
}

type ConfigViper struct {
	logger Logger
}

func (this *ConfigViper) Logger() Logger {
	if this.logger == nil {
		logger := logrus.New()
		this.logger = logger
	}
	return this.logger
}
func (this *ConfigViper) ListenPort() int64 {
	return GetInt(viperKeyListenPort, 3000)
}
func (this *ConfigViper) ListenHost() string {
	return GetString(viperKeyListenHost, "0.0.0.0")

}
func (this *ConfigViper) DSN() string {
	dsn := GetString(viperKeyDSN, "")
	if dsn == "" {
		this.logger.Fatal(errors.New("No DSN"))
	}
	return dsn
}

func (this *ConfigViper) ElasticsearchEndpoint() string {
	return GetString(viperKeyElasticsearchEndpoint, "http://elasticsearch:9200")
}
func (this *ConfigViper) ElasticsearchAuthUsername() string {
	return GetString(viperKeyElasticsearchAuthUsername, "")
}
func (this *ConfigViper) ElasticsearchAuthPassword() string {
	return GetString(viperKeyElasticsearchAuthPassword, "")
}

func (this *ConfigViper) MinioEndpoint() string {
	return GetString(viperKeyMinioEndpoint, "http://elasticsearch:9000")
}
func (this *ConfigViper) MinioAuthID() string {
	return GetString(viperKeyMinioAuthUsername, "")

}
func (this *ConfigViper) MinioAuthKey() string {
	return GetString(viperKeyMinioAuthPassword, "")
}

func NewConfigViper() *ConfigViper {
	return &ConfigViper{}
}
