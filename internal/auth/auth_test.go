package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

func TestHashPassword(t *testing.T) {
	cases := []struct {
		input       string
		expectedOut string
		expectedErr error
	}{
		{
			input:       "Rttf**65_p3ter",
			expectedOut: "$2a$12$.uXkotxlzOGVEO3IQJuzPejvuUzZqmc2PlX3UVS8VhSJQZ4.qfDQO",
			expectedErr: nil,
		},
		{
			input:       "descartes&14plaTo",
			expectedOut: "$2a$12$vdNDttFuh2Te2TK3kxWj0eDAC.drsCG1yUxeCYLz45VoxXzhN.3G6",
			expectedErr: nil,
		},
		{
			input:       "tnanever2015-Eht",
			expectedOut: "$2a$12$wOGn.W7ARE5x/3IBsL5Xm..fclin4OuIjJXbNLsWuwYQKb5p34GaG",
			expectedErr: nil,
		},
	}

	for _, c := range cases {
		actual, err := HashPassword(c.input)
		if err != c.expectedErr && actual != c.expectedOut {
			t.Errorf("Test Failed => actual output: %s is not same as expected output: %s", actual, c.expectedOut)
			continue
		}
	}
}

func TestCheckPasswordHash(t *testing.T) {
	cases := []struct {
		inputHash     string
		inputPassword string
		hasError      bool
	}{
		{
			inputHash:     "$2a$12$HB9Rxi8ZNZTiTkjj4Ydko.GBxc10..OV915/ahJr7lH5mRodQ7K5i",
			inputPassword: "X9f$2pL!zW",
			hasError:      false,
		},
		{
			inputHash:     "$2a$12$Qc3ALQkOA8cpWd.cR5CAiOY25fpMvRWYT9PAxK.28DTKyCbR6C/t1",
			inputPassword: "mE7#vT@kQ4",
			hasError:      true,
		},
		{
			inputHash:     "$2a$12$IpnUQ.hUQnVzGzKVEqHpo.OxPRaJ31qSoxs08Nb.KdSkIpILHGbGi",
			inputPassword: "!R8c%bN2Yx",
			hasError:      false,
		},
	}

	for _, c := range cases {
		actual := CheckPasswordHash(c.inputHash, c.inputPassword)
		hasErr := false

		if actual != nil {
			hasErr = true
		}
		if hasErr != c.hasError {
			t.Errorf("Test Failed => actual has error: %v is not same as expected has error: %v", hasErr, c.hasError)
			continue
		}
	}
}

func TestMakeValidateJWT(t *testing.T) {
	id1 := uuid.New()
	token1, _ := MakeJWT(id1, "secret", time.Hour)

	cases := []struct {
		name        string
		tokenSecret string
		tokenString string
		wantUserID  uuid.UUID
		wantErr     bool
	}{
		{
			name:        "Valid token",
			tokenSecret: "secret",
			tokenString: token1,
			wantUserID:  id1,
			wantErr:     false,
		},
		{
			name:        "Invalid token",
			tokenSecret: "secret",
			tokenString: "invalid.token.string",
			wantUserID:  uuid.Nil,
			wantErr:     true,
		},
		{
			name:        "Wrong secret",
			tokenSecret: "wrong_secret",
			tokenString: token1,
			wantUserID:  uuid.Nil,
			wantErr:     true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			gotUserID, err := ValidateJWT(c.tokenString, c.tokenSecret)
			if (err != nil) != c.wantErr {
				t.Errorf("ValidateJWT() error = %v, wantErr %v", err, c.wantErr)
				return
			}
			if gotUserID != c.wantUserID {
				t.Errorf("ValidateJWT() gotUserID = %v, want %v", gotUserID, c.wantUserID)
			}
		})
	}
}
