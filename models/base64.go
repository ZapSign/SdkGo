package models

type Base64Pdf struct {
	Name       string   `json:"name"`
	Base64_pdf string   `json:"base64_pdf"`
	Signers    []Signer `json:"signers"`
}

type Base64Docx struct {
	Name        string   `json:"name"`
	Base64_docx string   `json:"base64_docx"`
	Signers     []Signer `json:"signers"`
}
