package dto

type LoginFormDto struct {
	Phone    string `json:"phone,omitempty"`
	Code     string `json:"code"`
	Password string `json:"password"`
}
