package util

import "context"

func VerifyToken(token string, ctx context.Context) error {
	_, err := Verify.Validate(ctx, token, "")

	if err != nil {
		return err
	} else {
		return nil
	}
}
