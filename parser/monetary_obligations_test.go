package parser

import (
	"os"
	"reflect"
	"testing"
)

func TestFindMonetaryObligations(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		want     MonetaryObligations
	}{
		{
			name:     "arb",
			filepath: "../testdata/monetaryObligations.xml",
			want: MonetaryObligations{
				Entrepreneurship: []MonetaryObligation{},
				NonEntrepreneurship: []MonetaryObligation{
					{
						CreditorName:   "Общество с ограниченной ответственностью \"Ренессанс Кредит\"",
						CreditorRegion: "г. Москва",
						Content:        "Задолженность по кредитным платежам (кроме ипотеки)",
						Basis:          "Судебный приказ от 09.12.2022 №СП2-4624/2022-12 Судебный участок №12 г. Петрозаводска",
						TotalSum:       63449.69,
						DebtSum:        63449.69,
						PenaltySum:     0,
					},
					{
						CreditorName:   "Акционерное общество \"Банк Русский Стандарт\"",
						CreditorRegion: "г. Москва",
						Content:        "Задолженность по кредитным платежам (кроме ипотеки)",
						Basis:          "Судебный приказ от 15.01.2022 №СП2-4583/2021-12 Судебный участок №12 г. Петрозаводска",
						TotalSum:       55020.91,
						DebtSum:        55020.91,
						PenaltySum:     0,
					},
					{
						CreditorName:   "Публичное акционерное общество КБ \"Восточный\"",
						CreditorRegion: "г. Москва",
						Content:        "Задолженность по кредитным платежам (кроме ипотеки)",
						Basis:          "Судебный приказ от 22.06.2021 №СП2-1775/2021-12 Судебный участок №12 г. Петрозаводска",
						TotalSum:       61694.17,
						DebtSum:        61694.17,
						PenaltySum:     0,
					},
					{
						CreditorName:   "Публичное акционерное общество \"Совкомбанк\"",
						CreditorRegion: "Костромская область",
						Content:        "Задолженность по кредитным платежам (кроме ипотеки)",
						Basis:          "Судебный приказ от 28.05.2021 №СП2-1393/2021-12 Судебный участок №12 г. Петрозаводска",
						TotalSum:       102928.34,
						DebtSum:        102928.34,
						PenaltySum:     0,
					},
					{
						CreditorName:   "Публичное акционерное общество \"Совкомбанк\"",
						CreditorRegion: "Костромская область",
						Content:        "Задолженность по кредитным платежам (кроме ипотеки)",
						Basis:          "Судебный приказ от 25.05.2021 №СП2-1319/2021-12 Судебный участок №12 г. Петрозаводска",
						TotalSum:       10272.32,
						DebtSum:        10272.32,
						PenaltySum:     0,
					},
				},
			},
		},
		{
			name:     "Not Found",
			filepath: "../testdata/arb.xml",
			want: MonetaryObligations{
				NonEntrepreneurship: []MonetaryObligation{},
				Entrepreneurship:    []MonetaryObligation{},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := os.ReadFile(tt.filepath)
			if err != nil {
				t.Fatal(err)
			}
			if got := FindMonetaryObligations(string(b)); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("FindMonetaryObligations() = %v, want %v", got, tt.want)
			}
		})
	}
}
