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

type SroFilter struct {
	DateLastModifBegin time.Time
	DateLastModifEnd   time.Time
	SearchString       string
	IsActive           ThreePositionSwitch
	Guid               []string
}

type Sro struct {
	Guid             string    `json:"guid"`
	DateLastModifRaw string    `json:"dateLastModif"`
	DateLastModif    time.Time `json:"-"`
	Shortname        string    `json:"shortName"`
	Fullname         string    `json:"fullName"`
	Ogrn             string    `json:"ogrn"`
	Inn              string    `json:"inn"`
	Kpp              string    `json:"kpp"`
	Regnum           string    `json:"regnum"`
	DateRegisterRaw  string    `json:"dateRegister"`
	DateRegister     time.Time `json:"-"`
	DateExcludeRaw   string    `json:"dateExclude"`
	DateExclude      time.Time `json:"-"`
	CauseExclude     string    `json:"causeExclude"`
	Address          string    `json:"address"`
}

type SroResult struct {
	Total int   `json:"total"`
	Items []Sro `json:"pageData"`
}

// Sros возвращает список СРО
func (c *Client) Sros(ctx context.Context, filter SroFilter, offset, limit int) (result SroResult, err error) {
	if limit < 1 {
		return
	}

	if limit > 500 {
		err = NewErrParamLimit()
		return
	}

	if len(filter.Guid) == 0 && filter.SearchString == "" && filter.IsActive == PositionSwitchUnknown &&
		(filter.DateLastModifBegin.IsZero() || filter.DateLastModifEnd.IsZero()) {
		err = ErrRequiredParam{
			message: "Не заполнен обязательный параметр запроса - dateLastModifBegin и dateLastModifEnd, тк не указан searchString или guid или isActive"}
		return
	}

	data := url.Values{}
	data.Set("searchString", filter.SearchString)
	data.Set("guid", strings.Join(filter.Guid, ","))
	data.Set("limit", strconv.Itoa(limit))
	data.Set("offset", strconv.Itoa(offset))

	switch filter.IsActive {
	case PositionSwitchYes:
		data.Set("isActive", "true")
	case PositionSwitchNo:
		data.Set("isActive", "false")
	}

	if !filter.DateLastModifBegin.IsZero() {
		data.Set("dateLastModifBegin", filter.DateLastModifBegin.Format(asRFC3339))
	}
	if !filter.DateLastModifEnd.IsZero() {
		data.Set("dateLastModifEnd", filter.DateLastModifEnd.Format(asRFC3339))
	}

	var b []byte
	if b, err = c.get(ctx, "/v1/sros?"+data.Encode()); err != nil {
		return
	}

	if err = json.Unmarshal(b, &result); err != nil {
		err = fmt.Errorf("Unmarshal: %w Body: %s", err, string(b))
	}

	for i, item := range result.Items {
		result.Items[i].DateLastModif = parser.DateTime(item.DateLastModifRaw, asRFC3339)
		result.Items[i].DateRegister = parser.DateTime(item.DateRegisterRaw, asRFC3339)
		result.Items[i].DateExclude = parser.DateTime(item.DateExcludeRaw, asRFC3339)
	}
	return
}
