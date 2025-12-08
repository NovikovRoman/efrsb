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

type MessageFilter struct {
	DatePublishBegin  time.Time
	DatePublishEnd    time.Time
	Number            []string
	Guid              []string
	Type              []string
	CourtDecisionType []string
	IsAnnulled        ThreePositionSwitch
	IsLocked          ThreePositionSwitch
	BankruptGuid      []string
	IncludeContent    bool
	Sort              SortOrder
}

type MessageResult struct {
	Total int       `json:"total"`
	Items []Message `json:"pageData"`
}

// Messages возвращает список сообщений
func (c *Client) Messages(ctx context.Context, filter MessageFilter, offset, limit int) (result MessageResult, err error) {
	if limit < 1 {
		return
	}

	if limit > 500 {
		err = NewErrParamLimit()
		return
	}

	if len(filter.Guid) == 0 && len(filter.BankruptGuid) == 0 && len(filter.Number) == 0 &&
		(filter.DatePublishBegin.IsZero() || filter.DatePublishEnd.IsZero()) {
		err = ErrRequiredParam{
			message: "Не заполнен обязательный параметр запроса - datePublishBegin и datePublishEnd, тк не указан number или guid или bankruptGuid"}
		return
	}

	data := url.Values{}
	data.Set("number", strings.Join(filter.Number, ","))
	data.Set("guid", strings.Join(filter.Guid, ","))
	data.Set("type", strings.Join(filter.Type, ","))
	data.Set("courtDecisionType", strings.Join(filter.CourtDecisionType, ","))
	data.Set("bankruptGUID", strings.Join(filter.BankruptGuid, ","))
	data.Set("limit", strconv.Itoa(limit))
	data.Set("offset", strconv.Itoa(offset))
	data.Set("includeContent", fmt.Sprintf("%t", filter.IncludeContent))
	data.Set("sort", string(filter.Sort))

	switch filter.IsAnnulled {
	case PositionSwitchYes:
		data.Set("isAnnulled", "true")
	case PositionSwitchNo:
		data.Set("isAnnulled", "false")
	}

	switch filter.IsLocked {
	case PositionSwitchYes:
		data.Set("isLocked", "true")
	case PositionSwitchNo:
		data.Set("isLocked", "false")
	}

	if !filter.DatePublishBegin.IsZero() {
		data.Set("datePublishBegin", filter.DatePublishBegin.Format(asRFC3339))
	}
	if !filter.DatePublishEnd.IsZero() {
		data.Set("datePublishEnd", filter.DatePublishEnd.Format(asRFC3339))
	}

	var b []byte
	if b, err = c.get(ctx, "/v1/messages?"+data.Encode()); err != nil {
		return
	}

	if err = json.Unmarshal(b, &result); err != nil {
		err = fmt.Errorf("Unmarshal: %w Body: %s", err, string(b))
	}

	for i, item := range result.Items {
		result.Items[i].DatePublish = parser.DateTime(item.DatePublishRaw, asRFC3339)
	}
	return
}
