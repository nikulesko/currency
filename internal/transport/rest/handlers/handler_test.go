package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

//todo: Doesn’t works
func TestGetLatest(t *testing.T) {
	// Setup
	e := echo.New()
	req := httptest.NewRequest(http.MethodGet, "/currency/latest", nil)
	rec := httptest.NewRecorder()
	c := e.NewContext(req, rec)

	// Assertions
	if assert.NoError(t, GetLatest(c)) {
		assert.Equal(t, http.StatusOK, rec.Code)
		//assert.Equal(t, userJSON, rec.Body.String())
	}
}
