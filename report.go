package efrsb

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/NovikovRoman/efrsb/parser"
)

type Report struct {
	Guid                string    `json:"guid"`
	BankruptGuid        string    `json:"bankruptGUID"`
	AnnulmentReportGuid string    `json:"annulmentReportGUID"`
	Number              string    `json:"number"`
	DatePublishRaw      string    `json:"datePublish"`
	DatePublish         time.Time `json:"-"`
	Content             string    `json:"content"`
	Type                string    `json:"type"`
	ProcedureType       string    `json:"procedureType"`
	LockReason          string    `json:"lockReason"`
}

// Report возвращает данные по отчету
func (c *Client) Report(ctx context.Context, guid string) (result Report, err error) {
	if guid == "" {
		err = ErrRequiredParam{message: "Не указан обязательный параметр guid"}
		return
	}

	var b []byte
	if b, err = c.get(ctx, "/v1/reports/"+guid); err != nil {
		return
	}

	if err = json.Unmarshal(b, &result); err != nil {
		err = fmt.Errorf("Unmarshal: %w Body: %s", err, string(b))
	}

	result.DatePublish = parser.DateTime(result.DatePublishRaw, asRFC3339)
	return
}
