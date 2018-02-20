package validations

import (
	"regexp"
	"strconv"

	"corelab.mkcl.org/MKCLOS/coredevelopmentplatform/coreospackage/logginghelper"
)

func IsEmpty(fields ...string) bool {
	logginghelper.LogInfo("Inside Validation util:: IsEmpty")
	for _, field := range fields {
		if len(field) == 0 {
			return true
		}
	}
	return false
}

func IsNumber(fields ...string) bool {
	logginghelper.LogInfo("Inside Validation util:: IsNumber")
	for _, field := range fields {
		_, err := strconv.Atoi(field)
		if err != nil {
			return false
		}
	}
	return true
}
func IsMobileNumber(field string) bool {
	logginghelper.LogInfo("Inside Validation util:: isMobileNumber")
	matched, _ := regexp.MatchString("^[7-9][0-9]{9}$", field)
	if matched {
		return true
	} else {
		return false
	}
}

func IsValidEmail(email string) bool {
	logginghelper.LogInfo("Inside Validation util:: IsValidEmail")
	// matched, _ := regexp.MatchString(`^([\w\.\_]{2,10})@(\w{1,}).([a-z]{2,4})$`, email)
	matched, _ := regexp.MatchString(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`, email)

	if matched {
		// valid email
		return true
	} else {
		// Invalid email
		return false
	}
}
