package sms

import "fmt"

type Config struct {
	Endpoint string `json:"endpoint"`
	Account  string `json:"account"`
	Secret   string `json:"secret"`
	Template string `json:"template"`
}

func (sms *Config) Render(code string) string {
	return fmt.Sprintf(sms.Template, code)
}
