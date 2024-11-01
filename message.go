package efrsb

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/NovikovRoman/efrsb/parser"
)

type Message struct {
	Guid                 string    `json:"guid"`
	BankruptGuid         string    `json:"bankruptGUID"`
	AnnulmentMessageGuid string    `json:"annulmentMessageGuid"`
	Number               string    `json:"number"`
	DatePublishRaw       string    `json:"datePublish"`
	DatePublish          time.Time `json:"-"`
	Content              string    `json:"content"`
	Type                 string    `json:"type"`
	LockReason           string    `json:"lockReason"`
	HasViolation         bool      `json:"hasViolation"`
}

// Message возвращает сообщение
func (c *Client) Message(ctx context.Context, guid string) (result Message, err error) {
	if guid == "" {
		err = ErrRequiredParam{message: "Не указан обязательный параметр guid"}
		return
	}

	var b []byte
	if b, err = c.get(ctx, "/v1/messages/"+guid); err != nil {
		return
	}

	if err = json.Unmarshal(b, &result); err != nil {
		err = fmt.Errorf("Unmarshal: %w Body: %s", err, string(b))
	}

	result.DatePublish = parser.DateTime(result.DatePublishRaw, asRFC3339)
	return
}
