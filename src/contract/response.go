package contract

type (
	MessageResponse struct {
		Message string `json:"message"`
	}
	ErrorResponse struct {
		StatusCode int    `json:"status"`
		Message    string `json:"message"`
	}
)
