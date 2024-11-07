package efrsb

import (
	"context"
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessage(t *testing.T) {
	tests := []struct {
		name    string
		guid    string
		wantErr error
	}{
		{
			name:    "b94e0a31",
			guid:    "b94e0a31-ebc7-c728-d754-41fbd20fa4f8",
			wantErr: ErrNotFound{},
		},
		{
			name:    "deea9d05",
			guid:    "deea9d05-9b04-44f5-9f55-64ef53108021",
			wantErr: nil,
		},
		{
			name:    "64b009a4",
			guid:    "64b009a4-543c-4f33-bcdd-746cc358ee48",
			wantErr: nil,
		},
		{
			name:    "empty",
			guid:    "",
			wantErr: errors.New("Не указан обязательный параметр guid"),
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := testClient.Message(ctx, tt.guid)
			if err != nil && tt.wantErr == nil || err == nil && tt.wantErr != nil {
				t.Errorf("Message() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}

			assert.NotNil(t, m.DatePublish)
		})
	}
}
