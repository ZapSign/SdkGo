package models

type SignBatchAPI struct {
	User_token    string    `json:"user_token"`
	Signer_tokens [2]string `json:"signer_tokens"`
}
