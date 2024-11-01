package parser

import (
	"os"
	"reflect"
	"testing"
	"time"
)

func TestFindCourtDecree(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		want     *CourtDecree
	}{
		{
			name:     "arb",
			filepath: "../testdata/arb.xml",
			want: &CourtDecree{
				ID:              "79",
				Name:            "Арбитражный суд города Санкт-Петербурга и Ленинградской области",
				Number:          "А56-105696/2019",
				DecisionDateRaw: "2020-01-20",
				DecisionDate:    time.Date(2020, 1, 20, 0, 0, 0, 0, time.UTC),
				DecisionType: &DecisionType{
					Name: "о введении наблюдения",
					ID:   "11",
				},
			},
		},
		{
			name:     "arb2",
			filepath: "../testdata/arb2.xml",
			want: &CourtDecree{
				ID:              "46",
				Name:            "Арбитражный суд Республики Северная Осетия - Алания",
				Number:          "А61-1055/2016",
				DecisionDateRaw: "2016-05-19",
				DecisionDate:    time.Date(2016, 5, 19, 0, 0, 0, 0, time.UTC),
				DecisionType: &DecisionType{
					Name: "о признании действий (бездействий) арбитражного управляющего незаконными",
					ID:   "22",
				},
			},
		},
		{
			name:     "arb3",
			filepath: "../testdata/arb3.xml",
			want: &CourtDecree{
				ID:              "66",
				Name:            "Арбитражный суд Владимирской области",
				Number:          "А11-32165/2019",
				DecisionDateRaw: "2023-07-06",
				DecisionDate:    time.Date(2023, 7, 6, 0, 0, 0, 0, time.UTC),
				DecisionType: &DecisionType{
					Name: "о завершении реализации имущества гражданина",
					ID:   "25",
				},
				Discharged: true,
			},
		},
		{
			name:     "Not Found",
			filepath: "../testdata/monetaryObligations.xml",
			want:     nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := os.ReadFile(tt.filepath)
			if err != nil {
				t.Fatal(err)
			}
			if got := FindCourtDecree(string(b)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindCourtDecree() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
