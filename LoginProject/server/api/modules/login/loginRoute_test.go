// NOTE: SET -->	viper.SetConfigName("../../../config")

package login

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

/*
//correct Username

func Test1ValidateUsernameRoute(test *testing.T) {
	e := echo.New()
	userJSON := `{"username":"sushmita"}`

	req, _ := http.NewRequest(echo.POST, "o/login/validateusername", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, ValidateUsernameRoute(c)) {
		assert.Equal(test, http.StatusOK, rec.Code)
	}
}

//BLANK Username
func Test2ValidateUsernameRoute(test *testing.T) {
	e := echo.New()
	userJSON := `{"username":""}`

	req, _ := http.NewRequest(echo.POST, "o/login/validateusername", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, ValidateUsernameRoute(c)) {
		assert.Equal(test, http.StatusExpectationFailed, rec.Code)
	}
}

//NIL JSON
func Test3ValidateUsernameRoute(test *testing.T) {
	e := echo.New()
	req, _ := http.NewRequest(echo.POST, "o/login/validateusername", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, ValidateUsernameRoute(c)) {
		assert.Equal(test, http.StatusExpectationFailed, rec.Code)
	}
}

//WRONG Username
func Test4ValidateUsernameRoute(test *testing.T) {
	e := echo.New()
	userJSON := `{"username":"QWERTYUI"}`

	req, _ := http.NewRequest(echo.POST, "o/login/validateusername", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, ValidateUsernameRoute(c)) {
		assert.Equal(test, http.StatusExpectationFailed, rec.Code)
	}
}


//CORRECT USERNAME & PASSWORD
func Test5ValidatePasswordRoute(test *testing.T) {
	e := echo.New()
	userJSON := `{"username":"sushmita","password": "1234"}`

	req, _ := http.NewRequest(echo.POST, "/o/login/validateuser", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("ClientId", "ERA_1")

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, ValidatePasswordRoute(c)) {
		assert.Equal(test, http.StatusOK, rec.Code)
	}
}

//	WRONG PASSWORD
func Test6ValidatePasswordRoute(test *testing.T) {
	e := echo.New()
	userJSON := `{"username":"sushmita","password": "1234ERTYUIO"}`

	req, _ := http.NewRequest(echo.POST, "/o/login/validateuser", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("ClientId", "ERA_1")

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, ValidatePasswordRoute(c)) {
		assert.Equal(test, http.StatusExpectationFailed, rec.Code)
	}
}

//	NIL JSON
func Test7ValidatePasswordRoute(test *testing.T) {
	e := echo.New()
	req, _ := http.NewRequest(echo.POST, "/o/login/validateuser", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("ClientId", "ERA_1")

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, ValidatePasswordRoute(c)) {
		assert.Equal(test, http.StatusExpectationFailed, rec.Code)
	}
}

//	EMPTY USERNAME
func Test8ValidatePasswordRoute(test *testing.T) {
	e := echo.New()
	userJSON := `{"username":"","password": "1234"}`

	req, _ := http.NewRequest(echo.POST, "/o/login/validateuser", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("ClientId", "ERA_1")

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, ValidatePasswordRoute(c)) {
		assert.Equal(test, http.StatusExpectationFailed, rec.Code)
	}
}

// WRONG CLIENTID
func Test9ValidatePasswordRoute(test *testing.T) {
	redisSessionManager.Init()

	e := echo.New()
	userJSON := `{"username":"sushmita","password": "1234"}`

	req, _ := http.NewRequest(echo.POST, "/o/login/validateuser", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("ClientId", "SOLAR_2")

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, ValidatePasswordRoute(c)) {
		assert.Equal(test, http.StatusOK, rec.Code)
	}
}

// WRONG CLIENTID
func Test10ValidatePasswordRoute(test *testing.T) {
	e := echo.New()
	userJSON := `{"username":"sushmita","password": "1234"}`

	req, _ := http.NewRequest(echo.POST, "/o/login/validateuser", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("ClientId", "")

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, ValidatePasswordRoute(c)) {
		assert.Equal(test, http.StatusBadRequest, rec.Code)
	}
}


//NOTE: GoogleAuthAppRoute
//CORRECT USERNAME
func Test11GoogleAuthAppRoute(test *testing.T) {

	e := echo.New()
	o := e.Group("/o")
	r := e.Group("/r")
	c := r.Group("/c")
	middleware.Init(e, o, r, c)
	userJSON := `{"username":"sushmita"}`

	req, _ := http.NewRequest(echo.POST, "/o/login/googleauthapp", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c1 := e.NewContext(req, rec)

	if assert.NoError(test, GoogleAuthApp(c1)) {
		assert.Equal(test, http.StatusOK, rec.Code)
	}
}

// NIL JSON
func Test12GoogleAuthAppRoute(test *testing.T) {
	e := echo.New()
	req, _ := http.NewRequest(echo.POST, "/o/login/googleauthapp", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, GoogleAuthApp(c)) {
		assert.Equal(test, http.StatusExpectationFailed, rec.Code)
	}
}
*/
//NOTE: GoogleAuthAppCheck
//CORRECT USERNAME & OTP
// func Test13GoogleAuthAppCheck(test *testing.T) {
// 	e := echo.New()
// 	userJSON := `{"username":"imti","otp":"767818"}`

