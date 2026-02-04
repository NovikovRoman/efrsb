package efrsb

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestAmSroMemberships(t *testing.T) {
	dBegin := time.Date(2020, 6, 1, 0, 0, 0, 0, time.UTC)
	dEnd := time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name    string
		filter  AmSroFilter
		wantErr bool
	}{
		{
			name: "By datePublish",
			filter: AmSroFilter{
				DateLastModifBegin: dBegin,
				DateLastModifEnd:   dEnd,
			},
			wantErr: false,
		},
		{
			name: "By isAnnulled",
			filter: AmSroFilter{
				DateLastModifBegin: dBegin,
				DateLastModifEnd:   dEnd,
				IsAnnulled:         PositionSwitchYes,
			},
			wantErr: false,
		},
		{
			name: "Error by datePublishEnd",
			filter: AmSroFilter{
				DateLastModifBegin: dBegin,
			},
			wantErr: true,
		},
		{
			name: "Error by isAnnulled",
			filter: AmSroFilter{
				IsAnnulled: PositionSwitchYes,
			},
			wantErr: true,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := testClient.AmSroMemberships(ctx, tt.filter, 0, 500)
			if err != nil && !tt.wantErr || err == nil && tt.wantErr {
				t.Errorf("AmSroMemberships() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}

			assert.Greater(t, m.Total, 0)
			assert.LessOrEqual(t, len(m.Items), m.Total)
		})
	}
}
