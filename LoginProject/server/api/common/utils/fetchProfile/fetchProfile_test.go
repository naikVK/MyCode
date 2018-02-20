package fetchProfile

import (
	"LoginProject/server/api/common/model"
	"reflect"
	"testing"
)

func TestGetByUserName(t *testing.T) {
	type args struct {
		username string
	}
	tests := []struct {
		name  string
		args  args
		want  model.ProfileDetail
		want1 bool
	}{
		{
			name: "FetchProfile Test",
			args:args{
				username:"priyanka",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := GetByUserName(tt.args.username)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetByUserName() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("GetByUserName() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
