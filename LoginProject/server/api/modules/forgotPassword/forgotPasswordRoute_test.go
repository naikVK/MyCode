package forgotPassword

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func Test1Init(test *testing.T) {
	e := echo.New()
	o := e.Group("/o")
	r := e.Group("/r")
	c := r.Group("/c")
	Init(o, r, c)
}

func Test1SendOTPRoute(test *testing.T) {
	e := echo.New()
	userJSON := `{"username":"priyankam"}`

	req, _ := http.NewRequest(echo.POST, "o/login/sendotp", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, SendOTPRoute(c)) {
		assert.Equal(test, http.StatusOK, rec.Code)
	}
}
