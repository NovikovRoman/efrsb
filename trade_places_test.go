package efrsb

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTradePlaces(t *testing.T) {
	dBegin := time.Date(2022, 6, 1, 0, 0, 0, 0, time.UTC)
	dEnd := time.Date(2022, 7, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name    string
		filter  TradePlaceFilter
		wantErr bool
	}{
		{
			name: "By datePublish",
			filter: TradePlaceFilter{
				DateLastModifBegin: dBegin,
				DateLastModifEnd:   dEnd,
			},
			wantErr: false,
		},
		{
			name: "By isActive",
			filter: TradePlaceFilter{
				IsActive: PositionSwitchYes,
			},
			wantErr: false,
		},
		{
			name: "By GUID",
			filter: TradePlaceFilter{
				Guid: []string{
					"ccab978a-7cb2-47be-bec3-911b14a22d95",
				},
			},
			wantErr: false,
		},
		{
			name: "Error by datePublishEnd",
			filter: TradePlaceFilter{
				DateLastModifBegin: dBegin,
			},
			wantErr: true,
		},
		{
			name: "By searchString",
			filter: TradePlaceFilter{
				SearchString: "КУПЕЦЪ",
			},
			wantErr: false,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := testClient.TradePlaces(ctx, tt.filter, 0, 500)
			if err != nil && !tt.wantErr || err == nil && tt.wantErr {
				t.Errorf("TradePlaces() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}
			assert.Greater(t, m.Total, 0)
			assert.LessOrEqual(t, len(m.Items), m.Total)
		})
	}
}
