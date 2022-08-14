package config

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestNewConfig_error_json(t *testing.T) {
	conf, err := NewConfig("../../config/config.json")
	assert.NotNil(t, err)
	assert.Nil(t, conf)
}

func TestNewConfig_error_yml(t *testing.T) {
	conf, err := NewConfig("../../config/test.yml")
	assert.NotNil(t, err)
	assert.Nil(t, conf)
}

func TestNewConfig_error_unn(t *testing.T) {
	conf, err := NewConfig("../../config/test")
	assert.NotNil(t, err)
	assert.Nil(t, conf)
}

func TestNewConfig_error_unread(t *testing.T) {
	conf, err := NewConfig("../../config/test1")
	assert.NotNil(t, err)
	assert.Nil(t, conf)
}

func TestNewConfig_json(t *testing.T) {
	assert.Nil(t, os.Setenv("TIME_SEC", "15"))
	assert.Nil(t, os.Setenv("PG_USER", "test"))
	assert.Nil(t, os.Setenv("PG_PASSWORD", "test"))

	conf, err := NewConfig("../../config/config.json")
	assert.Nil(t, err)
	assert.NotNil(t, conf)
	if conf.GetPeriod() != 15 {
		t.Errorf("Unexpected period")
	}
	assert.Nil(t, os.Setenv("TIME_SEC", ""))
	assert.Nil(t, os.Setenv("PG_USER", ""))
	assert.Nil(t, os.Setenv("PG_PASSWORD", ""))
}

func TestNewConfig_yml(t *testing.T) {
	assert.Nil(t, os.Setenv("TIME_SEC", "15"))
	assert.Nil(t, os.Setenv("PG_USER", "test"))
	assert.Nil(t, os.Setenv("PG_PASSWORD", "test"))

	conf, err := NewConfig("../../config/test.yml")
	assert.Nil(t, err)
	assert.NotNil(t, conf)
	if conf.GetPeriod() != 15 {
		t.Errorf("Unexpected period")
	}
	assert.Nil(t, os.Setenv("TIME_SEC", ""))
	assert.Nil(t, os.Setenv("PG_USER", ""))
	assert.Nil(t, os.Setenv("PG_PASSWORD", ""))
}

func TestNewConfig_error_user(t *testing.T) {
	assert.Nil(t, os.Setenv("TIME_SEC", "15"))
	//assert.Nil(t, os.Setenv("PG_USER", "test"))
	assert.Nil(t, os.Setenv("PG_PASSWORD", "test"))

	conf, err := NewConfig("../../config/config.json")
	assert.NotNil(t, err)
	assert.Nil(t, conf)
	assert.Nil(t, os.Setenv("TIME_SEC", ""))
	assert.Nil(t, os.Setenv("PG_USER", ""))
	assert.Nil(t, os.Setenv("PG_PASSWORD", ""))
}

func TestNewConfig_error_pass(t *testing.T) {
	assert.Nil(t, os.Setenv("TIME_SEC", "15"))
	assert.Nil(t, os.Setenv("PG_USER", "test"))
	//assert.Nil(t, os.Setenv("PG_PASSWORD", "test"))

	conf, err := NewConfig("../../config/config.json")
	assert.NotNil(t, err)
	assert.Nil(t, conf)
	assert.Nil(t, os.Setenv("TIME_SEC", ""))
	assert.Nil(t, os.Setenv("PG_USER", ""))
	assert.Nil(t, os.Setenv("PG_PASSWORD", ""))
}

func TestNewConfig_error_user_yml(t *testing.T) {
	assert.Nil(t, os.Setenv("TIME_SEC", "15"))
	//assert.Nil(t, os.Setenv("PG_USER", "test"))
	assert.Nil(t, os.Setenv("PG_PASSWORD", "test"))

	conf, err := NewConfig("../../config/test.yml")
	assert.NotNil(t, err)
	assert.Nil(t, conf)
	assert.Nil(t, os.Setenv("TIME_SEC", ""))
	assert.Nil(t, os.Setenv("PG_USER", ""))
	assert.Nil(t, os.Setenv("PG_PASSWORD", ""))
}

func TestNewConfig_error_pass_yml(t *testing.T) {
	assert.Nil(t, os.Setenv("TIME_SEC", "15"))
	assert.Nil(t, os.Setenv("PG_USER", "test"))
	//assert.Nil(t, os.Setenv("PG_PASSWORD", "test"))

	conf, err := NewConfig("../../config/test.yml")
	assert.NotNil(t, err)
	assert.Nil(t, conf)
	assert.Nil(t, os.Setenv("TIME_SEC", ""))
	assert.Nil(t, os.Setenv("PG_USER", ""))
	assert.Nil(t, os.Setenv("PG_PASSWORD", ""))
}
