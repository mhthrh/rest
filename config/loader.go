package config

import (
	"bytes"
	e "errors"
	"github.com/spf13/viper"
	"log"
	"restfullApi/pkg/errors"
	cryptx "restfullApi/util/cryptox"
	env "restfullApi/util/environment"
	"restfullApi/util/file/text"
)

const (
	appName          = "myApp"
	environment      = "environment"
	Extension        = "json"
	fileName         = "plane.json"
	encFileName      = "coded.ENC"
	variableFileName = "file_name"
	filePath         = "app_path"
	path             = "/Users/mohsen/Projects/Golang/practice/app/user/config/file"
	IP               = ""
	Port             = 0
	servicePath      = ""
)

type Local struct {
	v *viper.Viper
}

type Remote struct {
	UserName  string
	Password  string
	Encrypted bool
}

func New(key, extension string, enc bool) (IConfig, *errors.Error) {
	if e := env.GetEnv(environment, ""); e == "service" {
		return Remote{
			Encrypted: false,
		}, nil
	}

	//for local config resource
	newViper := viper.New()
	newViper.SetConfigName(env.GetEnv(fileName, path))
	newViper.SetConfigType(env.GetEnv(extension, Extension))

	err := newViper.ReadConfig(bytes.NewBuffer(func() []byte {
		if !enc {
			txt := text.New(env.GetEnv(filePath, path), env.GetEnv(variableFileName, fileName))
			byt, err := txt.Read()
			if err != nil {
				log.Fatalf("read file failed: %v", err)
			}
			return byt
		}
		c, err := cryptx.New(key)
		if err != nil {
			log.Fatalf("crypto failed: %v", err)
		}
		txt := text.New(env.GetEnv(filePath, path), env.GetEnv(variableFileName, encFileName))
		tt, err := txt.Read()
		if err != nil {
			log.Fatalf("read file failed: %v", err)
		}
		content, err := c.Decrypt(string(tt))
		if err != nil {
			log.Fatalf("decription failed: %v", err)
		}
		return []byte(content)
	}()))
	if err != nil {
		return nil, errors.FailedResource(err, nil)
	}
	return Local{
		v: newViper,
	}, nil
}

func (l Local) DbConfig() (PostgresConfig, *errors.Error) {
	config := PostgresConfig{
		Host:           l.v.GetString("postgresql.host"),
		Port:           l.v.GetInt("postgresql.port"),
		UserName:       l.v.GetString("postgresql.user"),
		Password:       l.v.GetString("postgresql.password"),
		SSLModeEnabled: l.v.GetBool("postgresql.sslEnabled"),
		DatabaseName:   l.v.GetString("postgresql.databaseName"),
		Schema:         l.v.GetString("postgresql.schema"),
		SSLMode:        "",
	}

	if config.Host == "" {
		return PostgresConfig{}, errors.FailedResource(e.New("host is required"), nil)
	}
	if config.Port == 0 {
		return PostgresConfig{}, errors.FailedResource(e.New("port is required"), nil)
	}
	if config.UserName == "" {
		return PostgresConfig{}, errors.FailedResource(e.New("username is required"), nil)
	}
	if config.Password == "" {
		return PostgresConfig{}, errors.FailedResource(e.New("password is required"), nil)
	}

	return config, nil
}

func (l Local) GetRootAdmin() (AdminUser, *errors.Error) {
	admin := AdminUser{
		UserName: l.v.GetString("Redis.Host"),
		Password: l.v.GetString("Redis.Host"),
	}
	if admin.UserName == "" {
		return AdminUser{}, errors.FailedResource(e.New("username is required"), nil)
	}
	if admin.Password == "" {
		return AdminUser{}, errors.FailedResource(e.New("password is required"), nil)
	}
	return admin, nil
}

func (s Remote) DbConfig() (PostgresConfig, *errors.Error) {
	return PostgresConfig{}, errors.NotImplemented("service")
}

func (s Remote) GetRootAdmin() (AdminUser, *errors.Error) {
	return AdminUser{}, errors.NotImplemented("service")

}
