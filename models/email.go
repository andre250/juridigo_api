package models

type SendEmail struct {
	To   string `json:"to"`
	Name string `json:"name,omitempty"`
	Type int    `json:"type"`
}
