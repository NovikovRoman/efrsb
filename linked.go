package efrsb

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/NovikovRoman/efrsb/parser"
)

type LinkedMessage struct {
	Guid                 string    `json:"guid"`
	Number               string    `json:"number"`
	DatePublishRaw       string    `json:"datePublish"`
	DatePublish          time.Time `json:"-"`
	DateInvisibleRaw     string    `json:"dateInvisible"`
	DateInvisible        time.Time `json:"-"`
	Type                 string    `json:"type"`
	AnnulmentMessageGuid string    `json:"annulmentMessageGuid"`
	LockReason           string    `json:"lockReason"`
	ContentMessageGuids  []string  `json:"contentMessageGuids"`
}

// LinkedMessages возвращает список связанных сообщений
func (c *Client) LinkedMessages(ctx context.Context, guid string) (result []LinkedMessage, err error) {
	if guid == "" {
		err = ErrRequiredParam{message: "Не указан обязательный параметр guid"}
		return
	}

	var b []byte
	if b, err = c.get(ctx, "/v1/messages/"+guid+"/linked"); err != nil {
		return
	}

	if err = json.Unmarshal(b, &result); err != nil {
		err = fmt.Errorf("Unmarshal: %w Body: %s", err, string(b))
	}

	for i := range result {
		result[i].DatePublish = parser.DateTime(result[i].DatePublishRaw, asRFC3339)
		result[i].DateInvisible = parser.DateTime(result[i].DateInvisibleRaw, asRFC3339)
	}
	return
}

type LinkedReport struct {
	Guid                 string    `json:"guid"`
	Number               string    `json:"number"`
	DatePublishRaw       string    `json:"datePublish"`
	DatePublish          time.Time `json:"-"`
	Type                 string    `json:"type"`
	ProcedureType        string    `json:"procedureType"`
	AnnulmentMessageGuid string    `json:"annulmentMessageGuid"`
	LockReason           string    `json:"lockReason"`
	ContentReportGuids   []string  `json:"contentReportGuids"`
}

// LinkedReports возвращает список связанных отчетов
func (c *Client) LinkedReports(ctx context.Context, guid string) (result []LinkedReport, err error) {
	if guid == "" {
		err = ErrRequiredParam{message: "Не указан обязательный параметр guid"}
		return
	}

	var b []byte
	if b, err = c.get(ctx, "/v1/reports/"+guid+"/linked"); err != nil {
		return
	}

	if err = json.Unmarshal(b, &result); err != nil {
		err = fmt.Errorf("Unmarshal: %w Body: %s", err, string(b))
	}

	for i := range result {
		result[i].DatePublish = parser.DateTime(result[i].DatePublishRaw, asRFC3339)
	}
	return
}
