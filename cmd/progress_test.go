/*
	These tests are designed to run against an instance of Progress API
	with a PostgreSQL data store backing configured.

	The goal is to provide the end-to-end validation of all components.
*/

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gopkg.in/resty.v1"

	"github.com/hoanhan101/progress/internal/config"
)

const (
	host = "http://localhost:8000"
)

func TestConfig(t *testing.T) {
	cfg := new(config.Config)
	_, err := resty.New().SetHostURL(host).R().SetResult(cfg).Get("/config")
	assert.NoError(t, err)
	assert.Equal(t, ":8000", cfg.Server.Address)
	assert.Equal(t, "postgres", cfg.Database.User)
	assert.Equal(t, "postgres", cfg.Database.Password)
	assert.Equal(t, "postgres", cfg.Database.Host)
	assert.Equal(t, 5432, cfg.Database.Port)
	assert.Equal(t, "disable", cfg.Database.SSLMode)
}

func TestHealth(t *testing.T) {
	health := &struct {
		Status string
	}{}

	_, err := resty.New().SetHostURL(host).R().SetResult(health).Get("/health")
	assert.NoError(t, err)
	assert.Equal(t, "ok", health.Status)
}
