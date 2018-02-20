package registration

//TEST FOR INIT
import (
	"LoginProject/server/api/common/model"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/labstack/echo"
	"github.com/stretchr/testify/assert"
)

func TestInit(test *testing.T) {
	e := echo.New()
	o := e.Group("/o")

	r := e.Group("/r")

	c := r.Group("/c")

	Init(o, r, c)
}

//NOTE: TEST CASE FOR CORRECT CREDENTIALS
func Test1IsUsernameAvailable(test *testing.T) {
	e := echo.New()
	// userJSON := `{"clientid":"ERA_1"}`
	profileDetail := model.ProfileDetail{}
	profileDetail.UserName = "vivek1234"
	profileDetail.PersonalDetails.FullName = "vivek naik"
	profileDetail.PersonalDetails.Dob = "2017-11-02"
	profileDetailbyte, _ := json.Marshal(profileDetail)

	req, _ := http.NewRequest(echo.POST, "/isUsernameAvailable", strings.NewReader(string(profileDetailbyte)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, IsUsernameAvailable(c)) {
		assert.Equal(test, http.StatusOK, rec.Code)
	}
}

func Test2IsUsernameAvailable(test *testing.T) {
	e := echo.New()
	// userJSON := `{"clientid":"ERA_1"}`
	profileDetail := model.ProfileDetail{}
	profileDetail.UserName = "vivekn"
	profileDetail.PersonalDetails.FullName = "vivek naik"
	profileDetail.PersonalDetails.Dob = "2017-11-02"
	profileDetailbyte, _ := json.Marshal(profileDetail)

	req, _ := http.NewRequest(echo.POST, "/isUsernameAvailable", strings.NewReader(string(profileDetailbyte)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, IsUsernameAvailable(c)) {
		assert.Equal(test, http.StatusAlreadyReported, rec.Code)
	}
}

// empty username
func Test3IsUsernameAvailable(test *testing.T) {
	e := echo.New()
	// userJSON := `{"clientid":"ERA_1"}`
	profileDetail := model.ProfileDetail{}
	profileDetail.UserName = ""
	profileDetail.PersonalDetails.FullName = "vivek naik"
	profileDetail.PersonalDetails.Dob = "2017-11-02"
	profileDetailbyte, _ := json.Marshal(profileDetail)

	req, _ := http.NewRequest(echo.POST, "/isUsernameAvailable", strings.NewReader(string(profileDetailbyte)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, IsUsernameAvailable(c)) {
		assert.Equal(test, http.StatusBadRequest, rec.Code)
	}
}

//validation
func Test1RegisterUser(test *testing.T) {
	e := echo.New()
	// userJSON := `{"clientid":"ERA_1"}`
	profileDetail := model.ProfileDetail{}
	profileDetail.PersonalDetails.FullName = "vivek naik"
	profileDetail.PersonalDetails.Gender = "M"
	profileDetail.PersonalDetails.Dob = "2017-11-02"
	profileDetail.ContactDetails.Email.Address = "viveknaik@mkcl.org"
	profileDetail.ContactDetails.Mobile.Number = "7709191781"
	profileDetail.UserName = "vk"
	profileDetail.Password = "vknaik"
	profileDetailbyte, _ := json.Marshal(profileDetail)

	req, _ := http.NewRequest(echo.POST, "/register", strings.NewReader(string(profileDetailbyte)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, RegisterRoute(c)) {
		assert.Equal(test, http.StatusBadRequest, rec.Code)
	}
}

// username already present
func Test2RegisterUser(test *testing.T) {
	e := echo.New()
	// userJSON := `{"clientid":"ERA_1"}`
	profileDetail := model.ProfileDetail{}

	profileDetail.PersonalDetails.FullName = "vivek naik"
	profileDetail.PersonalDetails.Gender = "M"
	profileDetail.PersonalDetails.Dob = "2017-11-02"
	profileDetail.ContactDetails.Email.Address = "viveknaik@mkcl.org"
	profileDetail.ContactDetails.Mobile.Number = "7709191781"
	profileDetail.UserName = "vivekn"
	profileDetail.Password = "vknaik"
	profileDetailbyte, _ := json.Marshal(profileDetail)

	req, _ := http.NewRequest(echo.POST, "/register", strings.NewReader(string(profileDetailbyte)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, RegisterRoute(c)) {
		assert.Equal(test, http.StatusBadRequest, rec.Code)
	}
}

// success
func Test3RegisterUser(test *testing.T) {
	e := echo.New()
	// userJSON := `{"clientid":"ERA_1"}`
	profileDetail := model.ProfileDetail{}

	profileDetail.PersonalDetails.FullName = "vivek naik"
	profileDetail.PersonalDetails.Gender = "M"
	profileDetail.PersonalDetails.Dob = "2017-11-02"
	profileDetail.ContactDetails.Email.Address = "viveknaik@mkcl.org"
	profileDetail.ContactDetails.Mobile.Number = "7709191781"
	profileDetail.UserName = "naikVivek0071"
	profileDetail.Password = "vknaik"
	profileDetailbyte, _ := json.Marshal(profileDetail)

	req, _ := http.NewRequest(echo.POST, "/register", strings.NewReader(string(profileDetailbyte)))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, RegisterRoute(c)) {
		assert.Equal(test, http.StatusOK, rec.Code)
	}
}
