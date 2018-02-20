package model

// otp keeps information of all otp
type OTP struct {
	Username string `json:"username,omitempty" bson:"USERNAME" validate:"required"`

	OTP           string `json:"otp,omitempty" bson:"OTP" validate:"required"`
	OTPVerifiedOn string `json:"otpverifiedon,omitempty" bson:"OTP_VERIFIED_ON,omitempty"`
}
