package config

import (
	"fmt"
	"time"
)

type Server struct {
	Host         string        `json:"host"`
	Port         int           `json:"port"`
	ReadTimeOut  time.Duration `json:"readTimeOut"`
	WriteTimeOut time.Duration `json:"writeTimeOut"`
	IdleTimeOut  time.Duration `json:"idleTimeOut"`
}

type AdminUser struct {
	UserName string `json:"userName"`
	Password string `json:"password"`
}
type PostgresConfig struct {
	Host           string  `yaml:"host" json:"host"`
	Port           int     `yaml:"port" json:"port"`
	UserName       string  `yaml:"username" json:"username"`
	Password       string  `yaml:"password" json:"password"`
	SSLModeEnabled bool    `yaml:"sslEnabled" json:"sslEnabled"`
	DatabaseName   string  `yaml:"database" json:"database"`
	Schema         string  `yaml:"schema" json:"schema"`
	SSLMode        SSLMode `yaml:"sslmode" json:"sslmode"`
}

type SSLMode string

const (
	_          SSLMode = ""
	Disabled           = "disable"
	Require            = "require"
	VerifyCA           = "verify-ca"
	VerifyFull         = "verify-full"
)

func IsValid(v interface{}) (SSLMode, error) {
	var s SSLMode
	switch v {
	case Disabled:
		s = Disabled
	case Require:
		s = Require
	case VerifyCA:
		s = VerifyCA
	case VerifyFull:
		s = VerifyFull
	default:
		return "", fmt.Errorf("invalid ssl mode type")
	}
	return s, nil
}
