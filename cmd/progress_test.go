/*
	These tests are designed to run against an instance of Progress with
	a PostgreSQL data store backing configured.

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

func TestE2ESuccess(t *testing.T) {
	// Get all goals.
	goals := new([]model.Goal)
	_, err := request.SetResult(goals).Get("/goal")
	assert.NoError(t, err)
	assert.Equal(t, 0, len(*goals))

	// Get all systems.
	systems := new([]model.System)
	_, err = request.SetResult(systems).Get("/system")
	assert.NoError(t, err)
	assert.Equal(t, 0, len(*systems))

	// Create a new goal.
	newGoal := new(model.Goal)
	_, err = request.SetBody(map[string]interface{}{"name": "new goal"}).SetResult(newGoal).Post("/goal")
	assert.NoError(t, err)
	assert.NotNil(t, newGoal.ID)
	assert.Equal(t, "new goal", newGoal.Name)
	assert.NotNil(t, newGoal.DateCreated)

	// Create a new system.
	newSystem := new(model.System)
	_, err = request.SetBody(map[string]interface{}{"goal_id": newGoal.ID, "name": "new system", "repeat": "everyday"}).SetResult(newSystem).Post("/system")
	assert.NoError(t, err)
	assert.NotNil(t, newSystem.ID)
	assert.Equal(t, newGoal.ID, newSystem.GoalID)
	assert.Equal(t, "new system", newSystem.Name)
	assert.Equal(t, "everyday", newSystem.Repeat)
	assert.NotNil(t, newSystem.DateCreated)

	// Check if the new goal is added.
	_, err = request.SetResult(goals).Get("/goal")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(*goals))

	for _, g := range *goals {
		assert.Equal(t, newGoal.ID, g.ID)
		assert.Equal(t, newGoal.Name, g.Name)
		assert.NotNil(t, g.DateCreated)
	}

	// Check if the new system is added.
	_, err = request.SetResult(systems).Get("/system")
	assert.NoError(t, err)
	assert.Equal(t, 1, len(*systems))

	for _, s := range *systems {
		assert.Equal(t, newSystem.ID, s.ID)
		assert.Equal(t, newSystem.GoalID, s.GoalID)
		assert.Equal(t, newSystem.Name, s.Name)
		assert.Equal(t, newSystem.Repeat, s.Repeat)
		assert.NotNil(t, s.DateCreated)
	}

	// Get the new goal.
	goal := new(model.Goal)
	_, err = request.SetResult(goal).Get("/goal/" + newGoal.ID)
	assert.NoError(t, err)
	assert.Equal(t, newGoal.ID, goal.ID)
	assert.Equal(t, newGoal.Name, goal.Name)
	assert.NotNil(t, goal.DateCreated)

	// Get the new system.
	system := new(model.System)
	_, err = request.SetResult(system).Get("/system/" + newSystem.ID)
	assert.NoError(t, err)
	assert.Equal(t, newSystem.ID, system.ID)
	assert.Equal(t, newSystem.GoalID, system.GoalID)
	assert.Equal(t, newSystem.Name, system.Name)
	assert.Equal(t, newSystem.Repeat, system.Repeat)
	assert.NotNil(t, system.DateCreated)

	// Update the new goal.
	putGoal := new(model.Goal)
	_, err = request.SetBody(map[string]interface{}{"name": "updated goal"}).SetResult(putGoal).Put("/goal/" + newGoal.ID)
	assert.NoError(t, err)
	assert.Equal(t, newGoal.ID, putGoal.ID)
	assert.Equal(t, "updated goal", putGoal.Name)
	assert.NotNil(t, putGoal.DateCreated)

	// Update the new system.
	putSystem := new(model.System)
	_, err = request.SetBody(map[string]interface{}{"goal_id": newGoal.ID, "name": "updated system", "repeat": "every week"}).SetResult(putSystem).Put("/system/" + newSystem.ID)
	assert.NoError(t, err)
	assert.Equal(t, newSystem.ID, putSystem.ID)
	assert.Equal(t, newGoal.ID, putSystem.GoalID)
	assert.Equal(t, "updated system", putSystem.Name)
	assert.Equal(t, "every week", putSystem.Repeat)
	assert.NotNil(t, putGoal.DateCreated)

	// Delete the new goal.
	status := &struct {
		Message string
	}{}
	_, err = request.SetResult(status).Delete("/goal/" + newGoal.ID)
	assert.NoError(t, err)
	assert.Equal(t, "deleted successfully", status.Message)

	// Delete the new system.
	_, err = request.SetResult(status).Delete("/system/" + newSystem.ID)
	assert.NoError(t, err)
	assert.Equal(t, "deleted successfully", status.Message)

	// Check if there is no goal.
	_, err = request.SetResult(goals).Get("/goal")
	assert.NoError(t, err)
	assert.Equal(t, 0, len(*goals))

	// Check if there is no system.
	_, err = request.SetResult(systems).Get("/system")
	assert.NoError(t, err)
	assert.Equal(t, 0, len(*systems))
}

func TestE2EGoalError(t *testing.T) {
	status := &struct {
		Message string
	}{}
	_, err := request.SetError(status).Get("/goal/a235be9e-ab5d-44e6-a987-fa1c749264c7")
	assert.NoError(t, err)
	assert.Equal(t, "sql: no rows in result set", status.Message)

	_, err = request.SetBody(map[string]interface{}{"foo": "bar"}).SetError(status).Post("/goal")
	assert.NoError(t, err)
	assert.Equal(t, "Key: 'NewGoal.Name' Error:Field validation for 'Name' failed on the 'required' tag", status.Message)

	_, err = request.SetBody(map[string]interface{}{"name": "foo"}).SetError(status).Put("/goal/a235be9e-ab5d-44e6-a987-fa1c749264c7")
	assert.NoError(t, err)
	assert.Equal(t, "sql: no rows in result set", status.Message)

	_, err = request.SetError(status).Delete("/goal/a235be9e-ab5d-44e6-a987-fa1c749264c7")
	assert.NoError(t, err)
	assert.Equal(t, "sql: no rows in result set", status.Message)
}
