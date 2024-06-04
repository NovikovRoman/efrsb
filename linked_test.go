package efrsb

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLinkedMessages(t *testing.T) {
	tests := []struct {
		name    string
		guid    string
		wantErr bool
	}{
		{
			name:    "Available",
			guid:    "21851419-F027-0259-B5C4-17A442A5319C",
			wantErr: false,
		},
		{
			name:    "Not found",
			guid:    "43be2e40-74c4",
			wantErr: true,
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
			m, err := client.LinkedMessages(ctx, tt.guid)
			if err != nil && !tt.wantErr || err == nil && tt.wantErr {
				t.Errorf("LinkedMessages() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}

			assert.Greater(t, len(m), 0)
		})
	}
}

func TestLinkedReports(t *testing.T) {
	tests := []struct {
		name    string
		guid    string
		wantErr bool
	}{
		{
			name:    "Available",
			guid:    "0a9393e6-d292-4490-b9df-11ed3b858759",
			wantErr: false,
		},
		{
			name:    "Not found",
			guid:    "0a9393e6-d292",
			wantErr: true,
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
			m, err := client.LinkedReports(ctx, tt.guid)
			if err != nil && !tt.wantErr || err == nil && tt.wantErr {
				t.Errorf("LinkedReports() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}

			assert.Greater(t, len(m), 0)
		})
	}
}
