package efrsb

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestArbitrManagers(t *testing.T) {
	dBegin := time.Date(2020, 6, 1, 0, 0, 0, 0, time.UTC)
	dEnd := time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name    string
		filter  ArbitrManagerFilter
		wantErr bool
	}{
		{
			name: "By datePublish",
			filter: ArbitrManagerFilter{
				DateLastModifBegin: dBegin,
				DateLastModifEnd:   dEnd,
			},
			wantErr: false,
		},
		{
			name: "By bankrupt GUID",
			filter: ArbitrManagerFilter{
				SroGuid: "e85a144d-bf94-44e7-8e56-d7271f1e1ec7",
			},
			wantErr: false,
		},
		{
			name: "By GUID",
			filter: ArbitrManagerFilter{
				Guid: []string{
					"0572ce9f-11d9-4507-8bbe-954c592baf7d",
				},
			},
			wantErr: false,
		},
		{
			name: "Error by datePublishEnd",
			filter: ArbitrManagerFilter{
				DateLastModifBegin: dBegin,
			},
			wantErr: true,
		},
		{
			name: "By isActive",
			filter: ArbitrManagerFilter{
				IsActive: PositionSwitchNo,
			},
			wantErr: false,
		},
		{
			name: "By searchString",
			filter: ArbitrManagerFilter{
				SearchString: "Спинозович",
			},
			wantErr: false,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := testClient.ArbitrManagers(ctx, tt.filter, 0, 500)
			if err != nil && !tt.wantErr || err == nil && tt.wantErr {
				t.Errorf("ArbitrManagers() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}

			assert.Greater(t, m.Total, 0)
			assert.LessOrEqual(t, len(m.Items), m.Total)
		})
	}
}
