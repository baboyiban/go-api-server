package dto

type ErrorResponse struct {
	Error   string `json:"error"`
	Details string `json:"details,omitempty"`
}
