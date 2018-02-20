package clientConfiguration

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

//TEST FOR INIT
func TestInit(test *testing.T) {
	e := echo.New()
	o := e.Group("/o")

	r := e.Group("/r")

	c := r.Group("/c")

	Init(o, r, c)
}

//NOTE: TEST CASE FOR CORRECT CREDENTIALS
func Test1GetClientConfigRoute(test *testing.T) {
	e := echo.New()
	userJSON := `{"clientid":"ERA_1"}`

	req, _ := http.NewRequest(echo.POST, "/getclientconfig", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, GetClientConfigRoute(c)) {
		assert.Equal(test, http.StatusOK, rec.Code)
	}
}
