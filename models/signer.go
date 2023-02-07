package models

type Signer struct {
	Name                    string `json:"name"`
	Email                   string `json:"email"`
	Lock_email              bool   `json:"lock_email"`
	Lock_phone              bool   `json:"lock_phone"`
	Phone_country           string `json:"phone_country"`
	Phone_number            string `json:"phone_number"`
	Auth_mode               string `json:"auth_mode"`
	Send_automatic_email    bool   `json:"send_automatic_email"`
	Send_automatic_whatsapp bool   `json:"send_automatic_whatsapp"`
}

func (signer Signer) CreateSigner() []Signer {
	var fakeSigners = []Signer{
		{
			Name:                    "John Doe",
			Email:                   "test@test.com",
			Auth_mode:               "assinaturaTela",
			Send_automatic_email:    false,
			Send_automatic_whatsapp: true,
		},
		{
			Name: "Jane Doe",
		},
	}

	return fakeSigners
}
