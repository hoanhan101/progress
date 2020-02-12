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
	"github.com/hoanhan101/progress/internal/model"
)

var request = resty.New().SetHostURL("http://localhost:8000").R()

func TestE2EConfig(t *testing.T) {
	cfg := new(config.Config)
	_, err := request.SetResult(cfg).Get("/config")
	assert.NoError(t, err)
	assert.Equal(t, ":8000", cfg.Server.Address)
	assert.Equal(t, "postgres", cfg.Database.User)
	assert.Equal(t, "postgres", cfg.Database.Password)
	assert.Equal(t, "postgres", cfg.Database.Host)
	assert.Equal(t, 5432, cfg.Database.Port)
	assert.Equal(t, "disable", cfg.Database.SSLMode)
}

func TestE2EHealth(t *testing.T) {
	health := &struct {
		Status string
	}{}
	_, err := request.SetResult(health).Get("/health")
	assert.NoError(t, err)
	assert.Equal(t, "ok", health.Status)
}

func TestE2EGoal(t *testing.T) {
	// Get all goals.
	goals := new([]model.Goal)
	_, err := request.SetResult(goals).Get("/goal")
	assert.NoError(t, err)
	assert.Equal(t, 0, len(*goals))

	// Create a new goal.
	newGoal := new(model.Goal)
	_, err = request.SetBody(map[string]interface{}{"name": "foo"}).SetResult(newGoal).Post("/goal")
	assert.NoError(t, err)
	assert.NotNil(t, newGoal.ID)
	assert.Equal(t, "foo", newGoal.Name)
	assert.NotNil(t, newGoal.DateCreated)

	// Check if the new goal is added.
	_, err = request.SetResult(goals).Get("/goal")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(*goals))

	for _, g := range *goals {
		assert.Equal(t, newGoal.ID, g.ID)
		assert.Equal(t, newGoal.Name, g.Name)
		assert.NotNil(t, newGoal.DateCreated)
	}

	// Get the new goal.
	goal := new(model.Goal)
	_, err = request.SetResult(goal).Get("/goal/" + newGoal.ID)
	assert.NoError(t, err)
	assert.Equal(t, newGoal.ID, goal.ID)
	assert.Equal(t, newGoal.Name, goal.Name)
	assert.NotNil(t, newGoal.DateCreated)

	// Update the new goal.
	putGoal := new(model.Goal)
	_, err = request.SetBody(map[string]interface{}{"name": "bar"}).SetResult(putGoal).Put("/goal/" + newGoal.ID)
	assert.NoError(t, err)
	assert.NotNil(t, putGoal.ID)
	assert.Equal(t, "bar", putGoal.Name)
	assert.NotNil(t, putGoal.DateCreated)
	assert.NotEqual(t, newGoal.Name, putGoal.Name)
	assert.Equal(t, newGoal.ID, putGoal.ID)

	// Delete the new goal.
	status := &struct {
		Message string
	}{}
	_, err = request.SetResult(status).Delete("/goal/" + newGoal.ID)
	assert.NoError(t, err)
	assert.Equal(t, "deleted successfully", status.Message)

	// Check if there is nothing in the system now.
	_, err = request.SetResult(goals).Get("/goal")
	assert.NoError(t, err)
	assert.Equal(t, 0, len(*goals))

	// Catch the error.
	_, err = request.SetError(status).Get("/goal/" + newGoal.ID)
	assert.NoError(t, err)
	assert.Equal(t, "sql: no rows in result set", status.Message)

	// _, err = request.SetError(status).Post("/goal")
	// assert.NoError(t, err)
	// assert.Equal(t, "sql: no rows in result set", status.Message)

	_, err = request.SetBody(map[string]interface{}{"name": "foobar"}).SetError(putGoal).Put("/goal/" + newGoal.ID)
	assert.NoError(t, err)
	assert.Equal(t, "sql: no rows in result set", status.Message)

	_, err = request.SetError(status).Delete("/goal/" + newGoal.ID)
	assert.NoError(t, err)
	assert.Equal(t, "sql: no rows in result set", status.Message)
}
