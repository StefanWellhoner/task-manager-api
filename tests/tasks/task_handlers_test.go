package handlers_test

import (
    "net/http"
    "net/http/httptest"
    "testing"

    "github.com/StefanWellhoner/task-manager-api/internal/app"
    "github.com/stretchr/testify/assert"
)

func TestGetTasks(t *testing.T) {
    router := app.SetupRouter()

    w := httptest.NewRecorder()
    req, _ := http.NewRequest("GET", "/v1/tasks", nil)
    router.ServeHTTP(w, req)

    assert.Equal(t, 200, w.Code)
}