package efrsb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMessageFiles(t *testing.T) {
	tests := []struct {
		name    string
		guid    string
		wantErr bool
	}{
		{
			name:    "No files",
			guid:    "deea9d05-9b04-44f5-9f55-64ef53108021",
			wantErr: false,
		},
		{
			name:    "Files available",
			guid:    "A56A9F5F-0FD0-4DDB-A13B-74962D894DD8",
			wantErr: false,
		},
		{
			name:    "empty guid",
			guid:    "",
			wantErr: true,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := testClient.MessageFiles(ctx, tt.guid, true)
			if err != nil && !tt.wantErr || err == nil && tt.wantErr {
				t.Errorf("MessageFiles() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}

			assert.NotEmpty(t, b)
		})
	}
}
func TestReportFiles(t *testing.T) {
	tests := []struct {
		name    string
		guid    string
		wantErr bool
	}{
		{
			name:    "No files",
			guid:    "0a9393e6-d292-4490-b9df-11ed3b858759",
			wantErr: false,
		},
		{
			name:    "empty guid",
			guid:    "",
			wantErr: true,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := testClient.ReportFiles(ctx, tt.guid, true)
			if err != nil && !tt.wantErr || err == nil && tt.wantErr {
				t.Errorf("ReportFiles() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}

			assert.NotEmpty(t, b)
		})
	}
}
