package parser

import (
	"os"
	"reflect"
	"testing"
)

func TestFindPersonCategory(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		want     *PersonCategory
	}{
		{
			name:     "arb",
			filepath: "../testdata/monetaryObligations.xml",
			want: &PersonCategory{
				Code:        "PensionRecipient",
				Description: "На основании соблюдения условий, предусмотренных п.п. 2 п. 1 ст. 223.2 Федерального закона от 26.10.2002 № 127-ФЗ \"О несостоятельности (банкротстве)\"",
			},
		},
		{
			name:     "Not Found",
			filepath: "../testdata/arb.xml",
			want:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := os.ReadFile(tt.filepath)
			if err != nil {
				t.Fatal(err)
			}
			if got := FindPersonCategory(string(b)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindPersonCategory() = %v, want %v", got, tt.want)
			}
		})
	}
}
