package parser

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestFindPublisher(t *testing.T) {
	tests := []struct {
		name     string
		filepath string
		want     *Publisher
	}{
		{
			name:     "arb",
			filepath: "../testdata/arb.xml",
			want: &Publisher{
				Person: &PublisherPerson{
					Lastname:              "Васильев",
					Firstname:             "Иван",
					Middlename:            "Петрович",
					Inn:                   "110115686508",
					Snils:                 "11111111111",
					CorrespondenceAddress: "test",
				},
				Sro: &Sro{
					Name:    "Ассоциация \"Первая СРО АУ\" - Ассоциация «Первая Саморегулируемая Организация Арбитражных Управляющих зарегистрированная в едином государственном реестре саморегулируемых организаций арбитражных управляющих»",
					Ogrn:    "1025203032150",
					Inn:     "5260111551",
					Address: "109029, г. Москва, ул. Скотопрогонная, д. 29/1",
				},
			},
		},
		{
			name:     "arb2",
			filepath: "../testdata/arb2.xml",
			want: &Publisher{
				Company: &PublisherCompany{
					Name: "Государственная корпорация «Агентство по страхованию вкладов»",
					Ogrn: "1047796046198",
					Inn:  "7708514824",
				},
			},
		},
		{
			name:     "arb3",
			filepath: "../testdata/arb3.xml",
			want: &Publisher{
				Person: &PublisherPerson{
					Lastname:              "Васильев",
					Firstname:             "Иван",
					Middlename:            "Петрович",
					Inn:                   "110115686508",
					Snils:                 "36819469638",
					CorrespondenceAddress: "Москва, улица Пушкина",
				},
				Sro: &Sro{
					Name:    "Ассоциация \"Первая СРО АУ\" - Ассоциация «Первая Саморегулируемая Организация Арбитражных Управляющих зарегистрированная в едином государственном реестре саморегулируемых организаций арбитражных управляющих»",
					Ogrn:    "1025203032150",
					Inn:     "5260111551",
					Address: "109029, г. Москва, ул. Скотопрогонная, д. 29/1",
				},
			},
		},
		{
			name:     "arb4",
			filepath: "../testdata/arb4.xml",
			want: &Publisher{
				Person: &PublisherPerson{
					Lastname:              "Тухикова",
					Firstname:             "Юлиана",
					Middlename:            "Александровна",
					Inn:                   "190507849809",
					Snils:                 "06880247094",
					CorrespondenceAddress: "125009, г. Москва, ул. Тверская, д. 19, а/я 46, для Тухиковой Ю.А.",
					Email:                 "email@email.ru",
				},
				Sro: &Sro{
					Name:    "Ассоциация \"Саморегулируемая организация арбитражных управляющих \"Меркурий\"",
					Ogrn:    "1037710023108",
					Inn:     "7710458616",
					Address: "127018, г Москва, Сущевский Вал, 16, 4, оф.301 (фактический адрес)",
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := os.ReadFile(tt.filepath)
			if err != nil {
				t.Fatal(err)
			}
			if got := FindPublisher(string(b)); !reflect.DeepEqual(got, tt.want) {
				fmt.Println(got.Sro.Name)
				t.Errorf("FindPublisher() = %+v, want %+v", got, tt.want)
			}
		})
	}
}
