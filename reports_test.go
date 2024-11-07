package efrsb

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestReports(t *testing.T) {
	dBegin := time.Date(2020, 6, 1, 0, 0, 0, 0, time.UTC)
	dEnd := time.Date(2020, 7, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name    string
		filter  ReportFilter
		wantErr bool
	}{
		{
			name: "By datePublish",
			filter: ReportFilter{
				DatePublishBegin: dBegin,
				DatePublishEnd:   dEnd,
			},
			wantErr: false,
		},
		{
			name: "By bankrupt GUID",
			filter: ReportFilter{
				BankruptGuid: []string{
					"281011dd-2228-0ada-3224-3e93055ad42c",
				},
			},
			wantErr: false,
		},
		{
			name: "By GUID",
			filter: ReportFilter{
				Guid: []string{
					"38cd692f-ec25-4a37-af0b-3bb6e213f809",
				},
			},
			wantErr: false,
		},
		{
			name: "Error by datePublishEnd",
			filter: ReportFilter{
				DatePublishBegin: dBegin,
			},
			wantErr: true,
		},
		{
			name: "By type",
			filter: ReportFilter{
				DatePublishBegin: dBegin,
				DatePublishEnd:   dEnd,
				Type: []string{
					ReportFinal,
				},
			},
			wantErr: false,
		},
	}

	ctx := context.Background()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := testClient.Reports(ctx, tt.filter, 0, 500)
			if err != nil && !tt.wantErr || err == nil && tt.wantErr {
				t.Errorf("Reports() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}

			assert.Greater(t, m.Total, 0)
			assert.Equal(t, len(m.Items), m.Total)
		})
	}
}
