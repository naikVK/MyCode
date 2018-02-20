package email

import "testing"

func TestSendmail(t *testing.T) {
	type args struct {
		toAddress        string
		body             string
		emailservicetype string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "Email test",
			args: args{
				toAddress:        "priyankam@mkcl.org",
				body:             "Verification",
				emailservicetype: "otpmsg",
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Sendmail(tt.args.toAddress, tt.args.body, tt.args.emailservicetype); got != tt.want {
				t.Errorf("Sendmail() = %v, want %v", got, tt.want)
			}
		})
	}
}
