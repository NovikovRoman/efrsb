package efrsb

import (
	"context"
	"encoding/json"
	"fmt"
	"net/url"
	"strconv"
	"time"

	"github.com/NovikovRoman/efrsb/parser"
)

type AmSroFilter struct {
	DateLastModifBegin time.Time
	DateLastModifEnd   time.Time
	ArbitrmanagerGuid  string
	IsAnnulled         ThreePositionSwitch
}

type AmSroMembership struct {
	Guid              string    `json:"guid"`
	DateLastModifRaw  string    `json:"dateLastModif"`
	DateLastModif     time.Time `json:"-"`
	ArbitrmanagerGuid string    `json:"arbitrmanagerGuid"`
	SroGuid           string    `json:"sroGuid"`
	Site              string    `json:"site"`
	DateActionRaw     string    `json:"dateAction"`
	DateAction        time.Time `json:"-"`
	IsIncludeAction   bool      `json:"isIncludeAction"`
	Cause             string    `json:"cause"`
}

type AmSroMembershipResult struct {
	Total int               `json:"total"`
	Items []AmSroMembership `json:"pageData"`
}

// AmSroMemberships возвращает список истории переходов АУ между СРО
func (c *Client) AmSroMemberships(ctx context.Context, filter AmSroFilter, offset, limit int) (result AmSroMembershipResult, err error) {
	if limit < 1 {
		return
	}

	if limit > 500 {
		err = NewErrParamLimit()
		return
	}

	if filter.ArbitrmanagerGuid == "" && (filter.DateLastModifBegin.IsZero() || filter.DateLastModifEnd.IsZero()) {
		err = ErrRequiredParam{
			message: "Не заполнен обязательный параметр запроса - dateLastModifBegin и dateLastModifEnd, тк не указан ArbitrmanagerGuid"}
		return
	}

	data := url.Values{}
	data.Set("arbitrmanagerGuid", filter.ArbitrmanagerGuid)
	data.Set("limit", strconv.Itoa(limit))
	data.Set("offset", strconv.Itoa(offset))

	switch filter.IsAnnulled {
	case PositionSwitchYes:
		data.Set("isAnnulled", "true")
	case PositionSwitchNo:
		data.Set("isAnnulled", "false")
	}

	if !filter.DateLastModifBegin.IsZero() {
		data.Set("dateLastModifBegin", filter.DateLastModifBegin.Format(asRFC3339))
	}
	if !filter.DateLastModifEnd.IsZero() {
		data.Set("dateLastModifEnd", filter.DateLastModifEnd.Format(asRFC3339))
	}

	var b []byte
	if b, err = c.get(ctx, "/v1/am-sro-memberships?"+data.Encode()); err != nil {
		return
	}

	if err = json.Unmarshal(b, &result); err != nil {
		err = fmt.Errorf("Unmarshal: %w Body: %s", err, string(b))
	}

	for i, item := range result.Items {
		result.Items[i].DateLastModif = parser.DateTime(item.DateLastModifRaw, asRFC3339)
		result.Items[i].DateAction = parser.DateTime(item.DateActionRaw, asRFC3339)
	}
	return
}