// 	req, _ := http.NewRequest(echo.POST, "/o/login/googleauthappcheck", strings.NewReader(userJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set(echo.HeaderAuthorization, "Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImltdGkiLCJpc0FkbWluIjpmYWxzZSwic2Vzc2lvbklkIjoiaW10aTp5Y3Fsd1hpTiIsImNsaWVudElkIjoiRVJBXzEiLCJleHAiOjE1MTE1MDUwMDh9.R0f9cnnHt7R99Y_aolUX2sd24OjDOoKWAJOH8CZv8U8")

// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)

// 	if assert.NoError(test, GoogleAuthAppCheck(c)) {
// 		assert.Equal(test, http.StatusOK, rec.Code)
// 	}
// }

// //WRONG  OTP
// func Test14GoogleAuthAppCheck(test *testing.T) {
// 	e := echo.New()
// 	userJSON := `{"username":"sushmita","otp":"431384"}`

// 	req, _ := http.NewRequest(echo.POST, "/o/login/googleauthappcheck", strings.NewReader(userJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)

// 	if assert.NoError(test, GoogleAuthAppCheck(c)) {
// 		assert.Equal(test, http.StatusExpectationFailed, rec.Code)
// 	}
// }

// //  NIL JSON
// func Test15GoogleAuthAppCheck(test *testing.T) {
// 	e := echo.New()
// 	req, _ := http.NewRequest(echo.POST, "/o/login/googleauthappcheck", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)

// 	if assert.NoError(test, GoogleAuthAppCheck(c)) {
// 		assert.Equal(test, http.StatusExpectationFailed, rec.Code)
// 	}
// }

//NOTE: GoogleAuthGetKey
//CORRECT USERNAME
// func Test16GoogleAuthGetKey(test *testing.T) {
// 	e := echo.New()
// 	userJSON := `{"username":"sushmita"}`

// 	req, _ := http.NewRequest(echo.POST, "/o/login/googleauthgetkey", strings.NewReader(userJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)

// 	if assert.NoError(test, GoogleAuthGetKey(c)) {
// 		assert.Equal(test, http.StatusOK, rec.Code)
// 	}
// }

// //WRONG USERNAME
// func Test17GoogleAuthGetKey(test *testing.T) {
// 	e := echo.New()
// 	userJSON := `{"username":"SDFGUJ"}`

// 	req, _ := http.NewRequest(echo.POST, "/o/login/googleauthgetkey", strings.NewReader(userJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)

// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)

// 	if assert.NoError(test, GoogleAuthGetKey(c)) {
// 		assert.Equal(test, http.StatusExpectationFailed, rec.Code)
// 	}
// }

