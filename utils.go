package efrsb

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func responseErrHandler(body []byte, statusCode int) (err error) {
	if statusCode == http.StatusUnauthorized {
		err = ErrUnauthorized{}
		return
	}

	if statusCode == http.StatusBadRequest {
		var sErr errService
		if err = json.Unmarshal(body, &sErr); err != nil {
			err = fmt.Errorf("ErrHandler Unmarshal: %w StatusCode: %d Body: %s", err, statusCode, body)
			return
		}

		err = sErr.Error()
		return
	}

	if statusCode == http.StatusNotFound {
		err = ErrNotFound{}
		return
	}

	if statusCode != http.StatusOK {
		err = fmt.Errorf("StatusCode: %d %s", statusCode, body)
	}
	return
}
