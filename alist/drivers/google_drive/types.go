package google_drive

type TokenError struct {
	Error            string `json:"error"`
	ErrorDescription string `json:"error_description"`
}