// //NIL JSON
// func Test171GoogleAuthGetKey(test *testing.T) {
// 	e := echo.New()

// 	req, _ := http.NewRequest(echo.POST, "/o/login/googleauthgetkey", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)

// 	if assert.NoError(test, GoogleAuthGetKey(c)) {
// 		assert.Equal(test, http.StatusExpectationFailed, rec.Code)
// 	}
// }

//NOTE: sendOTPonPhone
//CORRECT USERNAME
// func Test16sendOTPonPhone(test *testing.T) {
// 	e := echo.New()
// 	userJSON := `{"username":"sushmita"}`

// 	req, _ := http.NewRequest(echo.POST, "/o/login/sendOTPonPhone", strings.NewReader(userJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set("ClientId", "ERA_1")

// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)

// 	if assert.NoError(test, sendOTPonPhone(c)) {
// 		assert.Equal(test, http.StatusOK, rec.Code)
// 	}
// }

// //WRONG USERNAME
// func Test17sendOTPonPhone(test *testing.T) {
// 	e := echo.New()
// 	userJSON := `{"username":"sushmitaSDFGHLK"}`

// 	req, _ := http.NewRequest(echo.POST, "/o/login/sendOTPonPhone", strings.NewReader(userJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set("ClientId", "ERA_1")

// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)

// 	if assert.NoError(test, sendOTPonPhone(c)) {
// 		assert.Equal(test, http.StatusExpectationFailed, rec.Code)
// 	}
// }

// //EMPTY USERNAME
// func Test18sendOTPonPhone(test *testing.T) {
// 	e := echo.New()
// 	userJSON := `{"username":""}`

// 	req, _ := http.NewRequest(echo.POST, "/o/login/sendOTPonPhone", strings.NewReader(userJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set("ClientId", "ERA_1")

// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)

// 	if assert.NoError(test, sendOTPonPhone(c)) {
// 		assert.Equal(test, http.StatusExpectationFailed, rec.Code)
// 	}
// }

// //NIL JSON
// func Test19sendOTPonPhone(test *testing.T) {
// 	e := echo.New()
// 	req, _ := http.NewRequest(echo.POST, "/o/login/sendOTPonPhone", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set("ClientId", "ERA_1")

// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)

// 	if assert.NoError(test, sendOTPonPhone(c)) {
// 		assert.Equal(test, http.StatusExpectationFailed, rec.Code)
// 	}
// }

//NOTE: VerifyOTP
//CORRECT  USERNAME & OTP
// func Test19VerifyOTP(test *testing.T) {
// 	e := echo.New()
// 	userJSON := `{"username":"sushmita","otp":"602407"}`

// 	req, _ := http.NewRequest(echo.POST, "/o/login/verifyOTP", strings.NewReader(userJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set("ClientId", "ERA_1")

// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)

// 	if assert.NoError(test, VerifyOTP(c)) {
// 		assert.Equal(test, http.StatusOK, rec.Code)
// 	}
// }

// //WRONG  USERNAME & OTP
// func Test20VerifyOTP(test *testing.T) {
// 	e := echo.New()
// 	userJSON := `{"username":"sushmitaFSDF","otp":"264010"}`

// 	req, _ := http.NewRequest(echo.POST, "/o/login/verifyOTP", strings.NewReader(userJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set("ClientId", "ERA_1")

// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)

// 	if assert.NoError(test, VerifyOTP(c)) {
// 		assert.Equal(test, http.StatusExpectationFailed, rec.Code)
// 	}
// }

// // NIL JSON
// func Test21VerifyOTP(test *testing.T) {
// 	e := echo.New()
// 	req, _ := http.NewRequest(echo.POST, "/o/login/verifyOTP", nil)
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set("ClientId", "ERA_1")

// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)

// 	if assert.NoError(test, VerifyOTP(c)) {
// 		assert.Equal(test, http.StatusExpectationFailed, rec.Code)
// 	}
// }

