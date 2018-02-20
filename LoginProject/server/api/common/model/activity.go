package model

// Login keeps information of all Logins
type ActivityLog struct {
	Username       string `json:"username,omitempty" bson:"ACITIVITY_FOR" validate:"required"`
	ActivitType    string `json:"activitytype,omitempty" bson:"ACITIVITY_TYPE" validate:"required"`
	ActivityResult string `json:"activityresult,omitempty" bson:"ACTIVITY_RESULT" validate:"required"`
	ActivityOn     string `json:"activityon,omitempty" bson:"ACTIVITY_ON" validate:"required"`
	ActivityBy     string `json:"activityby,omitempty" bson:"ACTIVITY_BY" validate:"required"`
}
