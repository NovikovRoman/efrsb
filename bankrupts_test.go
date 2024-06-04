package efrsb

import (
	"context"
	"testing"
)

func TestBankrupts(t *testing.T) {
	tests := []struct {
		name    string
		filter  BankruptFilter
		wantErr bool
	}{
		{
			name: "Person name",
			filter: BankruptFilter{
				Type: TypePerson,
				Name: "Иванов",
			},
			wantErr: false,
		},
		{
			name: "Company name",
			filter: BankruptFilter{
				Type: TypePerson,
				Name: "Иванов",
			},
			wantErr: false,
		},
		{
			name: "no type",
			filter: BankruptFilter{
				Name: "Иванов",
			},
			wantErr: true,
		},
	}

	ctx := context.Background()
	client := New(testAuth)

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := client.Bankrupts(ctx, tt.filter, 0, 1)
			if err != nil && !tt.wantErr || err == nil && tt.wantErr {
				t.Errorf("Bankrupts() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
