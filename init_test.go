package efrsb

import (
	"context"
)

const (
	testLogin    = "demowebuser"
	testPassword = "Ax!761BN"
)

var testAuth *Auth

func init() {
	ctx := context.Background()
	cfg := NewAuthConfig(testLogin, testPassword).Dev()

	var err error
	testAuth, err = NewAuth(ctx, cfg)
	if err != nil {
		panic(err)
	}
}
