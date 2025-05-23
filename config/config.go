package config

import (
	"field-service/common/util"
	"github.com/sirupsen/logrus"
)

var Config AppConfig

type AppConfig struct {
	Port                  int             `json:"port"`
	AppName               string          `json:"appName"`
	AppEnv                string          `json:"appEnv"`
	SignatureKey          string          `json:"signatureKey"`
	Database              Database        `json:"database"`
	RateLimiterMaxRequest float64         `json:"rateLimiterMaxRequest"`
	RateLimiterTimeSecond int             `json:"rateLimiterTimeSecond"`
	InternalService       InternalService `json:"internalService"`
	Minio                 Minio           `json:"minio"`
}

type Database struct {
	Host                  string `json:"host"`
	Port                  int    `json:"port"`
	Name                  string `json:"name"`
	Username              string `json:"username"`
	Password              string `json:"password"`
	MaxOpenConnections    int    `json:"maxOpenConnections"`
	MaxLifeTimeConnection int    `json:"maxLifeTimeConnection"`
	MaxIdleConnections    int    `json:"maxIdleConnections"`
	MaxIdleTime           int    `json:"maxIdleTime"`
}

type InternalService struct {
	User User `json:"user"`
}

type User struct {
	Host         string `json:"host"`
	SignatureKey string `json:"signatureKey"`
}

type Minio struct {
	Address    string `json:"address"`
	AccessKey  string `json:"accessKey"`
	Secret     string `json:"secret"`
	UseSsl     bool   `json:"useSsl"`
	BucketName string `json:"bucketName"`
}

func Init() {
	err := util.BindFromJSON(&Config, "config", ".")
	if err != nil {
		logrus.Infof("failed to bind config: %v", err)
		err = util.BindFromEnv(&Config)
		if err != nil {
			panic(err)
		}
	}
}
