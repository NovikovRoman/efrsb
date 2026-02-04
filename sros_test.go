package efrsb

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestSros(t *testing.T) {
	dBegin := time.Date(2017, 2, 1, 0, 0, 0, 0, time.UTC)
	dEnd := time.Date(2017, 3, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name    string
		filter  SroFilter
		wantErr bool
	}{
		{
			name: "By datePublish",
			filter: SroFilter{
				DateLastModifBegin: dBegin,
				DateLastModifEnd:   dEnd,
			},
			wantErr: false,
		},
		{
			name: "By isActive",
			filter: SroFilter{
				IsActive: PositionSwitchYes,
			},
			wantErr: false,
		},
		{
			name: "By GUID",
			filter: SroFilter{
				Guid: []string{
					"6f51518d-6181-41f9-88c2-70717006c51a",
				},
			},
			wantErr: false,
		},
		{
			name: "Error by datePublishEnd",
			filter: SroFilter{
				DateLastModifEnd: dBegin,
			},
			wantErr: true,
		},
		{
			name: "By searchString",
			filter: SroFilter{
				SearchString: "Ассоциация",
			},
			wantErr: false,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := testClient.Sros(ctx, tt.filter, 0, 500)
			if err != nil && !tt.wantErr || err == nil && tt.wantErr {
				t.Errorf("Sros() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}

			assert.Greater(t, m.Total, 0)
			assert.LessOrEqual(t, len(m.Items), m.Total)
		})
	}
}
