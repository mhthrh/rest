package config

import "restfullApi/pkg/errors"

type IConfig interface {
	DbConfig() (PostgresConfig, *errors.Error)
	GetRootAdmin() (AdminUser, *errors.Error)
}
