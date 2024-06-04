package efrsb

import (
	"context"
	"strconv"
)

// MessageFiles возвращает архив файлов сообщения
func (c *Client) MessageFiles(ctx context.Context, guid string, onlySafe bool) (b []byte, err error) {
	if guid == "" {
		err = ErrRequiredParam{message: "Не указан обязательный параметр guid"}
		return
	}

	b, err = c.get(ctx, "/v1/messages/"+guid+"/files/archive?onlySafe="+strconv.FormatBool(onlySafe))
	return
}

// ReportFiles возвращает архив файлов сообщения
func (c *Client) ReportFiles(ctx context.Context, guid string, onlySafe bool) (b []byte, err error) {
	if guid == "" {
		err = ErrRequiredParam{message: "Не указан обязательный параметр guid"}
		return
	}

	b, err = c.get(ctx, "/v1/reports/"+guid+"/files/archive?onlySafe="+strconv.FormatBool(onlySafe))
	return
}
