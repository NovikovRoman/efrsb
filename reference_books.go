package efrsb

import (
	"context"
	"encoding/json"
	"fmt"
)

type ReferenceBookMessageType struct {
	Code      string `json:"code"`
	Name      string `json:"name"`
	IsOld     bool   `json:"isOld"`
	UnionType string `json:"unionType"`
}

// ReferenceBookMessageTypes возвращает типы сообщений
func (c *Client) ReferenceBookMessageTypes(ctx context.Context) (result []ReferenceBookMessageType, err error) {
	var b []byte
	if b, err = c.get(ctx, "/v1/reference-books/message-types"); err != nil {
		return
	}

	if err = json.Unmarshal(b, &result); err != nil {
		err = fmt.Errorf("Unmarshal: %w Body: %s", err, string(b))
	}
	return
}

type CourtDecisionType struct {
	Code  string `json:"code"`
	Name  string `json:"name"`
	IsOld bool   `json:"isOld"`
}

// ReferenceBookMessageTypes возвращает типы судебных актов
func (c *Client) ReferenceBookCourtDecisionTypes(ctx context.Context) (result []CourtDecisionType, err error) {
	var b []byte
	if b, err = c.get(ctx, "/v1/reference-books/court-decision-types"); err != nil {
		return
	}

	if err = json.Unmarshal(b, &result); err != nil {
		err = fmt.Errorf("Unmarshal: %w Body: %s", err, string(b))
	}
	return
}
