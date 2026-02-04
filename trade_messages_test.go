package efrsb

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTradeMessages(t *testing.T) {
	dBegin := time.Date(2020, 6, 1, 0, 0, 0, 0, time.UTC)
	dEnd := time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name    string
		filter  TradeMessageFilter
		wantErr bool
	}{
		{
			name: "By datePublish",
			filter: TradeMessageFilter{
				DatePublishBegin: dBegin,
				DatePublishEnd:   dEnd,
			},
			wantErr: false,
		},
		{
			name: "By bankrupt GUID",
			filter: TradeMessageFilter{
				BankruptGuid: []string{
					"e7c22dfb-59ef-b28b-8ed4-b6fc93173563",
				},
			},
			wantErr: false,
		},
		{
			name: "By GUID",
			filter: TradeMessageFilter{
				Guid: []string{
					"66bdf555-1a3f-4a3e-98f6-c196c13c5d60",
				},
			},
			wantErr: false,
		},
		{
			name: "By TradeNumber",
			filter: TradeMessageFilter{
				TradeNumber: []string{
					"108–ОАОФ",
				},
			},
			wantErr: false,
		},
		{
			name: "Error by datePublishEnd",
			filter: TradeMessageFilter{
				DatePublishBegin: dBegin,
			},
			wantErr: true,
		},
		{
			name: "By type",
			filter: TradeMessageFilter{
				DatePublishBegin: dBegin,
				DatePublishEnd:   dEnd,
				Type: []string{
					EtpBiddingResult,
				},
			},
			wantErr: false,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := testClient.TradeMessages(ctx, tt.filter, 0, 500)
			if err != nil && !tt.wantErr || err == nil && tt.wantErr {
				t.Errorf("TradeMessages() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}
			assert.Greater(t, m.Total, 0)
			assert.LessOrEqual(t, len(m.Items), m.Total)
		})
	}
}
