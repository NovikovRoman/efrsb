package efrsb

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCmpTradeOrgs(t *testing.T) {
	dBegin := time.Date(2013, 6, 1, 0, 0, 0, 0, time.UTC)
	dEnd := time.Date(2013, 7, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name    string
		filter  TradeOrgFilter
		wantErr bool
	}{
		{
			name: "By datePublish",
			filter: TradeOrgFilter{
				DateLastModifBegin: dBegin,
				DateLastModifEnd:   dEnd,
			},
			wantErr: false,
		},
		{
			name: "By GUID",
			filter: TradeOrgFilter{
				Guid: []string{
					"4e4b4a72-2668-41a7-8896-e91590cf2a92",
				},
			},
			wantErr: false,
		},
		{
			name: "Error by datePublishEnd",
			filter: TradeOrgFilter{
				DateLastModifBegin: dBegin,
			},
			wantErr: true,
		},
		{
			name: "By searchString",
			filter: TradeOrgFilter{
				SearchString: "Систем",
			},
			wantErr: false,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := testClient.CmpTradeOrgs(ctx, tt.filter, 0, 500)
			if err != nil && !tt.wantErr || err == nil && tt.wantErr {
				t.Errorf("CmpTradeOrgs() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}
			assert.Greater(t, m.Total, 0)
			assert.LessOrEqual(t, len(m.Items), m.Total)
		})
	}
}

func TestPrsnTradeOrgs(t *testing.T) {
	dBegin := time.Date(2013, 6, 1, 0, 0, 0, 0, time.UTC)
	dEnd := time.Date(2013, 7, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name    string
		filter  TradeOrgFilter
		wantErr bool
	}{
		{
			name: "By datePublish",
			filter: TradeOrgFilter{
				DateLastModifBegin: dBegin,
				DateLastModifEnd:   dEnd,
			},
			wantErr: false,
		},
		{
			name: "By GUID",
			filter: TradeOrgFilter{
				Guid: []string{
					"fb783383-73a6-4aec-af03-f99a574696f0",
				},
			},
			wantErr: false,
		},
		{
			name: "Error by datePublishEnd",
			filter: TradeOrgFilter{
				DateLastModifBegin: dBegin,
			},
			wantErr: true,
		},
		{
			name: "By searchString",
			filter: TradeOrgFilter{
				SearchString: "Алюкаев",
			},
			wantErr: false,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := testClient.PrsnTradeOrgs(ctx, tt.filter, 0, 500)
			if err != nil && !tt.wantErr || err == nil && tt.wantErr {
				t.Errorf("PrsnTradeOrgs() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}
			assert.Greater(t, m.Total, 0)
			assert.LessOrEqual(t, len(m.Items), m.Total)
		})
	}
}
