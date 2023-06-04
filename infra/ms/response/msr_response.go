package response

import "encoding/json"

type MsrResponse struct {
	Code int
	Msg  string
	Data json.RawMessage
}
