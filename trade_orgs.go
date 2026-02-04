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

type TradeOrgFilter struct {
	DateLastModifBegin time.Time
	DateLastModifEnd   time.Time
	SearchString       string
	Guid               []string
}

type CmpTradeOrg struct {
	Guid             string    `json:"guid"`
	DateLastModifRaw string    `json:"dateLastModif"`
	DateLastModif    time.Time `json:"-"`
	Shortname        string    `json:"shortName"`
	Fullname         string    `json:"fullName"`
	Ogrn             string    `json:"ogrn"`
	Inn              string    `json:"inn"`
	Kpp              string    `json:"kpp"`
}

type CmpTradeOrgResult struct {
	Total int           `json:"total"`
	Items []CmpTradeOrg `json:"pageData"`
}

// CmpTradeOrgs возвращает список организаторов торгов – юридических лиц
func (c *Client) CmpTradeOrgs(ctx context.Context, filter TradeOrgFilter, offset, limit int) (result CmpTradeOrgResult, err error) {
	if limit < 1 {
		return
	}

	if limit > 500 {
		err = NewErrParamLimit()
		return
	}

	if len(filter.Guid) == 0 && filter.SearchString == "" &&
		(filter.DateLastModifBegin.IsZero() || filter.DateLastModifEnd.IsZero()) {
		err = ErrRequiredParam{
			message: "Не заполнен обязательный параметр запроса - dateLastModifBegin и dateLastModifEnd, тк не указан searchString или guid"}
		return
	}

	data := url.Values{}
	data.Set("searchString", filter.SearchString)
	data.Set("guid", strings.Join(filter.Guid, ","))
	data.Set("limit", strconv.Itoa(limit))
	data.Set("offset", strconv.Itoa(offset))

	if !filter.DateLastModifBegin.IsZero() {
		data.Set("dateLastModifBegin", filter.DateLastModifBegin.Format(asRFC3339))
	}
	if !filter.DateLastModifEnd.IsZero() {
		data.Set("dateLastModifEnd", filter.DateLastModifEnd.Format(asRFC3339))
	}

	var b []byte
	if b, err = c.get(ctx, "/v1/cmp-trade-orgs?"+data.Encode()); err != nil {
		return
	}

	if err = json.Unmarshal(b, &result); err != nil {
		err = fmt.Errorf("Unmarshal: %w Body: %s", err, string(b))
	}

	for i, item := range result.Items {
		result.Items[i].DateLastModif = parser.DateTime(item.DateLastModifRaw, asRFC3339)
	}
	return
}

type PrsnTradeOrg struct {
	Guid             string      `json:"guid"`
	DateLastModifRaw string      `json:"dateLastModif"`
	DateLastModif    time.Time   `json:"-"`
	Lastname         string      `json:"lastName"`
	Firstname        string      `json:"firstName"`
	Middlename       string      `json:"middleName"`
	Ogrnip           string      `json:"ogrnip"`
	Inn              string      `json:"inn"`
	NameHistory      NameHistory `json:"nameHistory"`
}

type PrsnTradeOrgResult struct {
	Total int            `json:"total"`
	Items []PrsnTradeOrg `json:"pageData"`
}

// PrsnTradeOrgs возвращает список организаторов торгов – физических лиц
func (c *Client) PrsnTradeOrgs(ctx context.Context, filter TradeOrgFilter, offset, limit int) (result PrsnTradeOrgResult, err error) {
	if limit < 1 {
		return
	}

	if limit > 500 {
		err = NewErrParamLimit()
		return
	}

	if len(filter.Guid) == 0 && filter.SearchString == "" &&
		(filter.DateLastModifBegin.IsZero() || filter.DateLastModifEnd.IsZero()) {
		err = ErrRequiredParam{
			message: "Не заполнен обязательный параметр запроса - dateLastModifBegin и dateLastModifEnd, тк не указан searchString или guid"}
		return
	}

	data := url.Values{}
	data.Set("searchString", filter.SearchString)
	data.Set("guid", strings.Join(filter.Guid, ","))
	data.Set("limit", strconv.Itoa(limit))
	data.Set("offset", strconv.Itoa(offset))

	if !filter.DateLastModifBegin.IsZero() {
		data.Set("dateLastModifBegin", filter.DateLastModifBegin.Format(asRFC3339))
	}
	if !filter.DateLastModifEnd.IsZero() {
		data.Set("dateLastModifEnd", filter.DateLastModifEnd.Format(asRFC3339))
	}

	var b []byte
	if b, err = c.get(ctx, "/v1/prsn-trade-orgs?"+data.Encode()); err != nil {
		return
	}

	if err = json.Unmarshal(b, &result); err != nil {
		err = fmt.Errorf("Unmarshal: %w Body: %s", err, string(b))
	}

	for i, item := range result.Items {
		result.Items[i].DateLastModif = parser.DateTime(item.DateLastModifRaw, asRFC3339)
	}
	return
}
