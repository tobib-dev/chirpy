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
	id2 := uuid.New()
	id3 := uuid.New()

	secret1 := "gY8vFzM9L1bXnqC7sZ0eT2dPaRmVkWjU"
	secret2 := "tQ4eYkH7cLpXvRgZwN1oDbMuA6sJfB2K"
	secret3 := "Zx9VrEwTbG5mKoYqNdSfApLuQ3cXnJ7M"

	expiresHour := time.Hour * 1
	expiresMinute := time.Minute * 30
	expiresSeconds := time.Second * 45

	token1, _ := MakeJWT(id1, secret1, expiresHour)
	token2, _ := MakeJWT(id2, secret2, expiresMinute)
	token3, _ := MakeJWT(id3, secret3, expiresSeconds)

	cases := []struct {
		name          string
		inputID       uuid.UUID
		inputSecret   string
		expectedToken string
		wantErr       bool
	}{
		{
			name:          "Correct id",
			inputID:       id1,
			inputSecret:   secret1,
			expectedToken: token1,
			wantErr:       false,
		},
		{
			name:          "Incorrect id",
			inputID:       id2,
			inputSecret:   secret2,
			expectedToken: token3,
			wantErr:       true,
		},
		{
			name:          "Empty id",
			inputID:       uuid.Nil,
			inputSecret:   secret1,
			expectedToken: token2,
			wantErr:       true,
		},
		{
			name:          "Invalid Token",
			inputID:       id3,
			inputSecret:   secret3,
			expectedToken: "invalid token",
			wantErr:       true,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			_, err := ValidateJWT(c.expectedToken, c.inputSecret)
			if (err != nil) != c.wantErr {
				t.Errorf("ValidateJWT() error = %v, wantErr %v", err, c.wantErr)
			}
		})
	}
}
