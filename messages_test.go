package efrsb

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestMessages(t *testing.T) {
	dBegin := time.Date(2023, 6, 1, 0, 0, 0, 0, time.UTC)
	dEnd := time.Date(2023, 7, 1, 0, 0, 0, 0, time.UTC)
	tests := []struct {
		name    string
		filter  MessageFilter
		wantErr bool
	}{
		/* {
			name: "By datePublish",
			filter: MessageFilter{
				DatePublishBegin: dBegin,
				DatePublishEnd:   dEnd,
			},
			wantErr: false,
		},
		{
			name: "By bankrupt GUID",
			filter: MessageFilter{
				BankruptGuid: []string{
					"a79f9366-32f4-ef38-b8b4-22253ffd47a9",
					"c8796d66-2a15-a47a-23a4-22824c0160e2",
				},
			},
			wantErr: false,
		},
		{
			name: "By GUID",
			filter: MessageFilter{
				Guid: []string{
					"deea9d05-9b04-44f5-9f55-64ef53108021",
					"64b009a4-543c-4f33-bcdd-746cc358ee48",
				},
			},
			wantErr: false,
		},
		{
			name: "Error by datePublishEnd",
			filter: MessageFilter{
				DatePublishBegin: dBegin,
			},
			wantErr: true,
		}, */
		{
			name: "By type",
			filter: MessageFilter{
				DatePublishBegin: dBegin,
				DatePublishEnd:   dEnd,
				Type: []string{
					MessageArbitralDecree,
				},
			},
			wantErr: false,
		},
	}

	ctx := context.Background()
	client := New(testAuth)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m, err := client.Messages(ctx, tt.filter, 0, 500)
			if err != nil && !tt.wantErr || err == nil && tt.wantErr {
				t.Errorf("Messages() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err != nil {
				return
			}

			assert.Greater(t, m.Total, 0)
			assert.Equal(t, len(m.Items), m.Total)

			for _, item := range m.Items {
				mes, _ := client.Message(ctx, item.Guid)
				fmt.Printf("%+v\n\n\n", mes)
			}
		})
	}
}
