package test

import (
	"GoProject/main/enum"
	"GoProject/main/util"
	"testing"
)

func TestParseToken(t *testing.T) {
	type args struct {
		tokenString string
	}

	tests := []struct {
		name string
		args args
		want *util.ClaimsToken
	}{
		{"test1", args{util.SignToken(1, "admin", enum.Admin)}, nil},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := util.ParseToken(tt.args.tokenString)
			if err != nil {
				t.Errorf("ParseToken() error = %v", err)
			} else {
				t.Log(got)
			}

		})
	}
}

func TestSignToken(t *testing.T) {
	type args struct {
		id   int
		name string
		auth enum.Authority
	}

	tests := []struct {
		name string
		args args
		want string
	}{
		{"test1", args{1, "admin", enum.Admin}, ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := util.SignToken(tt.args.id, tt.args.name, tt.args.auth); got == tt.want {
				t.Errorf("SignToken() = %v, want %v", got, tt.want)
			} else {
				t.Log(got)
			}
		})
	}
}
