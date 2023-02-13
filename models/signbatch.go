package models

type SignBatchAPI struct {
	User_token    string    `json:"user_token"`
	Signer_tokens [2]string `json:"signer_tokens"`
}

func (signBatch SignBatchAPI) CreateSignBatch() SignBatchAPI {

	signerTokens := [...]string{"b15edab2-b691-4ecb-833a-564b3490963e", "9934376a-58fd-443a-ab9d-f5210895e0e0"}

	return SignBatchAPI{
		User_token:    "1c166935-5b8a-4059-a205-4f87d7119f9b",
		Signer_tokens: signerTokens,
	}

}
