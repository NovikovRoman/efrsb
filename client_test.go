package efrsb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNewClient(t *testing.T) {
	ctx := context.Background()
	client := New("demowebuser", "Ax!761BN", Dev())
	err := client.RefreshToken(ctx)
	require.Nil(t, err)

	assert.True(t, client.isDev)
	ok, err := client.IsActiveToken()
	assert.Nil(t, err)
	assert.True(t, ok)

	token := client.token.Raw
	err = client.RefreshToken(ctx)
	require.Nil(t, err)
	assert.NotEqual(t, token, client.token.Raw)
}
