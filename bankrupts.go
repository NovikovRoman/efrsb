package efrsb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/NovikovRoman/efrsb/parser"
)

type BankruptFilter struct {
	Type      string
	Guid      []string
	Name      string
	Ogrn      string
	Ogrnip    string
	Inn       string
	Snils     string
	Birthdate time.Time
}

type BankruptResult struct {
	Total int        `json:"total"`
	Items []Bankrupt `json:"pageData"`
}

type Bankrupt struct {
	Guid string `json:"guid"`
	Type string `json:"type"`
	Data struct {
		Name          string    `json:"name"`
		Address       string    `json:"address"`
		Lastname      string    `json:"lastName"`
		Firstname     string    `json:"firstName"`
		Middlename    string    `json:"middleName"`
		Ogrn          string    `json:"ogrn"`
		Ogrnip        string    `json:"ogrnip"`
		Inn           string    `json:"inn"`
		Snils         string    `json:"snils"`
		BirthdateRaw  string    `json:"birthdate"`
		Birthdate     time.Time `json:"-"`
		Birthplace    string    `json:"birthplace"`
		NameHistories []string  `json:"nameHistories"`
	} `json:"data"`
}

// Bankrupts возвращает список банкротов
func (c *Client) Bankrupts(ctx context.Context, filter BankruptFilter, offset, limit int) (result BankruptResult, err error) {
	if limit < 1 {
		return
	}

	if limit > 500 {
		err = NewErrParamLimit()
		return
	}

	if filter.Type == "" && len(filter.Guid) == 0 {
		err = ErrRequiredParam{"Не заполнен обязательный параметр запроса - type или guid"}
		return
	}

	// если указан тип, то обязательно указано что-то из
	// для Person:
	// name | ogrnip | inn | snils
	// для Company:
	// name | ogrn | inn
	switch filter.Type {
	case TypePerson:
		if filter.Name == "" && filter.Ogrnip == "" && filter.Inn == "" && filter.Snils == "" {
			err = ErrRequiredParam{
				message: "Не заполнен обязательный параметр запроса - name или ogrnip или inn или snils"}
			return
		}

	case TypeCompany:
		if filter.Name == "" && filter.Ogrn == "" && filter.Inn == "" {
			err = ErrRequiredParam{
				message: "Не заполнен обязательный параметр запроса - name или ogrn или inn"}
			return
		}

	default:
		if filter.Name == "" && filter.Ogrn == "" && filter.Ogrnip == "" && filter.Inn == "" && filter.Snils == "" {
			err = ErrRequiredParam{
				message: "Не заполнен обязательный параметр запроса - name или ogrn или ogrnip или inn или snils"}
			return
		}
	}

	data := url.Values{}
	data.Set("type", filter.Type)
	data.Set("guid", strings.Join(filter.Guid, ","))
	data.Set("name", filter.Name)
	data.Set("ogrn", filter.Ogrn)
	data.Set("ogrnip", filter.Ogrnip)
	data.Set("inn", filter.Inn)
	data.Set("snils", filter.Snils)
	data.Set("limit", strconv.Itoa(limit))
	data.Set("offset", strconv.Itoa(offset))
	if !filter.Birthdate.IsZero() {
		data.Set("birthdate", filter.Birthdate.Format(time.DateOnly))
	}

	var b []byte
	if b, err = c.get(ctx, "/v1/bankrupts?"+data.Encode()); err != nil {
		return
	}

	if err = json.Unmarshal(b, &result); err != nil {
		err = fmt.Errorf("Unmarshal: %w Body: %s", err, string(b))
	}

	for i, item := range result.Items {
		result.Items[i].Data.Birthdate = parser.DateTime(item.Data.BirthdateRaw, asRFC3339)
	}
	return
}
