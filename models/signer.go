package models

type Signer struct {
	Name                    string
	Email                   string `json:"email"`
	Lock_email              bool   `json:"lock_email"`
	Lock_phone              bool   `json:"lock_phone"`
	Phone_country           string `json:"phone_country"`
	Phone_number            string `json:"phone_number"`
	Auth_mode               string `json:"auth_mode"`
	Send_automatic_email    bool   `json:"send_automatic_email"`
	Send_automatic_whatsapp bool   `json:"send_automatic_whatsapp"`
}
