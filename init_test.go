package efrsb

import (
	"context"
)

const (
	testLogin    = "demowebuser"
	testPassword = "Ax!761BN"
)

var testClient *Client

func init() {
	testClient = New(testLogin, testPassword, Dev())

	ctx := context.Background()
	if err := testClient.Auth(ctx); err != nil {
		panic(err)
	}
}
