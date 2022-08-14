package config

import (
	"encoding/json"
	"fmt"
	"gopkg.in/yaml.v3"
	"log"
	"os"
	"strconv"
	"strings"
)

var (
	errNotFoundSecrets = "not found secret"
	pgUser             = "PG_USER"
	pgPass             = "PG_PASSWORD"
	dbHost             = "DB_HOST"
)

type IConfig interface {
	GetSQL() *SQL
	GetPort() string
	GetPeriod() int
}

type SQL struct {
	Host            string `json:"-" yml:"-"`
	Port            string `json:"port" yaml:"port"`
	UserID          string `json:"user_id" yaml:"user_id"`
	Password        string `json:"password" yaml:"password"`
	Database        string `json:"database" yaml:"database"`
	MaxIdleCons     int    `json:"max_idle_cons" yaml:"max_idle_cons"`
	MaxOpenCons     int    `json:"max_open_cons" yaml:"max_open_cons"`
	ConnMaxLifetime int    `json:"conn_max_lifetime" yaml:"conn_max_lifetime"`
}

type Server struct {
	Host string `json:"host" yml:"host"`
	Port string `json:"port" yaml:"port"`
}

type Config struct {
	sql    *SQL
	serv   *Server
	period int
}

func NewConfig(path string) (IConfig, error) {
	var instance IConfig
	var f func([]byte) (IConfig, error)

	bytes, err := os.ReadFile(path)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	switch {
	case strings.HasSuffix(path, ".yml"):
		f = fromYML
	case strings.HasSuffix(path, ".json"):
		f = fromJSON
	default:
		return nil, fmt.Errorf("Invalid type file: " + path)
	}

	instance, err = f(bytes)
	if err != nil {
		instance = nil
	}

	return instance, err
}

func fromJSON(b []byte) (IConfig, error) {
	temp := struct {
		SQL    SQL    `json:"sql"`
		Serv   Server `json:"server"`
		period int
	}{}

	err := json.Unmarshal(b, &temp)
	if err != nil {
		return nil, err
	}

	temp.period, err = strconv.Atoi(os.Getenv("TIME_SEC"))
	if err != nil {
		return nil, fmt.Errorf("error parse peroid time from .env")
	}

	if temp.SQL.UserID = os.Getenv(pgUser); temp.SQL.UserID == "" {
		return nil, fmt.Errorf("%s by %s", errNotFoundSecrets, pgUser)
	}
	if temp.SQL.Password = os.Getenv(pgPass); temp.SQL.Password == "" {
		return nil, fmt.Errorf("%s by %s", errNotFoundSecrets, pgPass)
	}

	temp.SQL.Host = os.Getenv(dbHost)
	return &Config{
		serv:   &temp.Serv,
		sql:    &temp.SQL,
		period: temp.period,
	}, nil
}

func fromYML(b []byte) (IConfig, error) {
	temp := struct {
		SQL    SQL    `yaml:"sql"`
		Serv   Server `yaml:"server"`
		period int
	}{}

	err := yaml.Unmarshal(b, &temp)
	if err != nil {
		return nil, err
	}

	temp.period, err = strconv.Atoi(os.Getenv("TIME_SEC"))
	if err != nil {
		return nil, fmt.Errorf("error parse peroid time from .env")
	}

	if temp.SQL.UserID = os.Getenv(pgUser); temp.SQL.UserID == "" {
		return nil, fmt.Errorf("%s by %s", errNotFoundSecrets, pgUser)
	}
	if temp.SQL.Password = os.Getenv(pgPass); temp.SQL.Password == "" {
		return nil, fmt.Errorf("%s by %s", errNotFoundSecrets, pgPass)
	}

	temp.SQL.Host = os.Getenv(dbHost)
	return &Config{
		serv:   &temp.Serv,
		sql:    &temp.SQL,
		period: temp.period,
	}, nil
}

func (c *Config) GetSQL() *SQL {
	return c.sql
}

func (c *Config) GetPort() string {
	return c.serv.Port
}

func (c *Config) GetPeriod() int {
	return c.period
}
