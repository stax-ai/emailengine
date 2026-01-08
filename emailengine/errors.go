package emailengine

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type EEError struct {
	Status    int
	Code      string
	Message   string
	Operation string
	Retryable bool
	Raw       []byte
}

func (e *EEError) Error() string {
	return fmt.Sprintf("emailengine: %s (status=%d code=%s)", e.Message, e.Status, e.Code)
}

func decodeError(op string, res *http.Response) error {
	if res == nil {
		return &EEError{Status: 0, Message: "no response", Operation: op, Retryable: true}
	}
	body, _ := io.ReadAll(res.Body)
	_ = res.Body.Close()
	// Try to parse common error shape
	var kv map[string]any
	var code, msg string
	if json.Unmarshal(body, &kv) == nil {
		if v, ok := kv["error"].(string); ok {
			code = v
		}
		if v, ok := kv["message"].(string); ok {
			msg = v
		} else if v, ok := kv["response"].(string); ok {
			msg = v
		}
	}
	if msg == "" {
		msg = http.StatusText(res.StatusCode)
	}
	retryable := res.StatusCode == http.StatusTooManyRequests || res.StatusCode == http.StatusBadGateway || res.StatusCode == http.StatusServiceUnavailable || res.StatusCode == http.StatusGatewayTimeout
	return &EEError{
		Status:    res.StatusCode,
		Code:      code,
		Message:   msg,
		Operation: op,
		Retryable: retryable,
		Raw:       body,
	}
}
