package efrsb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func Test_NewAuth(t *testing.T) {
	ctx := context.Background()
	cfg := NewAuthConfig("demowebuser", "Ax!761BN").Dev()
	auth, err := NewAuth(ctx, cfg)
	require.Nil(t, err)

	assert.True(t, auth.isDev)
	ok, err := auth.IsActiveToken()
	assert.Nil(t, err)
	assert.True(t, ok)
}
