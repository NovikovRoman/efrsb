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

type TradeMessageFilter struct {
	DatePublishBegin   time.Time
	DatePublishEnd     time.Time
	AuctionMessageGuid []string
	Number             []string
	Guid               []string
	Type               []string
	CourtDecisionType  []string
	IsAnnulled         ThreePositionSwitch
	IsLocked           ThreePositionSwitch
	TradeNumber        []string
	TradePlaceGuid     []string
	BankruptGuid       []string
	IncludeContent     bool
	Sort               SortOrder
}

type TradeMessageResult struct {
	Total int            `json:"total"`
	Items []TradeMessage `json:"pageData"`
}

type TradeMessage struct {
	Guid                 string    `json:"guid"`
	Number               string    `json:"number"`
	DatePublishRaw       string    `json:"datePublish"`
	DatePublish          time.Time `json:"-"`
	DateEventRaw         string    `json:"dateEvent"`
	DateEvent            time.Time `json:"-"`
	Type                 string    `json:"type"`
	AnnulmentMessageGuid string    `json:"annulmentMessageGuid"`
	IsLocked             bool      `json:"isLocked"`
	Content              string    `json:"content"`
	TradePlaceGuid       string    `json:"tradePlaceGUID"`
	BankruptGuid         string    `json:"bankruptGUID"`
	AuctionMessageGuid   string    `json:"auctionMessageGuid"`
	Trade                Trade     `json:"trade"`
}

type Trade struct {
	Guid   string `json:"guid"`
	Number string `json:"number"`
}

// TradeMessages возвращает список сообщений
func (c *Client) TradeMessages(ctx context.Context, filter TradeMessageFilter, offset, limit int) (result TradeMessageResult, err error) {
	if limit < 1 {
		return
	}

	if limit > 500 {
		err = NewErrParamLimit()
		return
	}

	if len(filter.Guid) == 0 && len(filter.BankruptGuid) == 0 && len(filter.Number) == 0 &&
		len(filter.TradeNumber) == 0 && (filter.DatePublishBegin.IsZero() || filter.DatePublishEnd.IsZero()) {
		err = ErrRequiredParam{
			message: "Не заполнен обязательный параметр запроса - datePublishBegin и datePublishEnd, тк не указан number или guid или bankruptGuid или tradeNumber"}
		return
	}

	data := url.Values{}
	data.Set("number", strings.Join(filter.Number, ","))
	data.Set("guid", strings.Join(filter.Guid, ","))
	data.Set("type", strings.Join(filter.Type, ","))
	data.Set("courtDecisionType", strings.Join(filter.CourtDecisionType, ","))
	data.Set("bankruptGUID", strings.Join(filter.BankruptGuid, ","))
	data.Set("auctionMessageGuid", strings.Join(filter.AuctionMessageGuid, ","))
	data.Set("tradeNumber", strings.Join(filter.TradeNumber, ","))
	data.Set("tradePlaceGUID", strings.Join(filter.TradePlaceGuid, ","))
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
	if b, err = c.get(ctx, "/v1/trade-messages?"+data.Encode()); err != nil {
		return
	}

	if err = json.Unmarshal(b, &result); err != nil {
		err = fmt.Errorf("Unmarshal: %w Body: %s", err, string(b))
	}

	for i, item := range result.Items {
		result.Items[i].DatePublish = parser.DateTime(item.DatePublishRaw, asRFC3339)
		result.Items[i].DateEvent = parser.DateTime(item.DateEventRaw, asRFC3339)
	}
	return
}
