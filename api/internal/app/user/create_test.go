package user_test

import (
	"strings"
	"testing"

	"github.com/doktorupnos/crow/api/internal/app/user"
)

func TestCreateRequest_Validate(t *testing.T) {
	cases := map[string]struct {
		req  user.CreateRequest
		want error
	}{
		"empty name": {
			req:  user.CreateRequest{Name: "", Password: ""},
			want: user.ErrNameEmpty,
		},
		"empty password": {
			req:  user.CreateRequest{Name: "user", Password: ""},
			want: user.ErrPasswordEmpty,
		},
		"name is too small": {
			req:  user.CreateRequest{Name: "ab", Password: "password"},
			want: user.ErrNameSmall,
		},
		"name is too big": {
			req:  user.CreateRequest{Name: "Maximiliano_Featherstone", Password: "password"},
			want: user.ErrNameBig,
		},
		"password is too big": {
			req: user.CreateRequest{
				Name:     "user",
				Password: strings.Repeat("a", 100),
			},
			want: user.ErrPasswordBig,
		},
		"name containing spaces": {
			req:  user.CreateRequest{Name: "Peter Parker", Password: "password"},
			want: user.ErrNameMalformed,
		},
		"password containing not allowed special characters": {
			req:  user.CreateRequest{Name: "user", Password: "(Correct) {Horse} ~Battery `Staple`"},
			want: user.ErrPasswordMalformed,
		},
	}

	for name, tt := range cases {
		t.Run(name, func(t *testing.T) {
			err := tt.req.Validate()
			if err != tt.want {
				t.Errorf("Validate(%#v) = %v, want %v", tt.req, err, tt.want)
			}
		})
	}

	t.Run("valid request returns no error", func(t *testing.T) {
		req := user.CreateRequest{
			Name:     "user",
			Password: "Correct Horse Battery Staple",
		}

		err := req.Validate()
		if err != nil {
			t.Errorf("Validate(%#v) unexpected error: %v", req, err)
		}
	})
}
