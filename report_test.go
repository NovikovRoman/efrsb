package efrsb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReport(t *testing.T) {
	tests := []struct {
		name    string
		guid    string
		wantErr bool
	}{
		{
			name:    "0a9393e6",
			guid:    "0a9393e6-d292-4490-b9df-11ed3b858759",
			wantErr: false,
		},
		{
			name:    "empty",
			guid:    "",
			wantErr: true,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, err := testClient.Report(ctx, tt.guid)
			if err != nil && !tt.wantErr || err == nil && tt.wantErr {
				t.Errorf("Report() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}

			assert.NotNil(t, r.DatePublish)
		})
	}
}