// //WRONG OTP
// func Test22VerifyOTP(test *testing.T) {
// 	e := echo.New()
// 	userJSON := `{"username":"sushmita","otp":"860481"}`

// 	req, _ := http.NewRequest(echo.POST, "/o/login/verifyOTP", strings.NewReader(userJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set("ClientId", "ERA_1")

// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)

// 	if assert.NoError(test, VerifyOTP(c)) {
// 		assert.Equal(test, http.StatusExpectationFailed, rec.Code)
// 	}
// }

// //WRONG CLIENTID
// func Test23VerifyOTP(test *testing.T) {
// 	e := echo.New()
// 	userJSON := `{"username":"sushmita","otp":"602407"}`

// 	req, _ := http.NewRequest(echo.POST, "/o/login/verifyOTP", strings.NewReader(userJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set("ClientId", "SOLAR_1")

// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)

// 	if assert.NoError(test, VerifyOTP(c)) {
// 		assert.Equal(test, http.StatusBadRequest, rec.Code)
// 	}
// }

// //DIFFERENT CLIENT
// func Test24VerifyOTP(test *testing.T) {
// 	e := echo.New()
// 	redisSessionManager.Init()

// 	userJSON := `{"username":"sushmita","otp":"602407"}`

// 	req, _ := http.NewRequest(echo.POST, "/o/login/verifyOTP", strings.NewReader(userJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set("ClientId", "SOLAR_2")

// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)

// 	if assert.NoError(test, VerifyOTP(c)) {
// 		assert.Equal(test, http.StatusOK, rec.Code)
// 	}
// }
// func Test25VerifyOTP(test *testing.T) {
// 	e := echo.New()
// 	userJSON := `{"username":"","otp":"264010"}`

// 	req, _ := http.NewRequest(echo.POST, "/o/login/verifyOTP", strings.NewReader(userJSON))
// 	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
// 	req.Header.Set("ClientId", "ERA_1")

// 	rec := httptest.NewRecorder()

// 	c := e.NewContext(req, rec)

// 	if assert.NoError(test, VerifyOTP(c)) {
// 		assert.Equal(test, http.StatusExpectationFailed, rec.Code)
// 	}
// }

//NOTE: sendOTPonEmail
//CORRECT USERNAME
func Test26sendOTPonEmail(test *testing.T) {
	e := echo.New()
	userJSON := `{"username":"sushmita"}`

	req, _ := http.NewRequest(echo.POST, "/o/login/sendOTPonEmail", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("ClientId", "ERA_1")

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, sendOTPonEmail(c)) {
		assert.Equal(test, http.StatusOK, rec.Code)
	}
}

//WRONG USERNAME
func Test27sendOTPonEmail(test *testing.T) {
	e := echo.New()
	userJSON := `{"username":"sushmitaGYUII"}`

	req, _ := http.NewRequest(echo.POST, "/o/login/sendOTPonEmail", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("ClientId", "ERA_1")

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, sendOTPonEmail(c)) {
		assert.Equal(test, http.StatusExpectationFailed, rec.Code)
	}
}

//NIL JSON
func Test28sendOTPonEmail(test *testing.T) {
	e := echo.New()
	req, _ := http.NewRequest(echo.POST, "/o/login/sendOTPonEmail", nil)
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("ClientId", "ERA_1")

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, sendOTPonEmail(c)) {
		assert.Equal(test, http.StatusExpectationFailed, rec.Code)
	}
}

//EMPTY USERNAME
func Test29sendOTPonEmail(test *testing.T) {
	e := echo.New()
	userJSON := `{"username":""}`

	req, _ := http.NewRequest(echo.POST, "/o/login/sendOTPonEmail", strings.NewReader(userJSON))
	req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationJSON)
	req.Header.Set("ClientId", "ERA_1")

	rec := httptest.NewRecorder()

	c := e.NewContext(req, rec)

	if assert.NoError(test, sendOTPonEmail(c)) {
		assert.Equal(test, http.StatusExpectationFailed, rec.Code)
	}
}
