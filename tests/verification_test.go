package tests

import (
	"context"
	"fmt"
	"line_integrals_fuel_efficiency/util"
	"line_integrals_fuel_efficiency/util/types"
	"testing"

	"github.com/stretchr/testify/assert"
)

var token string = "eyJhbGciOiJSUzI1NiIsImtpZCI6IjQ4NmYxNjQ4MjAwNWEyY2RhZjI2ZDkyMTQwMThkMDI5Y2E0NmZiNTYiLCJ0eXAiOiJKV1QifQ.eyJpc3MiOiJhY2NvdW50cy5nb29nbGUuY29tIiwiYXpwIjoiNjE4MTA0NzA4MDU0LTlyOXMxYzRhbGczNmVybGl1Y2hvOXQ1Mm4zMm42ZGdxLmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwiYXVkIjoiNjE4MTA0NzA4MDU0LTlyOXMxYzRhbGczNmVybGl1Y2hvOXQ1Mm4zMm42ZGdxLmFwcHMuZ29vZ2xldXNlcmNvbnRlbnQuY29tIiwic3ViIjoiMTAxNjUxMDI4NzA0NDM0NjI1NjE1IiwiZW1haWwiOiJ2bW9kMjAwNUBnbWFpbC5jb20iLCJlbWFpbF92ZXJpZmllZCI6dHJ1ZSwiYXRfaGFzaCI6IjdMakZwUV9vYXJiRVJydTgwTHdrMHciLCJpYXQiOjE2NTI5OTI1MzEsImV4cCI6MTY1Mjk5NjEzMSwianRpIjoiYmExMDgxODU1MGVhYTU4NzAzMDIxMDY3OTM2NWNhZTBjNjQ4MTcwYyJ9.Qgv3mrBaJxFQ3JtNFuzftHM4D4Q2e2Pci-HKtkhLQ-tJeRBsRCwYD5qYSQnzgB5ZpgZwF9nyzzjAoL99fIAPQGufjMuEHVuyjNpRcQXJw7-vdF14zmekqUu_mPVDccr2A0r-CftkLVKrDK6Au8X3Cwh8nbRBgU_t0rFNSAJm7KpeVnrzAteI6D1O7BebvQNpGq5YQPM-SQNsWDkWY_4diLYMHZrbgRfgM4C10rG2Eu1pRYRF8TKlrLl-g23ejnxsxinQFysh-4q6z0NCD12bUnjtf2bftKOtSmuFUHVIgr_ysFu-9s1UrAVGguuHV3nPXGCPhc2y0eMubQgd8Gy2PA"

func TestVerification(t *testing.T) {
	util.InitClient()
	valid, err := util.VerifyToken(token, context.Background(), "vmod2005@gmail.com", types.ClientId(types.TestMode))
	if err != nil {
		fmt.Println(err)
	}
	assert.Equal(t, valid, true, "Token is invalid")
}
