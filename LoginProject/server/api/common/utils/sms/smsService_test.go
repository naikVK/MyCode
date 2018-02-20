package sms

import "testing"

func Test_sendSingleSMS(t *testing.T) {
	type args struct {
		message        string
		mobileno       string
		smsservicetype string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "SMS service test case",
			args: args{
				message:        "",
				mobileno:       "7709191781",
				smsservicetype: "otpmsg",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			SendSingleSMS(tt.args.message, tt.args.mobileno, tt.args.smsservicetype)
		})
	}
}
