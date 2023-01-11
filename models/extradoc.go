package models

type ExtraDoc struct {
	Name       string `json:"name"`
	Base64_pdf string `json:"base64_pdf"`
}

type ExtraDocResponse struct {
	Open_id                                 int32
	Token, Name, Original_file, Signed_file string
}
