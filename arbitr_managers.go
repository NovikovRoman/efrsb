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

type ArbitrManagerFilter struct {
	DateLastModifBegin time.Time
	DateLastModifEnd   time.Time
	SearchString       string
	IsActive           ThreePositionSwitch
	SroGuid            string
	Guid               []string
}

type ArbitrManager struct {
	Guid             string        `json:"guid"`
	DateLastModifRaw string        `json:"dateLastModif"`
	DateLastModif    time.Time     `json:"-"`
	SroGuid          string        `json:"sroGuid"`
	Lastname         string        `json:"lastName"`
	Firstname        string        `json:"firstName"`
	Middlename       string        `json:"middleName"`
	Ogrnip           string        `json:"ogrnip"`
	Inn              string        `json:"inn"`
	Snils            string        `json:"snils"`
	Regnum           string        `json:"regnum"`
	DateRegisterRaw  string        `json:"dateRegister"`
	DateRegister     time.Time     `json:"-"`
	DateExcludeRaw   string        `json:"dateExclude"`
	DateExclude      time.Time     `json:"-"`
	CauseExclude     string        `json:"causeExclude"`
	RegionCode       string        `json:"regionCode"`
	NameHistory      []NameHistory `json:"nameHistory"`
}

type NameHistory struct {
	Lastname   string `json:"lastName"`
	Firstname  string `json:"firstName"`
	Middlename string `json:"middleName"`
}

type ArbitrManagerResult struct {
	Total int             `json:"total"`
	Items []ArbitrManager `json:"pageData"`
}

// ArbitrManagers возвращает список арбитражных управляющих
func (c *Client) ArbitrManagers(ctx context.Context, filter ArbitrManagerFilter, offset, limit int) (result ArbitrManagerResult, err error) {
	if limit < 1 {
		return
	}

	if limit > 500 {
		err = NewErrParamLimit()
		return
	}

	if len(filter.Guid) == 0 && filter.SearchString == "" && filter.IsActive == PositionSwitchUnknown &&
		filter.SroGuid == "" && (filter.DateLastModifBegin.IsZero() || filter.DateLastModifEnd.IsZero()) {
		err = ErrRequiredParam{
			message: "Не заполнен обязательный параметр запроса - dateLastModifBegin и dateLastModifEnd, тк не указан searchString или guid или isActive или sroGuid"}
		return
	}

	data := url.Values{}
	data.Set("searchString", filter.SearchString)
	data.Set("guid", strings.Join(filter.Guid, ","))
	data.Set("sroGuid", filter.SroGuid)
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
	if b, err = c.get(ctx, "/v1/arbitr-managers?"+data.Encode()); err != nil {
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
