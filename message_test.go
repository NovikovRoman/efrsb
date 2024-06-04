package efrsb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	tests := []struct {
		name    string
		guid    string
		wantErr bool
	}{
		{
			name:    "deea9d05",
			guid:    "deea9d05-9b04-44f5-9f55-64ef53108021",
			wantErr: false,
		},
		{
			name:    "64b009a4",
			guid:    "64b009a4-543c-4f33-bcdd-746cc358ee48",
			wantErr: false,
		},
		{
			name:    "empty",
			guid:    "",
			wantErr: true,
		},
	}

	ctx := context.Background()
	client := New(testAuth)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := client.Message(ctx, tt.guid)
			if err != nil && !tt.wantErr || err == nil && tt.wantErr {
				t.Errorf("Message() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}

			assert.NotNil(t, m.DatePublish)
		})
	}
}
