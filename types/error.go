package types

type JSONError struct {
	Status *Status `json:"status"`
}

type Status struct {
	Timestamp    string `json:"timestamp"`
	ErrorCode    int    `json:"error_code"`
	ErrorMessage string `json:"error_message"`
	CreditCount  int    `json:"credit_count,omitempty"`
}
