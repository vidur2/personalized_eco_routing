package util

import (
	"context"
	"encoding/json"
	"line_integrals_fuel_efficiency/util/types"
	"time"
)

// var clientId string =

type Claims struct {
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Name          string `json:"name"`
	Picture       string `json:"picture"`
	GivenName     string `json:"given_name"`
	FamilyName    string `json:"family_name"`
	Locale        string `json:"locale"`
}

func VerifyToken(token string, ctx context.Context, email string, clientId types.ClientId) (bool, error) {
	payload, err := Verify.Validate(ctx, token, string(clientId))

	if err != nil {
		return false, err
	} else {
		currentTime := time.Now().Unix()
		var claims Claims
		b, err := json.Marshal(payload.Claims)
		if err != nil {
			return false, err
		}
		err = json.Unmarshal(b, &claims)

		if err != nil {
			return false, err
		}

		if currentTime <= int64(payload.Expires) && payload.Audience == string(clientId) && email == claims.Email {
			return true, nil
		} else {
			return false, nil
		}
	}
}
